package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

var listTopicCmd = &cobra.Command{
	Use:   "list",
	Short: "List topics",
	Long:  "List all topics.",
	Run: func(cmd *cobra.Command, args []string) {
		i := client.Topics(ctx)
		s, err := i.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				log.Print("No topics found")
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
	topicCmd.AddCommand(listTopicCmd)
}
