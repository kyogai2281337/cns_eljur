package constructor_logic_entrypoint

import (
	"sync"

	"github.com/gofiber/fiber/v2/log"
	"github.com/nats-io/nats.go"
)

type LogicWorker struct {
	Broker  *nats.Conn
	dirChan chan Directive
	extChan chan struct{}
	IdemMap map[string]chan DirResp
	rwm     sync.RWMutex
}

func NewLogicWorker(brokerConnStr string, dirBuf int) *LogicWorker {
	log.Info("Creating constructor logic worker...")
	nc, err := nats.Connect(brokerConnStr)
	if err != nil {
		log.Fatal(err)
	}
	worker := &LogicWorker{
		Broker:  nc,
		dirChan: make(chan Directive, dirBuf),
		extChan: make(chan struct{}),
		IdemMap: make(map[string]chan DirResp),
		rwm:     sync.RWMutex{},
	}
	worker.Serve()
	return worker
}

func (w *LogicWorker) Serve() {

	log.Info("Starting constructor logic worker...")

	sub, err := w.Broker.Subscribe("constructor", func(msg *nats.Msg) {
		dir, err := UnmarshalDirective(msg.Data)
		if err != nil {
			log.Error(err)
			return
		}
		w.dirChan <- dir
	})
	if err != nil {
		log.Error("Failed to subscribe: ", err)
		return
	}
	defer sub.Unsubscribe()

	go func() {
		for {
			select {
			case dir := <-w.dirChan:
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
				w.Broker.PublishMsg(
					&nats.Msg{
						Subject: "constructor",
						Data:    data,
						Header: nats.Header{
							"idempotency": []string{dir.ID},
						},
					},
				)
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
	w.extChan <- struct{}{} // notify the goroutine to exit
	w.Broker.Close()
}

func (w *LogicWorker) handleDirective(dir Directive) *DirResp {
	log.Info("Handling directive: ", dir)
	resp := new(DirResp)
	switch dir.Type {
	case DirInsert:
		resp = w.InsertTask(dir)
	case DirDelete:
		resp = w.DeleteTask(dir)
	case DirTX:
		resp = w.TXTask(dir)
	}
	w.rwm.RLock()
	respch, ok := w.IdemMap[dir.ID]
	w.rwm.RUnlock()
	if ok {
		respch <- *resp
		close(respch)
	}
	return resp

}

func AssignTask(topic string, dir *Directive) DirResp {
	return DirResp{
		Err:  nil,
		Data: dir,
	}
}
