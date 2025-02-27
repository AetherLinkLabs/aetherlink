package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "aetherlink",
		Usage: "A command-line tool to interact with the Auto-Drive service",
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			if err := godotenv.Load(); err != nil {
				return nil, errors.Wrapf(err, "failed to load .env file")
			}
			return nil, nil
		},
		Commands: []*cli.Command{
			uploadCmd,
			contractCmd,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
