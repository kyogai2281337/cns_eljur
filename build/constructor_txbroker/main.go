package main

import (
	"flag"

	constructor_logic_entrypoint "github.com/kyogai2281337/cns_eljur/internal/constructor_logic/scd"
)

var (
	BrokerStr string // --broker
	BufSize   int    // --buffer
)

func main() {
	Init()
	worker := constructor_logic_entrypoint.NewLogicWorker(BrokerStr, BufSize)
	defer worker.Close()
	worker.Serve()
}

// Init parses flags and sets default values for them.
// It is called at the very beginning of main().
//
// Flags:
//
//	-broker:
//	    A string, representing the URL of NATS broker.
//	    Default value: "nats://admin:adminpass@localhost:4222".
func Init() {
	flag.StringVar(&BrokerStr, "broker", "nats://admin:adminpass@nats:4222", "broker url")
	flag.IntVar(&BufSize, "buffer", 25, "buffer size")
}
