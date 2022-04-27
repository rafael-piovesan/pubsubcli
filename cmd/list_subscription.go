package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

var listSubscriptionCmd = &cobra.Command{
	Use:   "list",
	Short: "List subscriptions",
	Long:  "List all subscriptions.",
	Run: func(cmd *cobra.Command, args []string) {
		i := client.Subscriptions(ctx)
		s, err := i.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				log.Print("No subscriptions found")
				return
			}

			log.Fatal(err)
		}

		for !errors.Is(err, iterator.Done) {
			log.Printf("id: [%v]\t\tdetails:[%v]", s.ID(), s.String())
			s, err = i.Next()
		}
	},
}

func init() {
	subscriptionCmd.AddCommand(listSubscriptionCmd)
}
