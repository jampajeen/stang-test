package ethereum

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/jampajeen/stang-test/monitor-service/core"
	"github.com/jampajeen/stang-test/monitor-service/db"
	"github.com/jampajeen/stang-test/monitor-service/model"
)

type TransactionListener struct {
	db *db.MongoDb
}

func NewTransactionListener(db *db.MongoDb) *TransactionListener {
	return &TransactionListener{db: db}
}

func (t *TransactionListener) Listen() {
	client, err := ethclient.Dial(core.Config.APP.RpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0xaB7B31d116927dA43496c20e8cEF350CE57b91AA")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:

			b1, _ := header.MarshalJSON()
			fmt.Printf("\n%s\n", b1)
			fmt.Println("===================================================")

			fmt.Println(header.Hash())
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("TxHash: " + block.TxHash().String())
			fmt.Println("block.Number: " + block.Number().Text(10))
			for _, tx := range block.Transactions() {
				fmt.Println("===================================================")
				fmt.Println(tx.Hash())              //
				fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
				fmt.Println(tx.Value().String())    // 10000000000000000
				fmt.Println(tx.Gas())               // 105000
				fmt.Println(tx.GasPrice().Uint64()) // 102000000000
				fmt.Println(tx.Nonce())             // 110644
				fmt.Println(tx.ChainId())           //
				fmt.Println(tx.Data())              // []
				fmt.Println("To: " + tx.To().Hex()) // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

				from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
				if err != nil {
					log.Println(err.Error())
					continue
				}

				fmt.Println("From: " + from.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258

				b, _ := tx.MarshalJSON()
				fmt.Printf("\n%s\n", b)
				fmt.Println("===================================================")

				record := model.TransactionRecord{
					TxHash:   tx.Hash().Hex(),
					TxMethod: tx.Type(),
					TxBlock:  block.Number().Uint64(),
					TxAt:     block.ReceivedAt.Unix(),
					TxFrom:   from.Hex(),
					TxTo:     tx.To().Hex(),
					TxValue:  tx.Value().Int64(),
				}

				fmt.Printf("\n%+v\n", record)

				t.db.InsertTransactionRecord(record)

			}

		}
	}
}
