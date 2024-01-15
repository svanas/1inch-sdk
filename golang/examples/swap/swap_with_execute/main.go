package main

import (
	"log"
	"os"

	"github.com/1inch/1inch-sdk/golang/actions"
	"github.com/1inch/1inch-sdk/golang/client"
	"github.com/1inch/1inch-sdk/golang/client/swap"
	"github.com/1inch/1inch-sdk/golang/helpers"
	"github.com/1inch/1inch-sdk/golang/helpers/consts/amounts"
	"github.com/1inch/1inch-sdk/golang/helpers/consts/chains"
	"github.com/1inch/1inch-sdk/golang/helpers/consts/tokens"
)

func main() {

	// Build the config for the client
	config := client.Config{
		DevPortalApiKey:  os.Getenv("DEV_PORTAL_TOKEN"),
		WalletKey:        os.Getenv("WALLET_KEY"),
		Web3HttpProvider: os.Getenv("WEB_3_HTTP_PROVIDER_URL_WITH_KEY_POLYGON"),
		ChainId:          chains.Polygon,
	}

	// Create the 1inch client
	c, err := client.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Build the config for the swap request
	swapParams := swap.AggregationControllerGetSwapParams{
		Src:             tokens.PolygonFrax,
		Dst:             tokens.PolygonWeth,
		From:            os.Getenv("WALLET_ADDRESS"),
		Amount:          amounts.Ten16,
		Slippage:        0.5,
		DisableEstimate: helpers.GetPtr(true),
	}

	err = actions.SwapTokens(c, swapParams, false)
	if err != nil {
		log.Fatalf("Failed to swap tokens: %v", err)
	}
}
