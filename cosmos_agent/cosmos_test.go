// Copyright 2018 The QOS Authors

package cosmos_agent

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/rpc/client"
)

func TestQueryTx(t *testing.T) {
	remote := "http://192.168.1.23:26657"
	tmcli := client.NewHTTP(remote, "/websocket")
	t.Run("block", func(t *testing.T) {
		height := int64(5883)
		for {
			b, err := tmcli.Block(&height)
			if err != nil {
				break
			}
			txs := b.Block.Txs
			if b.Block.NumTxs > 0 {
				log.Printf("height:%+v", height)
			}
			for _, v := range txs {
				result, err := QueryTx(remote, v.Hash())
				assert.Nil(t, err)
				d, _ := json.Marshal(result)
				log.Printf("result:%+v", string(d))
			}
			height++
		}
	})
}
