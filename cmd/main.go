package cmd

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var rootCmd = &cli.Command{
	Name:                  "masa",
	Version:               "v0.0.0",
	Usage:                 "Managing your journey",
	EnableShellCompletion: true,
	Suggest:               true,
}

func Execute() {
	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
