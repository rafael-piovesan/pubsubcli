package cmd

import (
	"github.com/spf13/cobra"
)

var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "Manage topics on PubSub Emulator",
	Long:  "Manage topics on PubSub Emulator.",
}

func init() {
	rootCmd.AddCommand(topicCmd)
}
