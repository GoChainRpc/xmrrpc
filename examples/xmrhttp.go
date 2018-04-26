package main

import "log"
import (
	"github.com/GoChainRpc/xmrrpc"
)

func main() {
	connCfg := &xmrrpc.ConnConfig{
		Host:         "localhost:18083",
		HTTPPostMode: true, //  supports HTTP POST mode
		DisableTLS:   true, //  does not provide TLS by default
	}

	client, err := xmrrpc.New(connCfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(client)
	defer client.Shutdown()

	if height, err := client.GetHeight(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("height: %d", height)
	}

	balance ,err := client.GetBalance()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("balance: %d",balance)
	}

}
