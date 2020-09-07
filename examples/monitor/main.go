package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/websocket"
	"github.com/rmrobinson/deconz-go"
)

func main() {
	var (
		host   = flag.String("host", "", "The IP or hostname of the gateway")
		port   = flag.Int("port", 80, "The port of the gateway")
		apiKey = flag.String("apiKey", "", "The API key of the gateway")
	)
	flag.Parse()

	c := deconz.NewClient(&http.Client{}, *host, *port, *apiKey)

	gw, err := c.GetGateway(context.Background())
	if err != nil {
		fmt.Printf("err getting gateway\n")
		return
	}

	wsu := url.URL{
		Scheme: "ws",
		Host:   *host + ":" + strconv.Itoa(gw.WebsocketPort),
	}

	wsc, _, err := websocket.DefaultDialer.Dial(wsu.String(), nil)
	if err != nil {
		fmt.Printf("err creating websocket: %s\n", err.Error())
		return
	}
	defer wsc.Close()

	go func() {
		for {
			msg := &deconz.WebsocketUpdate{}
			err := wsc.ReadJSON(msg)

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("err reading from websocket: %s\n", err.Error())
				return
			}

			spew.Dump(msg)
		}
	}()

	time.Sleep(time.Second * 30)
}
