package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v3"

	"github.com/aetherlink/autodrive/contract"
)

var contractCmd = &cli.Command{
	Name:  "contract",
	Usage: "Interact with the Ethereum contract",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "rpc-url",
			Value: "https://auto-evm.taurus.autonomys.xyz/ws",
			Usage: "Ethereum RPC URL",
		},
		&cli.StringFlag{
			Name:    "contract-address",
			Value:   "0xbdF673bd60232917Ce960AD268a8bF6441CeFDdD",
			Aliases: []string{"ca"},
			Usage:   "Ethereum contract address",
		},
	},
	Commands: []*cli.Command{
		{
			Name:  "bind",
			Usage: "Bind a file CID to a name",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "name", Usage: "Name to bind the CID to"},
				&cli.StringFlag{Name: "cid", Usage: "Content Identifier (CID) of the file"},
				&cli.StringFlag{Name: "private-key", Usage: "Ethereum private key", Sources: cli.EnvVars("TAURUS_PRIVATE_KEY")},
			},
			Action: bindCID,
		},
		{
			Name:  "register",
			Usage: "Register a name and associate it with an Ethereum address",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "name", Usage: "Name to register"},
				&cli.StringFlag{Name: "private-key", Usage: "Ethereum private key", Sources: cli.EnvVars("TAURUS_PRIVATE_KEY")},
			},
			Action: register,
		},
		{
			Name:  "resolve",
			Usage: "Resolve a name to its associated CID",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "name", Usage: "Name to resolve"},
				&cli.StringFlag{Name: "address", Usage: "Ethereum contract address"},
			},
			Action: queryCID,
		},
	},
}

func createClientAndContract(rpcURL, contractAddress string) (*ethclient.Client, *contract.AETHER, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	contractAddr := common.HexToAddress(contractAddress)
	contract, err := contract.NewAETHER(contractAddr, client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create contract instance: %v", err)
	}

	return client, contract, nil
}

func getSigner(privateKeyHex string) (common.Address, bind.SignerFn, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to decode private key: %v", err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	signer := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewCancunSigner(big.NewInt(490000))
		return types.SignTx(tx, signer, privateKey)
	}

	return fromAddress, signer, nil
}

func getTransactionOptions(ctx context.Context, client *ethclient.Client, fromAddress common.Address, signer bind.SignerFn) (*bind.TransactOpts, error) {
	nonce, err := client.NonceAt(ctx, fromAddress, nil)
	if err != nil {
		return nil, err
	}

	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas tip cap: %v", err)
	}

	return &bind.TransactOpts{
		From:      fromAddress,
		Signer:    signer,
		GasLimit:  3000000,
		GasTipCap: gasTipCap,
		Nonce:     new(big.Int).SetUint64(nonce),
	}, nil
}

func register(ctx context.Context, cmd *cli.Command) error {
	client, contract, err := createClientAndContract(cmd.String("rpc-url"), cmd.String("contract-address"))
	if err != nil {
		return err
	}

	fromAddress, signer, err := getSigner(cmd.String("private-key"))
	if err != nil {
		return err
	}

	transactOpts, err := getTransactionOptions(ctx, client, fromAddress, signer)
	if err != nil {
		return err
	}

	signedTx, err := contract.RegisterName(transactOpts, cmd.String("name"))
	if err != nil {
		return errors.Wrapf(err, "failed to register name")
	}

	txHash := signedTx.Hash()
	slog.Info("Transaction sent", "txHash", txHash.Hex())
	return nil
}

func bindCID(ctx context.Context, cmd *cli.Command) error {
	client, contract, err := createClientAndContract(cmd.String("rpc-url"), cmd.String("contract-address"))
	if err != nil {
		return err
	}

	fromAddress, signer, err := getSigner(cmd.String("private-key"))
	if err != nil {
		return err
	}

	transactOpts, err := getTransactionOptions(ctx, client, fromAddress, signer)
	if err != nil {
		return err
	}

	signedTx, err := contract.RegisterFileCID(transactOpts, cmd.String("name"), cmd.String("cid"))
	if err != nil {
		return errors.Wrapf(err, "failed to register CID")
	}

	txHash := signedTx.Hash()
	slog.Info("Transaction sent", "txHash", txHash.Hex())
	return nil
}

func queryCID(ctx context.Context, cmd *cli.Command) error {
	_, contract, err := createClientAndContract(cmd.String("rpc-url"), cmd.String("contract-address"))
	if err != nil {
		return err
	}

	cid, err := contract.GetFileCID(nil, cmd.String("name"), common.HexToAddress(cmd.String("address")))
	if err != nil {
		return fmt.Errorf("failed to query CID: %v", err)
	}

	slog.Info("Query CID", "cid", cid, "name", cmd.String("name"))
	return nil
}
