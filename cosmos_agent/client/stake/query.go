package stake

import (
	"github.com/QOSGroup/qmoon_cosmos_agent/cosmos_agent"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/viper"
)

func QueryValidators(nodeURI string) (types.Validators, error) {
	viper.Set(flags.FlagNode, nodeURI)
	cliCtx := context.NewCLIContext().WithCodec(cosmos_agent.Cdc)

	resKVs, _, err := cliCtx.QuerySubspace(types.ValidatorsKey, staking.StoreKey)
	if err != nil {
		return nil, err
	}

	var validators types.Validators
	for _, kv := range resKVs {
		validators = append(validators, types.MustUnmarshalValidator(cosmos_agent.Cdc, kv.Value))
	}
	//return cliCtx.Codec.MarshalJSON(validators)
	return validators, nil
}
