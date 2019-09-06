// Copyright 2018 The QOS Authors

package cosmos_agent

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmbech32 "github.com/tendermint/tendermint/libs/bech32"
	"github.com/tendermint/tendermint/rpc/client"
)

func TestQueryTx(t *testing.T) {
	remote := "http://192.168.1.184:26657"
	tmcli := client.NewHTTP(remote, "/websocket")
	t.Run("block", func(t *testing.T) {
		height := int64(58578)
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

func TestDecodbeh32(t *testing.T) {
	bech32Pub := "cosmosvalconspub1zcjduepqmqrghhgxs6q8t6h7hlvwygawht78surjg3e47jdjxkkx8pc5n3psylm6gl"
	_, edPubb, _ := tmbech32.DecodeAndConvert(bech32Pub)
	var pubKey2 ed25519.PubKeyEd25519
	c := amino.NewCodec()
	c.RegisterConcrete(ed25519.PubKeyEd25519{},
		ed25519.PubKeyAminoName, nil)
	//c := cdc
	c.UnmarshalBinaryBare(edPubb, &pubKey2)
	t.Logf("old byte:%v", edPubb)
	t.Logf("new byte:%v", pubKey2.Bytes())
	t.Logf("pubKey2.Address:%s", pubKey2.Address().String())

}
