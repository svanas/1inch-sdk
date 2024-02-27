package main

import (
	"context"
	"log"
	"os"

	"github.com/svanas/1inch-sdk/golang/client"
	"github.com/svanas/1inch-sdk/golang/client/onchain"
	"github.com/svanas/1inch-sdk/golang/client/swap"
	"github.com/svanas/1inch-sdk/golang/helpers/consts/amounts"
	"github.com/svanas/1inch-sdk/golang/helpers/consts/chains"
	"github.com/svanas/1inch-sdk/golang/helpers/consts/tokens"
)

func main() {

	// Build the config for the client
	config := client.Config{
		DevPortalApiKey: os.Getenv("DEV_PORTAL_TOKEN"),
		Web3HttpProviders: []client.Web3ProviderConfig{
			{
				ChainId: chains.Polygon,
				Url:     os.Getenv("WEB_3_HTTP_PROVIDER_URL_WITH_KEY_POLYGON"),
			},
		},
	}

	// Create the 1inch client
	c, err := client.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Build the config for the swap request
	swapParams := swap.SwapTokensParams{
		ApprovalType:  onchain.PermitIfPossible,
		SkipWarnings:  false,
		PublicAddress: os.Getenv("WALLET_ADDRESS"),
		WalletKey:     os.Getenv("WALLET_KEY"),
		ChainId:       chains.Polygon,
		AggregationControllerGetSwapParams: swap.AggregationControllerGetSwapParams{
			Src:             tokens.PolygonFrax,
			Dst:             tokens.PolygonUsdc,
			From:            os.Getenv("WALLET_ADDRESS"),
			Amount:          amounts.Ten16,
			Slippage:        0.5,
			DisableEstimate: true,
		},
	}

	err = c.Actions.SwapTokens(context.Background(), swapParams)
	if err != nil {
		log.Fatalf("Failed to swap tokens: %v", err)
	}
}
