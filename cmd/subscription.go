package cmd

import (
	"github.com/spf13/cobra"
)

var subscriptionCmd = &cobra.Command{
	Use:   "subscription",
	Short: "Manage subscriptions on PubSub Emulator",
	Long:  "Manage subscriptions on PubSub Emulator",
}

func init() {
	rootCmd.AddCommand(subscriptionCmd)
}
