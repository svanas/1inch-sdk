package client

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/svanas/1inch-sdk/golang/client/onchain"
	"github.com/svanas/1inch-sdk/golang/client/orderbook"
	"github.com/svanas/1inch-sdk/golang/client/tenderly"
	"github.com/svanas/1inch-sdk/golang/helpers"
	"github.com/svanas/1inch-sdk/golang/helpers/consts/addresses"
	"github.com/svanas/1inch-sdk/golang/helpers/consts/amounts"
	"github.com/svanas/1inch-sdk/golang/helpers/consts/chains"
	"github.com/svanas/1inch-sdk/golang/helpers/consts/tokens"
	"github.com/svanas/1inch-sdk/golang/helpers/consts/web3providers"
)

func TestCreateOrderE2E(t *testing.T) {

	testcases := []struct {
		description       string
		config            Config
		createOrderParams orderbook.CreateOrderParams
		expectedOutput    string
	}{
		{
			description: "Arbitrum - Create limit order offering 1 FRAX for 1 DAI",
			config: Config{
				DevPortalApiKey: os.Getenv("DEV_PORTAL_TOKEN"),
				Web3HttpProviders: []Web3ProviderConfig{
					{
						ChainId: chains.Arbitrum,
						Url:     web3providers.Arbitrum,
					},
				},
			},
			createOrderParams: orderbook.CreateOrderParams{
				ChainId:      chains.Arbitrum,
				PrivateKey:   os.Getenv("WALLET_KEY_EMPTY"),
				Maker:        os.Getenv("WALLET_ADDRESS_EMPTY"),
				MakerAsset:   tokens.ArbitrumFrax,
				TakerAsset:   tokens.ArbitrumDai,
				MakingAmount: amounts.Ten18,
				TakingAmount: amounts.Ten18,
				Taker:        addresses.Zero,
				SkipWarnings: true,
			},
		},
		{
			description: "Polygon - Create limit order offering 1 FRAX for 1 DAI",
			config: Config{
				DevPortalApiKey: os.Getenv("DEV_PORTAL_TOKEN"),
				Web3HttpProviders: []Web3ProviderConfig{
					{
						ChainId: chains.Polygon,
						Url:     os.Getenv("WEB_3_HTTP_PROVIDER_URL_WITH_KEY_POLYGON"),
					},
				},
			},
			createOrderParams: orderbook.CreateOrderParams{
				ChainId:      chains.Polygon,
				PrivateKey:   os.Getenv("WALLET_KEY_EMPTY"),
				Maker:        os.Getenv("WALLET_ADDRESS_EMPTY"),
				MakerAsset:   tokens.PolygonFrax,
				TakerAsset:   tokens.PolygonDai,
				MakingAmount: amounts.Ten18,
				TakingAmount: amounts.Ten18,
				Taker:        addresses.Zero,
				SkipWarnings: true,
			},
		},
		{
			description: "Ethereum - Create limit order offering 1 1INCH for 1 DAI",
			config: Config{
				DevPortalApiKey: os.Getenv("DEV_PORTAL_TOKEN"),
				Web3HttpProviders: []Web3ProviderConfig{
					{
						ChainId: chains.Ethereum,
						Url:     os.Getenv("WEB_3_HTTP_PROVIDER_URL_WITH_KEY"),
					},
				},
			},
			createOrderParams: orderbook.CreateOrderParams{
				ChainId:      chains.Ethereum,
				PrivateKey:   os.Getenv("WALLET_KEY_EMPTY"),
				Maker:        os.Getenv("WALLET_ADDRESS_EMPTY"),
				MakerAsset:   tokens.Ethereum1inch,
				TakerAsset:   tokens.EthereumDai,
				MakingAmount: amounts.Ten18,
				TakingAmount: amounts.Ten18,
				Taker:        addresses.Zero,
				SkipWarnings: true,
				ApprovalType: onchain.PermitAlways,
			},
		},
		{
			description: "BSC - Create limit order offering 1 USDC for 1 DAI",
			config: Config{
				DevPortalApiKey: os.Getenv("DEV_PORTAL_TOKEN"),
				Web3HttpProviders: []Web3ProviderConfig{
					{
						ChainId: chains.Bsc,
						Url:     web3providers.Bsc,
					},
				},
			},
			createOrderParams: orderbook.CreateOrderParams{
				ApprovalType: onchain.PermitAlways,
				ChainId:      chains.Bsc,
				PrivateKey:   os.Getenv("WALLET_KEY_EMPTY"),
				Maker:        os.Getenv("WALLET_ADDRESS_EMPTY"),
				MakerAsset:   tokens.BscFrax,
				TakerAsset:   tokens.BscDai,
				MakingAmount: amounts.Ten18,
				TakingAmount: amounts.Ten18,
				Taker:        addresses.Zero,
				SkipWarnings: true,
			},
		},
	}

	//TODO set this up to have some form of configurations that enable the tests to run onchain and should also cleanup any previous test runs
	tenderlyApiKey := os.Getenv("TENDERLY_API_KEY")

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {

			t.Cleanup(func() {
				helpers.Sleep()
			})

			c, err := NewClient(tc.config)
			require.NoError(t, err)

			ctx := context.Background()
			if tenderlyApiKey != "" {
				ctx = context.WithValue(ctx, tenderly.SwapConfigKey, tenderly.SimulationConfig{
					TenderlyApiKey: tenderlyApiKey,
				})
			}
			_, _, err = c.Orderbook.CreateOrder(ctx, tc.createOrderParams)
			require.NoError(t, err)
		})
	}
}
