// Copyright 2018 The QOS Authors

package main

import (
	"github.com/QOSGroup/qmoon_cosmos_agent/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd moon主命令
var rootCmd = &cobra.Command{
	Use:   "qmoon_cosmos_agent",
	Short: "qmoon_cosmos_agent cli",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		return nil
	},
}

func main() {
	rootCmd.AddCommand(cmd.AgentCmd)
	rootCmd.AddCommand(cmd.ServerCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
