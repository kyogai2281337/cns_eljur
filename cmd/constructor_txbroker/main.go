package main

import (
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gofiber/fiber/v2/log"
	constructor_logic_entrypoint "github.com/kyogai2281337/cns_eljur/internal/constructor_logic/scd"
)

var (
	BrokerStr      string // --broker
	BufSize        int    // --buffer
	StorageConnStr string // --database
)

func main() {
	Init()
	log.Infof("Got flags: broker=%s buffer=%d database=%s", BrokerStr, BufSize, StorageConnStr)
	worker, err := constructor_logic_entrypoint.NewLogicWorker(BrokerStr, BufSize, StorageConnStr)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем WaitGroup для ожидания завершения воркера
	var wg sync.WaitGroup
	wg.Add(1)

	// Обработка сигнала завершения
	go func() {
		defer wg.Done()
		// Ожидание сигнала завершения
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		<-sigs         // Блокируемся до получения сигнала
		worker.Close() // Завершение работы воркера
	}()

	worker.Serve() // Запуск воркера

	// Ожидание завершения работы горутины обработки сигнала
	wg.Wait()
}

// Init parses flags and sets default values for them.
// It is called at the very beginning of main().
//
// Flags:
//
//	-broker:
//	    A string, representing the URL of NATS broker.
//	    Default value: "nats://@broker:4222".
//
//	-buffer:
//	    An integer, representing the size of the channel, which is used to communicate between
//	    the goroutine that handles incoming messages from broker and the goroutine that
//	    handles outgoing messages to the database.
//	    Default value: 25.
//
//	-database:
//	    A string, representing the URL of the database.
//	    Default value: "admin:Erunda228@tcp(db)/journal".
func Init() {
	flag.StringVar(&BrokerStr, "broker", "nats://@broker:4222", "broker url")
	flag.IntVar(&BufSize, "buffer", 25, "buffer size")
	flag.StringVar(&StorageConnStr, "database", "admin:Erunda228@tcp(db)/journal", "database url")
}
