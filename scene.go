package deconz

import (
	"context"
	"errors"
	"strconv"
)

// CreateScene creates a new scene for the specified group on the gateway. The new scene ID is returned on success.
func (c *Client) CreateScene(ctx context.Context, groupID int, req *CreateSceneRequest) (int, error) {
	resp, err := c.post(ctx, "groups/"+strconv.Itoa(groupID)+"/scenes", req)
	if err != nil {
		return 0, err
	}

	if len(*resp) < 1 {
		return 0, errors.New("new scene missing success entry")
	}
	if id, ok := (*resp)[0].Success["id"]; ok {
		if strID, ok := id.(string); ok {
			return strconv.Atoi(strID)
		}
		return 0, errors.New("new scene id not string")
	}

	return 0, errors.New("new scene missing id entry")
}

// GetScenes retrieves all the scenes for the specified group available on the gatway
func (c *Client) GetScenes(ctx context.Context, groupID int) (*GetScenesResponse, error) {
	scenesResp := &GetScenesResponse{}

	err := c.get(ctx, "groups/"+strconv.Itoa(groupID)+"/scenes", scenesResp)
	if err != nil {
		return nil, err
	}

	return scenesResp, nil
}

// GetScene retrieves the specified scene in the group
func (c *Client) GetScene(ctx context.Context, groupID int, sceneID int) (*Scene, error) {
	scene := &Scene{}

	err := c.get(ctx, "groups/"+strconv.Itoa(groupID)+"/scenes/"+strconv.Itoa(sceneID), scene)
	if err != nil {
		return nil, err
	}

	return scene, nil
}

// SetSceneConfig specifies the new config of a scene in the specified group
func (c *Client) SetSceneConfig(ctx context.Context, groupID, sceneID int, newConfig *SetSceneConfigRequest) error {
	return c.put(ctx, "groups/"+strconv.Itoa(groupID)+"/scenes/"+strconv.Itoa(sceneID), newConfig)
}

// StoreScene saves the current state of the lights in the group to the supplied scene ID
func (c *Client) StoreScene(ctx context.Context, groupID, sceneID int) error {
	req := &EmptyRequest{}
	return c.put(ctx, "groups/"+strconv.Itoa(groupID)+"/scenes/"+strconv.Itoa(sceneID)+"/store", req)
}

// RecallScene applies the saved light state from the specified scene ID to the lights in the specified group ID.
// If a light is not currently on, the recall will have no effect.
func (c *Client) RecallScene(ctx context.Context, groupID, sceneID int) error {
	req := &EmptyRequest{}
	return c.put(ctx, "groups/"+strconv.Itoa(groupID)+"/scenes/"+strconv.Itoa(sceneID)+"/recall", req)
}

// SetSceneLightState specifies the new state of a light in the specified scene.
// If the light is not a member of the group the scene is linked to, this will fail.
func (c *Client) SetSceneLightState(ctx context.Context, groupID, sceneID, lightID int, newState *SetSceneLightConfigRequest) error {
	return c.put(ctx, "groups/"+strconv.Itoa(groupID)+"/scenes/"+strconv.Itoa(sceneID)+"/lights/"+strconv.Itoa(lightID)+"/state", newState)
}

// DeleteScene removes the specified scene from the gateway
func (c *Client) DeleteScene(ctx context.Context, groupID, sceneID int) error {
	return c.delete(ctx, "groups/"+strconv.Itoa(groupID)+"/scenes/"+strconv.Itoa(sceneID))
}

// Scene contains the details of a scene
type Scene struct {
	// ID contains the bridge-specified ID of this scene.
	ID     string
	Lights []struct {
		ID             string  `json:"id"`
		IsOn           bool    `json:"on"`
		Brightness     int     `json:"bri"`
		TransitionTime int     `json:"transitiontime"`
		X              float64 `json:"x"`
		Y              float64 `json:"y"`
		CT             int     `json:"ct"`
		Hue            int     `json:"hue"`
		Saturation     int     `json:"sat"`
	} `json:"lights"`
	Name string `json:"name"`
}

// GetScenesResponse contains a collection of scenes in a group
type GetScenesResponse map[string]struct {
	LightIDs []string `json:"lights"`
	Name     string   `json:"name"`
}

// CreateSceneRequest is used to create a new scene with the specified name
type CreateSceneRequest struct {
	Name string `json:"name"`
}

// SetSceneConfigRequest is used to update the scene metadata
type SetSceneConfigRequest struct {
	Name string `json:"name"`
}

// SetSceneLightConfigRequest contains the fields which can be set on a single light in a scene
type SetSceneLightConfigRequest struct {
	IsOn           bool      `json:"on"`
	Brightness     int       `json:"bri"`
	TransitionTime int       `json:"transitiontime"`
	XY             []float64 `json:"xy"`
}
