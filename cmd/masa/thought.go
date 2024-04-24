package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/machshev/masa/config"
	pb "github.com/machshev/masa/pb/thoughts"
	"github.com/machshev/masa/pb/thoughts/thoughtsconnect"
	"github.com/machshev/masa/thoughts"
	"github.com/urfave/cli/v3"
)

func getMasaClient() thoughtsconnect.ThoughtServiceClient {
	return thoughtsconnect.NewThoughtServiceClient(
		http.DefaultClient,
		"http://"+config.GetMasaServerURL(),
	)

}

var thoughtCmd = &cli.Command{
	Name:    "thought",
	Aliases: []string{"t"},
	Usage:   "initial thoughts quickly capture new ideas",
	Commands: []*cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Action: func(ctx context.Context, cmd *cli.Command) error {
				client := getMasaClient()
				res, err := client.Add(
					context.Background(),
					connect.NewRequest(&pb.AddRequest{Thought: cmd.Args().First()}),
				)
				if err != nil {
					log.Fatalf("%v", err)
				}

				id, err := uuid.FromBytes(res.Msg.Uuid)
				if err != nil {
					log.Fatalf("UUID returned could not be parsed: %v", err)
				}
				fmt.Printf("Thought added with UUID: %s\n", id)

				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Action: func(ctx context.Context, cmd *cli.Command) error {
				client := getMasaClient()
				res, err := client.List(
					context.Background(),
					connect.NewRequest(&pb.ListRequest{}),
				)
				if err != nil {
					log.Fatalf("%v", err)
				}

				fmt.Print("Saved thoughts:\n")
				for _, pb_t := range res.Msg.Thoughts {
					t, err := thoughts.ThoughtFromPB(pb_t)
					if err != nil {
						log.Fatalf("Failed to parse thought (%v): %v", pb_t, err)
					}

					fmt.Printf("%s [%s] %s\n", t.Id, t.Created.Format(time.DateOnly), t.Text)
				}

				return nil
			},
		},
	},
}

func init() {

	rootCmd.Commands = append(rootCmd.Commands, thoughtCmd)
}
