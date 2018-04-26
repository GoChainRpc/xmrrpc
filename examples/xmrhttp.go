package main

import (
	"log"
	"github.com/GoChainRpc/xmrrpc"
	_ "github.com/GoChainRpc/xmrrpc/xmrjson"
	"github.com/GoChainRpc/xmrrpc/xmrjson"
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
	log.Println(connCfg)
	log.Println(client)
	defer client.Shutdown()

	if height, err := client.GetHeight(); err != nil {
		log.Fatal("GetHeight_err:", err)
	} else {
		log.Printf("height: %d", height)
	}

	var balanceResult xmrjson.GetBalanceResult

	log.Print(balanceResult)

	balance, err := client.GetBalance()
	if err != nil {
		log.Fatal("GetBalance_err:", err)
	} else {
		log.Printf("balance: %d", balance)
	}

}
