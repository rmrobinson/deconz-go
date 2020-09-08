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

		resourceID = flag.Int("id", 0, "The ID of the light to control")
	)
	flag.Parse()

	c := deconz.NewClient(&http.Client{}, *host, *port, *apiKey)

	// No resource ID, get all sensors
	if *resourceID < 1 {
		resp, err := c.GetSensors(context.Background())
		if err != nil {
			fmt.Printf("error getting sensors: %s\n", err.Error())
			return
		}
		spew.Dump(resp)
		return
	}

	resp, err := c.GetSensor(context.Background(), *resourceID)
	if err != nil {
		fmt.Printf("error getting sensor %d: %s\n", *resourceID, err.Error())
		return
	}
	spew.Dump(resp)
}
