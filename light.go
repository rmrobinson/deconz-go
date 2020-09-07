package deconz

// Light contains the fields of a light.
type Light struct {
	// ID contains the gateway-specified ID; could change.
	// Exists only for accessing by path; dedup using UniqueID instead
	ID              string
	CTMax           int        `json:"ctmax"`
	CTMin           int        `json:"ctmin"`
	LastAnnounced   string     `json:"lastannounced"`
	LastSeen        string     `json:"lastseen"`
	ETag            string     `json:"etag"`
	Manufacturer    string     `json:"manufacturer"`
	Name            string     `json:"name"`
	ModelID         string     `json:"modelid"`
	SoftwareVersion string     `json:"swversion"`
	Type            string     `json:"type"`
	State           LightState `json:"state"`
	UniqueID        string     `json:"uniqueid"`
}

// LightState contains the specific, controllable fields of a light.
type LightState struct {
	On         bool      `json:"on"`
	Brightness int       `json:"bri"`
	Hue        int       `json:"hue"`
	Saturation int       `json:"sat"`
	CT         int       `json:"ct"`
	XY         []float64 `json:"xy"`
	Alert      string    `json:"alert"`
	ColorMode  string    `json:"colormode"`
	Effect     string    `json:"effect"`
	Reachable  bool      `json:"reachable"`
}

// GetLightsResponse contains the result of all active lights.
type GetLightsResponse map[string]Light

// SetLightStateRequest lets a user update certain properties of the light.
// These are directly changing the active light and what it is showing.
type SetLightStateRequest struct {
	On         bool      `json:"on"`
	Brightness int       `json:"bri,omitempty"`
	Hue        int       `json:"hue,omitempty"`
	Saturation int       `json:"sat,omitempty"`
	CT         int       `json:"ct,omitempty"`
	XY         []float64 `json:"xy,omitempty"`
	Alert      string    `json:"alert,omitempty"`
	// Effect contains the light effect to apply. Either 'none' or 'colorloop'
	Effect string `json:"effect,omitempty"`
	// ColorLoopSpeed contains the speed of a colorloop.
	// 1 is very fast, 255 is very slow.
	// This is only read if the 'colorloop' effect is specifed
	ColorLoopSpeed int `json:"colorloopspeed,omitempty"`
	// TransitionTime is represented in 1/10th of a second between states
	TransitionTime int `json:"transitiontime,omitempty"`
}

// SetLightConfigRequest lets a user update certain properties of the light.
// This is metadata and not directly changing the light behaviour.
type SetLightConfigRequest struct {
	Name string `json:"name,omitempty"`
}
