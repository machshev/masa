/*
 */
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/adrg/xdg"
	"github.com/urfave/cli/v3"
)

// Capture thoughts for later processing into tasks, reminders, habits &c.
type Thought struct {
	Desc string
}

type Task struct {
	Desc string
}

type Habit struct {
	Desc string
}

type Routine struct {
	Desc string
}

// Step by step instructions for performing something
type Procedure struct {
}

func main() {
	cmd := &cli.Command{
		Name:                  "masa",
		Version:               "v0.0.0",
		Usage:                 "Managing your journey",
		EnableShellCompletion: true,
		Suggest:               true,
		Commands: []*cli.Command{
			{
				Name:    "thought",
				Aliases: []string{"t"},
				Usage:   "initial thoughts quickly capture new ideas",
				Commands: []*cli.Command{
					{
						Name:    "add",
						Aliases: []string{"a"},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							dbPath, err := xdg.DataFile("masa/db.sqlite")
							if err != nil {
								log.Fatal(err)
							}

							fmt.Println("Adding", cmd.Args().First(), "to", dbPath)
							return nil
						},
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
