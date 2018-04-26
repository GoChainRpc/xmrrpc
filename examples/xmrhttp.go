package main

import (
	"log"
	"github.com/GoChainRpc/xmrrpc"
)

func main() {
	connCfg := &xmrrpc.ConnConfig{
		Host:         "localhost:18083/json_rpc",
		HTTPPostMode: true, //  supports HTTP POST mode
		DisableTLS:   true, //  does not provide TLS by default
	}

	client, err := xmrrpc.New(connCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	if height, err := client.GetHeight(); err != nil {
		log.Fatal("GetHeight_err:", err)
	} else {
		log.Printf("height: %d", height)
	}

	balance, err := client.GetBalance()
	if err != nil {
		log.Fatal("GetBalance_err:", err)
	} else {
		log.Printf("balance: %d", balance)
	}

	transfersResult, err := client.GetTransfers(true, true, false, false,
		false, 1, 1)
	if err != nil {
		log.Fatal("GetTransfers_err:", err)
	} else {
		log.Println("transfersResult:", transfersResult)
	}

}
