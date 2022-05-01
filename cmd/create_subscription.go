package cmd

import (
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
)

var createSubscriptionCmd = &cobra.Command{
	Use:   "create",
	Short: "Create subscription",
	Long:  "Create a new subscription.",
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
				log.Print("dlt topic did not exist, was also created")
			}
		}

		subscription := client.Subscription(subscriptionID)
		subExists, err := subscription.Exists(ctx)
		if err != nil {
			log.Fatal(err)
		}
		if !subExists {
			subscription, err = client.CreateSubscription(
				ctx,
				subscriptionID,
				pubsub.SubscriptionConfig{
					Topic:       topic,
					AckDeadline: 20 * time.Second,
					DeadLetterPolicy: &pubsub.DeadLetterPolicy{
						DeadLetterTopic:     dltTopic.String(),
						MaxDeliveryAttempts: 5,
					},
				},
			)
			if err != nil {
				log.Fatal(err)
			} else {
				log.Print("subscription created")
			}
		} else {
			log.Print("subscription already exists")
		}
	},
}

func init() {
	subscriptionCmd.AddCommand(createSubscriptionCmd)

	createSubscriptionCmd.Flags().StringVarP(&topicID, "topic", "t", "", "the topic's id")
	createSubscriptionCmd.Flags().StringVarP(&subscriptionID, "sub", "s", "", "the subscription's id")
	_ = createSubscriptionCmd.MarkFlagRequired("topic")
	_ = createSubscriptionCmd.MarkFlagRequired("sub")
}
