package main

import (
	"context"
	"fmt"
	"time"

	"github.com/machshev/masa/thoughts"
	"github.com/urfave/cli/v3"
)

var thoughtCmd = &cli.Command{
	Name:    "thought",
	Aliases: []string{"t"},
	Usage:   "initial thoughts quickly capture new ideas",
	Commands: []*cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Action: func(ctx context.Context, cmd *cli.Command) error {
				tp, err := thoughts.GetThoughtPad()
				if err != nil {
					return fmt.Errorf("failed to open the thought pad: %v", err)
				}

				tp.Load()
				tp.Add(cmd.Args().First())
				tp.Save()

				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Action: func(ctx context.Context, cmd *cli.Command) error {
				tp, err := thoughts.GetThoughtPad()
				if err != nil {
					return fmt.Errorf("failed to open the thought pad: %v", err)
				}

				tp.Load()

				fmt.Print("Saved thoughts:\n")
				thoughts := tp.GetAll()
				for i, t := range thoughts {
					fmt.Printf("%d [%s] %s\n", i, t.Created.Format(time.DateOnly), t.Text)
				}

				return nil
			},
		},
	},
}

func init() {

	rootCmd.Commands = append(rootCmd.Commands, thoughtCmd)
}
