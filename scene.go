package deconz

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

// SetSceneConfig is used to update the scene metadata
type SetSceneConfig struct {
	Name string `json:"name"`
}

// SetSceneLightConfig contains the fields which can be set on a single light in a scene
type SetSceneLightConfig struct {
	IsOn           bool      `json:"on"`
	Brightness     int       `json:"bri"`
	TransitionTime int       `json:"transitiontime"`
	XY             []float64 `json:"xy"`
}
