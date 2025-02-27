package main

import (
	"context"
	"log"
	"os"

	"github.com/aetherlink/autodrive/autodrive"
	"github.com/urfave/cli/v3"
)

var uploadCmd = &cli.Command{
	Name:  "upload",
	Usage: "Upload a file to the Auto-Drive service",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "rpc-url",
			Value: "https://demo.auto-drive.autonomys.xyz/api",
			Usage: "The RPC URL for the Auto-Drive service",
		},
		&cli.StringFlag{
			Name:    "token",
			Usage:   "The API token for the Auto-Drive service",
			Sources: cli.EnvVars("TAURUS_API_TOKEN"),
		},
		&cli.StringFlag{
			Name:    "path",
			Usage:   "The path to the file to be uploaded",
			Aliases: []string{"p"},
		},
		&cli.StringFlag{
			Name:  "filename",
			Usage: "The name of the file to be uploaded",
		},
	},
	Action: upload,
}

func upload(ctx context.Context, cmd *cli.Command) error {
	client := autodrive.NewClient(ctx, cmd.String("rpc-url"), cmd.String("token"))
	path := cmd.String("path")
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	cid, err := client.Upload(cmd.String("filename"), file)
	if err != nil {
		return err
	}

	log.Printf("Uploaded file with CID: %s", cid)
	return nil
}
