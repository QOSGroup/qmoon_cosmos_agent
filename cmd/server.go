// Copyright 2018 The QOS Authors

package cmd

import (
	"encoding/hex"
	"github.com/QOSGroup/qmoon_cosmos_agent/cosmos_agent/client/stake"
	"log"
	"net/http"

	"github.com/QOSGroup/qmoon_cosmos_agent/cosmos_agent"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// ServerCmd qmoon http server
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "http server",
	RunE:  server,
}

var (
	laddr string
)

func init() {
	ServerCmd.PersistentFlags().StringVar(&laddr, "laddr", "0.0.0.0:19527", "listen addr")
}

type TxQuery struct {
	Txs []string `json:"txs"`
}

func server(cmd *cobra.Command, args []string) error {
	r := gin.Default()

	r.GET("/tx", func(ctx *gin.Context) {
		remote := ctx.Query("remote")
		hash := ctx.Query("hash")

		h, err := hex.DecodeString(hash)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "invalid hash")
			return
		}
		log.Printf("remote:%+v, hash:%+v", remote, hash)

		result, err := cosmos_agent.QueryTx(remote, h)
		log.Printf("res:%+v, err:%+v", result, err)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, result)
	})
	r.POST("/txs", func(ctx *gin.Context) {
		var tq TxQuery
		if err := ctx.ShouldBindJSON(&tq); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		result, err := cosmos_agent.ParseTxs(tq.Txs)
		log.Printf("[txs] res:%+v, err:%+v", result, err)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, result)
	})
	r.GET("/stake/validators", queryValidators)
	return r.Run(laddr)
}

func queryValidators(ctx *gin.Context) {
	remote := ctx.Query("remote")
	result, err := stake.QueryValidators(remote)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
