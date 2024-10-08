package constructor_logic_entrypoint

import (
	"errors"
	"sync"

	"log"

	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor_logic/logic"
	"github.com/kyogai2281337/cns_eljur/internal/mongo/primitives"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore"
	"github.com/nats-io/nats.go"
)

type CacheItem struct {
	Schedule *constructor.Schedule
	mu       sync.RWMutex
}

func NewCacheItem(sch *constructor.Schedule) *CacheItem {
	if sch == nil {
		return nil
	}
	return &CacheItem{Schedule: sch, mu: sync.RWMutex{}}
}

type LogicWorker struct {
	Broker  *nats.Conn
	dirChan chan *nats.Msg
	extChan chan struct{}
	IdemMap map[string]DirResp
	rwm     sync.RWMutex
	store   store.Store
	// InMem caching & syncing data
	schedBuf map[string]*CacheItem
}

func NewLogicWorker(brokerConnStr string, dirBuf int, storageConnStr string) (*LogicWorker, error) {
	db, err := server.NewDB(storageConnStr)
	if err != nil {
		return nil, err
	}

	store := sqlstore.New(db)
	nc, err := nats.Connect(brokerConnStr)
	if err != nil {
		return nil, err
	}
	nc.Flush()
	worker := &LogicWorker{
		store:   store,
		Broker:  nc,
		dirChan: make(chan *nats.Msg, dirBuf),
		extChan: make(chan struct{}),
		IdemMap: make(map[string]DirResp),
		rwm:     sync.RWMutex{},
		// caching
		schedBuf: make(map[string]*CacheItem),
	}
	// worker.Serve()
	return worker, nil
}

func (w *LogicWorker) Serve() {
	log.Println("Starting constructor logic worker...")
	sub, err := w.Broker.Subscribe("constructor.update", func(msg *nats.Msg) {
		w.dirChan <- msg
	})
	if err != nil {
		log.Println("Failed to subscribe: ", err)
		return
	}

	go func(subscriptions ...*nats.Subscription) {
		for {
			select {
			case msg := <-w.dirChan:
				log.Println("Received directive:", string(msg.Data))
				dir, err := UnmarshalDirective(msg.Data)
				if err != nil {
					log.Println("Failed to unmarshal directive:", err)
					data, _ := NewErrorResp(err).Marshal()
					msg.Respond(data)
					continue
				}

				log.Println("Marshalled directive:", dir)
				w.handleDirective(*dir)
				resp := w.IdemMap[dir.ID]
				log.Println("Response:", resp)
				if resp.Err != nil {
					log.Println("Error in response:", resp.Err)
					data, _ := resp.Marshal()
					msg.Respond(data)
				}

				log.Println("Found response:", resp)
				data, err := resp.Marshal()
				if err != nil {
					log.Println("Failed to marshal response:", err)
					continue
				}

				log.Println("Marshalled response:", data)
				if err := msg.Respond(data); err != nil {
					log.Println("Failed to send response:", err)
				}

				// * There is a logging and solution of the memory leak
				log.Printf("Handled directive for key %s\nGot response: %s\n", dir.ID, data)
				delete(w.IdemMap, dir.ID)

			case <-w.extChan:
				log.Println("Exiting constructor logic worker...")
				for _, sub = range subscriptions {
					sub.Unsubscribe()
					log.Println("Unsubscribed of theme", sub.Subject)
				}

				close(w.dirChan)
				close(w.extChan)
				if err := w.store.Close(); err != nil {
					log.Println("Error closing DB: ", err.Error())
				}
				return
			}
		}
	}(sub)
}

func (w *LogicWorker) Close() {
	w.Broker.Close()
	w.extChan <- struct{}{} // notify the goroutine to exit
}

func (w *LogicWorker) handleDirective(dir Directive) *DirResp {
	resp := new(DirResp)

	log.Printf("Handling directive: %+v\n", dir)

	// Получаем элемент кэша (расписание) с блокировкой чтения на уровне карты кэша
	w.rwm.RLock()
	item, err := w.GetSchedule(dir)
	if err != nil || item == nil {
		resp.Err = errors.New("schedule not found")
		log.Println("Schedule not found:", err)
		return resp
	}

	// Обрабатываем директиву
	switch dir.Type {
	case DirInsert:
		resp = w.InsertTask(dir, item)
	case DirDelete:
		resp = w.DeleteTask(dir, item)
	case DirTX:
		resp = w.TXTask(dir, item)
	case DirRename:
		resp = w.RenameTask(dir, item)
	}

	resp.Data = dir.ScheduleID
	// Проверяем на наличие idempotency и отправляем ответ
	w.rwm.RLock()
	w.IdemMap[dir.ID] = *resp
	w.rwm.RUnlock()
	primitives.NewMongoConn().Schedule().Update(dir.ScheduleID, item.Schedule)
	log.Println("Updated schedule:", dir.ScheduleID)
	return resp
}
