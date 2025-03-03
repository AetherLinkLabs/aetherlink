package gateway

import (
	"log/slog"
	"testing"

	"github.com/aetherlink/aetherlink/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestGetSubdomain(t *testing.T) {
	subdomain, err := getSubdomain("chat.example.com")
	require.NoError(t, err)
	require.Equal(t, "chat", subdomain)

	subdomain, err = getSubdomain("ai.chat.example.com")
	require.NoError(t, err)
	require.Equal(t, "ai.chat", subdomain)

	subdomain, err = getSubdomain("example.com")
	require.NoError(t, err)
	require.Equal(t, "", subdomain)

	subdomain, err = getSubdomain("a.b.c.example.com")
	require.NoError(t, err)
	require.Equal(t, "a.b.c", subdomain)

	subdomain, err = getSubdomain("http://aether.localhost:8081/hello.html")
	require.NoError(t, err)
	require.Equal(t, "aether", subdomain)

	subdomain, err = getSubdomain("http://aether.ai3.localhost:8081/hello.html")
	require.NoError(t, err)
	require.Equal(t, "aether.ai3", subdomain)
}

func TestSub(t *testing.T) {
	url := "wss://auto-evm.taurus.autonomys.xyz/ws"

	ethClient, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress("0xbdF673bd60232917Ce960AD268a8bF6441CeFDdD")
	aetherContract, err := contract.NewAETHER(contractAddress, ethClient)
	if err != nil {
		panic(err)
	}

	println("Subscribed to events")
	sink := make(chan *contract.AETHERFileCIDRegistered)
	sub, err := aetherContract.WatchFileCIDRegistered(nil, sink, nil)
	if err != nil {
		panic(err)
	}

	println("Subscribed to events")
	for {
		select {
		case event := <-sink:
			slog.Info("File CID registered", "name", event.User, "cid", event.Cid, "userAddress", event.FileName)
		case err := <-sub.Err():
			slog.Error("Error watching events", "error", err)
		}
	}
}
