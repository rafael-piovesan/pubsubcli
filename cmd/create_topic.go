package cmd

import (
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
)

var createTopicCmd = &cobra.Command{
	Use:   "create",
	Short: "Create topic",
	Long:  "Create a new topic.",
	Run: func(cmd *cobra.Command, args []string) {
		topic := client.Topic(topicID)
		exists, err := topic.Exists(ctx)
		if err != nil {
			log.Fatal(err)
		}

		if !exists {
			topic, err = client.CreateTopicWithConfig(ctx, topicID, &pubsub.TopicConfig{})
			if err != nil {
				log.Fatal(err)
			} else {
				log.Print("topic created")
			}
		} else {
			log.Print("topic already exists")
		}

		// Also create Dead Letter Topic
		dltTopicID := topicID + ".dlt"
		dltTopic := client.Topic(dltTopicID)

		dltExists, err := dltTopic.Exists(ctx)
		if err != nil {
			log.Fatal(err)
		}

		if !dltExists {
			dltTopic, err = client.CreateTopicWithConfig(ctx, dltTopicID, &pubsub.TopicConfig{})
			if err != nil {
				log.Fatal(err)
			} else {
				log.Print("dlt topic created")
			}
		}
	},
}

func init() {
	topicCmd.AddCommand(createTopicCmd)

	createTopicCmd.Flags().StringVarP(&topicID, "topic", "t", "", "the topic's id")
	_ = createTopicCmd.MarkFlagRequired("topic")
}
