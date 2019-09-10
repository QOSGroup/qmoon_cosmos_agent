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

type StakingValidator struct {
	Commission struct {
		MaxChangeRate string `json:"max_change_rate"`
		MaxRate       string `json:"max_rate"`
		Rate          string `json:"rate"`
		UpdateTime    string `json:"update_time"`
	} `json:"commission"`
	ConsensusPubkey string `json:"consensus_pubkey"`
	DelegatorShares string `json:"delegator_shares"`
	Description     struct {
		Details  string `json:"details"`
		Identity string `json:"identity"`
		Moniker  string `json:"moniker"`
		Website  string `json:"website"`
	} `json:"description"`
	Jailed            bool   `json:"jailed"`
	MinSelfDelegation string `json:"min_self_delegation"`
	OperatorAddress   string `json:"operator_address"`
	Status            int    `json:"status"`
	Tokens            string `json:"tokens"`
	UnbondingHeight   string `json:"unbonding_height"`
	UnbondingTime     string `json:"unbonding_time"`
}

func TestValidator(t *testing.T) {
	remote := "http://192.168.1.184:26657"
	tmcli := client.NewHTTP(remote, "/websocket")
	endpoint := "/store/staking/subspace"
	key := []byte{0x21}
	res1, err := tmcli.ABCIQuery(endpoint, key)
	assert.Nil(t, err)
	var sv []StakingValidator
	if res1.Response.Code == 0 {
		json.Unmarshal(res1.Response.Value, &sv)
	}

	t.Logf("name:%+v", sv.Description.Moniker)
	//cli := lib.TendermintClient("http://192.168.1.180:26657")
	//endpoint := "custom/staking/validator"
	//s, err := hex.DecodeString(strings.ToLower("1FBF89B0DC10144877BF8C93D7FBAF00E10FFC24"))
	//assert.Nil(t, err)
	//
	//address, err := bech32.ConvertAndEncode("cosmosvaloper", s)
	//assert.Nil(t, err)
	//t.Logf("address1:%s", address)
	//
	//address = lib.PubkeyToBech32Address("cosmosvalconspub", "tendermint/PubKeyEd25519", "9tK9IT+FPdf2qm+5c2qaxi10sWP+3erWTKgftn2PaQM=")
	//t.Logf("address2:%s", address)
	//
	//p := QueryValidatorParams{ValidatorAddr: address}
	//d, err := lib.Cdc.MarshalJSON(p)
	//t.Logf("d:%s", string(d))
	//assert.Nil(t, err)
	//res, err := cli.ABCIQuery(endpoint, d)
	//assert.Nil(t, err)
	//t.Logf("res:%+v", res)
	//
	//res1, err := cli.ABCIQuery("custom/staking/validators", nil)
	//t.Logf("code:%+v", res1.Response.Code)
	//t.Logf("res1.Response.Value:%+v", string(res1.Response.Value))
	//if res1.Response.Code == 0 {
	//	var sv []StakingValidator
	//	json.Unmarshal(res1.Response.Value, &sv)
	//	//t.Logf("sv:%+v", sv)
	//}
}
