package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/adrg/xdg"
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
				dbPath, err := xdg.DataFile("masa/db.sqlite")
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Adding", cmd.Args().First(), "to", dbPath)
				return nil
			},
		},
	},
}

func init() {

	rootCmd.Commands = append(rootCmd.Commands, thoughtCmd)
}
