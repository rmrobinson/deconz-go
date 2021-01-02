package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/rmrobinson/deconz-go"
)

func main() {
	var (
		host   = flag.String("host", "", "The IP or hostname of the gateway")
		port   = flag.Int("port", 80, "The port of the gateway")
		apiKey = flag.String("apiKey", "", "The API key of the gateway")

		resourceID = flag.Int("id", 0, "The ID of the light to control")
		isOn       = flag.Bool("isOn", false, "Whether to turn the light on or off")
		hue        = flag.Int("hue", 0, "The hue to set")
		sat        = flag.Int("sat", 0, "The saturation level to set")
		name       = flag.String("name", "", "The name of the device to set")

		setState  = flag.Bool("setState", false, "Whether to set the specified state fields")
		setConfig = flag.Bool("setConfig", false, "Whether to set the specified config fields")
	)
	flag.Parse()

	c := deconz.NewClient(&http.Client{}, *host, *port, *apiKey)

	// No resource ID, get all lights
	if *resourceID < 1 {
		resp, err := c.GetLights(context.Background())
		if err != nil {
			fmt.Printf("error getting lights: %s\n", err.Error())
			return
		}
		spew.Dump(resp)
		return
	}

	if *setState {
		req := &deconz.SetLightStateRequest{
			On: *isOn,
		}

		if *hue > 0 {
			req.Hue = *hue
		}
		if *sat > 0 {
			req.Saturation = *sat
		}

		err := c.SetLightState(context.Background(), strconv.Itoa(*resourceID), req)
		if err != nil {
			fmt.Printf("error setting light %d: %s\n", *resourceID, err.Error())
			return
		}
		fmt.Printf("set complete\n")
	}
	if *setConfig {
		req := &deconz.SetLightConfigRequest{
			Name: *name,
		}

		err := c.SetLightConfig(context.Background(), strconv.Itoa(*resourceID), req)
		if err != nil {
			fmt.Printf("error setting light %d config: %s\n", *resourceID, err.Error())
			return
		}
		fmt.Printf("set complete\n")
	}

	resp, err := c.GetLight(context.Background(), strconv.Itoa(*resourceID))
	if err != nil {
		fmt.Printf("error getting light %d: %s\n", *resourceID, err.Error())
		return
	}
	spew.Dump(resp)
}
