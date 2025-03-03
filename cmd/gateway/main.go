package main

import (
	"context"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v3"

	"github.com/aetherlink/aetherlink/gateway"
)

func main() {
	cmd := &cli.Command{
		Name:  "aetherlink gateway",
		Usage: "A gateway for the AetherLink service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "auto-drive-url",
				Value: "https://demo.auto-drive.autonomys.xyz/api",
				Usage: "The RPC URL for the Auto-Drive service",
			},
			&cli.StringFlag{
				Name:    "auto-drive-token",
				Usage:   "The API token for the Auto-Drive service",
				Sources: cli.EnvVars("TAURUS_API_TOKEN"),
			},
			&cli.StringFlag{
				Name:  "auto-evm-url",
				Value: "wss://auto-evm.taurus.autonomys.xyz/ws",
				Usage: "Ethereum Web socket RPC URL",
			},
			&cli.StringFlag{
				Name:    "contract-address",
				Value:   "0xbdF673bd60232917Ce960AD268a8bF6441CeFDdD",
				Aliases: []string{"ca"},
				Usage:   "Ethereum contract address",
			},
			&cli.UintFlag{
				Name:  "port",
				Value: 8081,
				Usage: "Port to listen on",
			},
			&cli.StringFlag{
				Name:     "db-path",
				Usage:    "Path to the SQLite database",
				Required: true,
			},
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			if err := godotenv.Load(); err != nil {
				return nil, errors.Wrapf(err, "failed to load .env file")
			}

			return nil, nil
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			autoDriveToken := cmd.String("auto-drive-token")
			if autoDriveToken == "" {
				autoDriveToken = os.Getenv("TAURUS_API_TOKEN")
			}

			contractAddr := common.HexToAddress(cmd.String("contract-address"))
			server, err := gateway.NewServer(
				ctx,
				cmd.Uint("port"),
				contractAddr,
				cmd.String("auto-evm-url"),
				cmd.String("auto-drive-url"),
				autoDriveToken,
				cmd.String("db-path"),
			)
			if err != nil {
				return errors.Wrap(err, "failed to create server")
			}

			return server.Start()
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
