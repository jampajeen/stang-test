package ethereum

import (
	"context"
	"log"

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

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			for _, tx := range block.Transactions() {
				from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
				if err != nil {
					log.Println(err.Error())
					continue
				}

				record := model.TransactionRecord{
					TxHash:   tx.Hash().Hex(),
					TxMethod: tx.Type(),
					TxBlock:  block.Number().Uint64(),
					TxAt:     block.ReceivedAt.Unix(),
					TxFrom:   from.Hex(),
					TxTo:     tx.To().Hex(),
					TxValue:  tx.Value().Int64(),
				}
				t.db.InsertTransactionRecord(record)

			}
		}
	}
}
