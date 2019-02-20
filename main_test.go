// Copyright 2018 The QOS Authors

package qmoon_cosmos_agent

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/rpc/client"
)

func TestCOSMOS(t *testing.T) {
	tmcli := client.NewHTTP("http://192.168.1.215:18080", "/websocket")
	t.Run("block", func(t *testing.T) {
		height := int64(1723)
		b, err := tmcli.Block(&height)
		assert.Nil(t, err)

		txs := b.Block.Txs
		assert.NotEqual(t, 0, len(txs))

		for _, v := range txs {
			ParseTx(v)
		}
	})
}
