package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var rootCmd = &cli.Command{
	Name:                  "masa",
	Version:               "v0.2.1",
	Usage:                 "Managing your journey",
	EnableShellCompletion: true,
	Suggest:               true,
}

func main() {
	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
