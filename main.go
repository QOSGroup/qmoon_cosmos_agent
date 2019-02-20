// Copyright 2018 The QOS Authors

package qmoon_cosmos_agent

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/encoding/amino"
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

func ParseTx(t []byte) error {
	var tx auth.StdTx
	err := cdc.UnmarshalBinaryLengthPrefixed(t, &tx)
	fmt.Printf("%+v", tx.Msgs)
	return err
}
