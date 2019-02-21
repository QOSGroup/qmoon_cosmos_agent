// Copyright 2018 The QOS Authors

package cmd

import (
	"github.com/spf13/cobra"
)

// AgentCmd qmoon http agent
var AgentCmd = &cobra.Command{
	Use:   "agent",
	Short: "cosmos agent",
	RunE:  agent,
}

var (
	qmoon  string
	cosmos string
)

func init() {
	AgentCmd.PersistentFlags().StringVar(&qmoon, "qmoon", "http://localhost:9527", "the remote of qmoon")
	AgentCmd.PersistentFlags().StringVar(&cosmos, "cosmos", "http://localhost:26657", "the remote of cosmos")
}

func agent(cmd *cobra.Command, args []string) error {
	return nil
}
