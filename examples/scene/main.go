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

		groupID = flag.Int("groupID", 0, "The ID of the group to manage")
		sceneID = flag.Int("sceneID", 0, "The ID of the scene to control")
		//isOn    = flag.Bool("isOn", false, "Whether to turn the light on or off")
		//hue     = flag.Int("hue", 0, "The hue to set")
		//sat     = flag.Int("sat", 0, "The saturation level to set")
		name = flag.String("name", "", "The name of the device to set")

		create      = flag.Bool("create", false, "Whether to create the specified scene")
		setConfig   = flag.Bool("setConfig", false, "Whether to set the specified config fields")
		storeState  = flag.Bool("storeState", false, "Whether to store the current state as the scene's profile")
		recallState = flag.Bool("recallState", false, "Whether to apply the scene state to the lights in the group")
		delete      = flag.Bool("delete", false, "Whether to delete the specified scene")
	)
	flag.Parse()

	c := deconz.NewClient(&http.Client{}, *host, *port, *apiKey)

	if *create {
		req := &deconz.CreateSceneRequest{
			Name: *name,
		}

		newID, err := c.CreateScene(context.Background(), *groupID, req)
		if err != nil {
			fmt.Printf("error creating scene: %s\n", err.Error())
			return
		}
		fmt.Printf("created scene %d in group %d\n", newID, groupID)
		*sceneID = newID
	}

	// No resource ID, get all groups
	if *sceneID < 1 {
		resp, err := c.GetScenes(context.Background(), *groupID)
		if err != nil {
			fmt.Printf("error getting scenes: %s\n", err.Error())
			return
		}
		spew.Dump(resp)
		return
	}

	if *storeState {
		err := c.StoreScene(context.Background(), *groupID, *sceneID)
		if err != nil {
			fmt.Printf("error storing scene %d in group %d: %s\n", *sceneID, *groupID, err.Error())
			return
		}

		fmt.Printf("store complete\n")
	}
	if *recallState {
		err := c.RecallScene(context.Background(), *groupID, *sceneID)
		if err != nil {
			fmt.Printf("error recalling scene %d in group %d: %s\n", *sceneID, *groupID, err.Error())
			return
		}

		fmt.Printf("recall complete\n")
	}
	if *setConfig {
		req := &deconz.SetSceneConfigRequest{}

		if len(*name) > 0 {
			req.Name = *name
		}

		err := c.SetSceneConfig(context.Background(), *groupID, *sceneID, req)
		if err != nil {
			fmt.Printf("error setting scene %d config in group %d: %s\n", *sceneID, *groupID, err.Error())
			return
		}

		fmt.Printf("set complete\n")
	}
	if *delete {
		err := c.DeleteScene(context.Background(), *groupID, *sceneID)
		if err != nil {
			fmt.Printf("error deleting scene %d in group %d: %s\n", *sceneID, *groupID, err.Error())
			return
		}

		fmt.Printf("delete complete\n")
	}

	resp, err := c.GetScene(context.Background(), *groupID, *sceneID)
	if err != nil {
		fmt.Printf("error getting scene %d in group %d: %s\n", *sceneID, *groupID, err.Error())
		return
	}
	spew.Dump(resp)
}
