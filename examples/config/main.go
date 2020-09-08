package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrobinson/deconz-go"
)

func main() {
	var (
		host   = flag.String("host", "", "The IP or hostname of the gateway")
		port   = flag.Int("port", 80, "The port of the gateway")
		apiKey = flag.String("apiKey", "", "The API key of the gateway")

		create  = flag.Bool("create", false, "Whether to create a new API key or not")
		appName = flag.String("appName", "", "The name of the app to specify when creating an API key")
		delete  = flag.Bool("delete", false, "Whether to delete the other API key")
		delKey  = flag.String("delKey", "", "The API key to delete")
	)
	flag.Parse()

	c := deconz.NewClient(&http.Client{}, *host, *port, *apiKey)

	if *create {
		req := &deconz.CreateAPIKeyRequest{
			ApplicationName: *appName,
		}

		key, err := c.CreateAPIKey(context.Background(), req)
		if err != nil {
			fmt.Printf("error creating API key: %s\n", err.Error())
			return
		}

		fmt.Printf("created new API key %s\n", key)
		*apiKey = key
	}
	if *delete {
		err := c.DeleteAPIKey(context.Background(), *delKey)
		if err != nil {
			fmt.Printf("error deleting API key: %s\n", err.Error())
			return
		}

		fmt.Printf("deleted api key\n")
	}

	gw, err := c.GetGatewayState(context.Background())
	if err != nil {
		fmt.Printf("err getting gateway\n")
		return
	}

	spew.Dump(gw)
}
