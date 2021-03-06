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

	getAddressResult ,err :=  client.GetAddress()
	if err  != nil {
		log.Fatal("getAddress_err:", err)
	} else {
		log.Println("GetAddress:", getAddressResult)
	}

	//transferDestinations := []xmrjson.TransferDestination{}
	//transferDestination := xmrjson.TransferDestination{}
	//transferDestination.Address = "46fPzT8eJGZdkHuoQSXQSLK8yDFiDLPt4WeRMk5BA87SB7fmrurK2woB1RF9TdmNuUdmp9ZzTEgUMTCWBQiP3SUC2NVZ7Sb"
	//transferDestination.Amount = 12560000000
	//transferDestinations = append(transferDestinations,transferDestination)
	//
	//transferDestination2 := xmrjson.TransferDestination{}
	//transferDestination2.Address = "46fPzT8eJGZdkHuoQSXQSLK8yDFiDLPt4WeRMk5BA87SB7fmrurK2woB1RF9TdmNuUdmp9ZzTEgUMTCWBQiP3SUC2NVZ7Sb"
	//transferDestination2.Amount = 11130000000
	//transferDestinations = append(transferDestinations,transferDestination2)
	//
	//transferResult ,err := client.Transfer(transferDestinations,0,"a6511848208d6e54893cb77e33368d9c245d41e6ee9d6d5e31b85bd51179ee1a")
	//
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	log.Println(transferResult)
	//}

	transferByTxid, err := client.GetTransferByTxid("062901dc61543a57d9ac112d08274501c10119566482bc06d963fb9c3ddd60b7")
	if err != nil {
		log.Fatal("GetTransferByTxid_err:", err)
	} else {
		log.Println("GetTransferByTxid:", transferByTxid)
	}

	err = client.OpenWallet("myWallet", "123456")
	if err != nil {
		log.Fatal("OpenWallet:", err)
	} else {
		log.Println("OpenWallet success")
	}


	err = client.CreateWallet("myWallet", "123456")
	if err != nil {
		log.Fatal("CreateWallet:", err)
	} else {
		log.Println("CreateWallet success")
	}
}
