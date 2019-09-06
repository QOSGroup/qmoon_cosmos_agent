// Copyright 2018 The QOS Authors

package cosmos_agent

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	jsoniter "github.com/json-iterator/go"
	"github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	"github.com/tendermint/tendermint/rpc/client"
)

var cdc = MakeCOSMOSCodec()

// cosmos
func MakeCOSMOSCodec() *amino.Codec {
	var cdc = amino.NewCodec()
	bank.RegisterCodec(cdc)
	staking.RegisterCodec(cdc)
	distribution.RegisterCodec(cdc)
	slashing.RegisterCodec(cdc)
	gov.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	types.RegisterCodec(cdc)
	cryptoAmino.RegisterAmino(cdc)
	return cdc
}

type Tx struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type ResultTx struct {
	Status int    `json:"status"`
	Txs    []Tx   `json:"txs"`
	Fee    string `json:"fee"`
	IsOK   bool   `json:"isOk"`
	Err    string `json:"err"`
}

func QueryTx(remote string, hash []byte) (*ResultTx, error) {
	tmcli := client.NewHTTP(remote, "/websocket")

	tx, err := tmcli.Tx(hash, false)
	if err != nil {
		return nil, err
	}
	result, err := ParseTx(tx.Tx)
	if err != nil {
		return nil, err
	}

	result.IsOK = tx.TxResult.IsOK()

	return result, nil
}

func ParseTxs(txs []string) ([]ResultTx, error) {
	var result []ResultTx

	for _, v := range txs {
		txd, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			result = append(result, ResultTx{Status: 500, Err: err.Error()})
			continue
		}
		var tx auth.StdTx
		if err := cdc.UnmarshalBinaryLengthPrefixed(txd, &tx); err != nil {
			result = append(result, ResultTx{Status: 500, Err: err.Error()})
			continue
		}

		var rt ResultTx
		for _, v := range tx.Msgs {
			d, _ := jsoniter.Marshal(v)
			rt.Txs = append(rt.Txs, Tx{
				Type: v.Type(),
				Data: d,
			})
		}

		var fees []string
		for _, v := range tx.Fee.Amount {
			fees = append(fees, v.String())
		}
		rt.Fee = strings.Join(fees, ",")
		result = append(result, rt)
	}

	return result, nil
}

func ParseTx(t []byte) (*ResultTx, error) {
	var result ResultTx

	var tx auth.StdTx
	err := cdc.UnmarshalBinaryLengthPrefixed(t, &tx)
	for _, v := range tx.Msgs {
		d, _ := jsoniter.Marshal(v)
		result.Txs = append(result.Txs, Tx{
			Type: v.Type(),
			Data: d,
		})
	}
	var fees []string
	for _, v := range tx.Fee.Amount {
		fees = append(fees, v.String())
	}
	result.Fee = strings.Join(fees, ",")
	return &result, err
}
