package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/svanas/1inch-sdk/golang/client"
	"github.com/svanas/1inch-sdk/golang/client/orderbook"
	"github.com/svanas/1inch-sdk/golang/helpers/consts/chains"
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

	// Build the config for the orders request
	limitOrdersParams := orderbook.GetAllOrdersParams{
		ChainId: 137,
		LimitOrderV3SubscribedApiControllerGetAllLimitOrdersParams: orderbook.LimitOrderV3SubscribedApiControllerGetAllLimitOrdersParams{
			Page:   1,
			Limit:  2,
			SortBy: "createDateTime",
		},
	}

	// Execute orders request
	allOrdersResponse, _, err := c.Orderbook.GetAllOrders(context.Background(), limitOrdersParams)
	if err != nil {
		log.Fatalf("Failed to get quote: %v", err)
	}

	prettyPrintOrderResponse(allOrdersResponse)
}

func prettyPrintOrderResponse(orders []orderbook.OrderResponse) {
	for _, order := range orders {
		jsonOrder, err := json.MarshalIndent(order, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling to JSON: %v", err)
		}
		fmt.Println(string(jsonOrder))
		fmt.Println("-------------------------------")
	}
}
