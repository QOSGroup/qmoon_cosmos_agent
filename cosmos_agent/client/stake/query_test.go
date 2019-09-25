package stake

import (
	"github.com/QOSGroup/qmoon_cosmos_agent/cosmos_agent"
	"log"
	"testing"
)

func TestValidators(t *testing.T) {
	result, err := QueryValidators("http://47.98.253.9:36657")
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		return
	}
	bytes, err := cosmos_agent.Cdc.MarshalJSON(result)
	log.Println(string(bytes))
}
