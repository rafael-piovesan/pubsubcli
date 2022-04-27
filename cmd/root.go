package cmd

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
)

var (
	addr           string
	projectID      string
	topicID        string
	subscriptionID string
	msg            string
	client         *pubsub.Client
	ctx            context.Context
)

var rootCmd = &cobra.Command{
	Use:   "pubsubcli",
	Short: "Manage a local PubSub Emulator environment",
	Long:  "Manage topics and subscriptions running on a local PubSub Emulator",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		err := os.Setenv("PUBSUB_EMULATOR_HOST", addr)
		if err != nil {
			os.Exit(1)
		}

		conn, err := net.DialTimeout("tcp", addr, time.Second)
		if err != nil {
			log.Fatalf("Connecting error: %v", err)
		} else {
			conn.Close()
		}

		ctx = context.Background()
		client, err = pubsub.NewClient(ctx, projectID)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&addr,
		"addr",
		"a",
		"localhost:8432",
		"PubSub Emulator address (following the pattern <server>:<port>)",
	)
	rootCmd.PersistentFlags().StringVarP(
		&projectID,
		"proj",
		"p",
		"my-project-id",
		"the project's ID",
	)
}
