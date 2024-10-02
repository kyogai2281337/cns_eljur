package constructor_logic_entrypoint

import (
	"sync"

	"github.com/gofiber/fiber/v2/log"
	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor_logic/logic"
	"github.com/kyogai2281337/cns_eljur/internal/mongo/primitives"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
	"github.com/nats-io/nats.go"
)

type CacheItem struct {
	Schedule *constructor.Schedule
	mu       sync.RWMutex
}

func NewCacheItem(sch *constructor.Schedule) *CacheItem {
	return &CacheItem{Schedule: sch, mu: sync.RWMutex{}}
}

type LogicWorker struct {
	Broker  *nats.Conn
	dirChan chan *nats.Msg
	extChan chan struct{}
	IdemMap map[string]chan DirResp
	rwm     sync.RWMutex
	store   store.Store
	// InMem caching & syncing data
	schedBuf map[string]*CacheItem
}

func NewLogicWorker(brokerConnStr string, dirBuf int) *LogicWorker {
	log.Info("Creating constructor logic worker...")
	nc, err := nats.Connect(brokerConnStr)
	if err != nil {
		return nil
	}
	worker := &LogicWorker{
		Broker:  nc,
		dirChan: make(chan *nats.Msg, dirBuf),
		extChan: make(chan struct{}),
		IdemMap: make(map[string]chan DirResp),
		rwm:     sync.RWMutex{},
		// caching
		schedBuf: make(map[string]*CacheItem),
	}
	// worker.Serve()
	return worker
}

func (w *LogicWorker) Serve() {

	log.Info("Starting constructor logic worker...")

	sub, err := w.Broker.Subscribe("constructor.update.request", func(msg *nats.Msg) {
		log.Info("Got message: ", string(msg.Data))
		w.dirChan <- msg
	})
	if err != nil {
		log.Error("Failed to subscribe: ", err)
		return
	}
	defer sub.Unsubscribe()

	log.Info("Successfully subscribed to constructor")

	go func() {
		for {
			select {
			case msg := <-w.dirChan:
				dir, err := UnmarshalDirective(msg.Data)
				if err != nil {
					log.Error(err)
					continue
				}
				log.Info("Got directive: ", dir)
				resp := w.handleDirective(dir)
				if resp.Err != nil {
					log.Error(resp.Err)
					continue
				}
				data, err := resp.Marshal()
				if err != nil {
					log.Error(err)
					continue
				}
				if err := msg.Respond(data); err != nil {
					log.Error(err)
				}

			case <-w.extChan:
				log.Info("Exiting constructor logic worker...")
				close(w.dirChan)
				close(w.extChan)
				return
			}
		}
	}()
}
func (w *LogicWorker) Close() {
	w.Broker.Close()
	w.extChan <- struct{}{} // notify the goroutine to exit
}

func (w *LogicWorker) handleDirective(dir Directive) *DirResp {
	log.Info("Handling directive: ", dir)
	resp := new(DirResp)

	// Получаем элемент кэша (расписание) с блокировкой чтения на уровне карты кэша
	w.rwm.RLock()
	item, exists := w.schedBuf[dir.ID]
	w.rwm.RUnlock()

	if !exists {
		// Если элемент не найден, загружаем и добавляем его в кэш
		w.rwm.Lock()
		item = w.GetSchedule(dir)
		w.schedBuf[dir.ID] = item
		w.rwm.Unlock()
	}

	// Блокируем доступ к расписанию на чтение, чтобы другие горутины могли также читать
	item.mu.RLock()

	// Обрабатываем директиву
	switch dir.Type {
	case DirInsert:
		resp = w.InsertTask(dir, item)
	case DirDelete:
		resp = w.DeleteTask(dir, item)
	case DirTX:
		resp = w.TXTask(dir, item)
	}

	// Проверяем на наличие idempotency и отправляем ответ
	w.rwm.RLock()
	respch, ok := w.IdemMap[dir.ID]
	w.rwm.RUnlock()
	if ok {
		respch <- *resp
		close(respch)
	}
	item.mu.RUnlock()
	primitives.NewMongoConn().Schedule().Update(resp.Data.(string), item.Schedule)

	return resp
}
