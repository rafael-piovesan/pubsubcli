package cmd

import (
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
)

var publishMsgCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish messages",
	Long:  "Publish a messages to a topic on PubSub Emulator",
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
				log.Print("topic did not exist, was also created")
			}
		}

		r := topic.Publish(ctx, &pubsub.Message{Data: []byte(msg)})
		_, err = r.Get(ctx)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(publishMsgCmd)

	publishMsgCmd.Flags().StringVarP(&topicID, "topic", "t", "", "the topic's id")
	publishMsgCmd.Flags().StringVarP(&msg, "msg", "m", "", "the message to be sent")
	_ = publishMsgCmd.MarkFlagRequired("topic")
	_ = publishMsgCmd.MarkFlagRequired("msg")
}
