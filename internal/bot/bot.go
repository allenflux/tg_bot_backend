package main

import (
	"context"
	"fmt"
	"github.com/gotd/td/telegram"
)

func main() {
	// https://core.telegram.org/api/obtaining_api_id
	appID := 814757
	appHash := "6705a5e22c1d46f27defab5e019d6754"

	client := telegram.NewClient(appID, appHash, telegram.Options{})
	if err := client.Run(context.Background(), func(ctx context.Context) error {
		// It is only valid to use client while this function is not returned
		// and ctx is not cancelled.
		api := client.API()
		fmt.Println(api)
		// Now you can invoke MTProto RPC requests by calling the API.
		// ...

		// Return to close client connection and free up resources.
		return nil
	}); err != nil {
		panic(err)
	}
	// Client is closed.
}
