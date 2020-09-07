package deconz

// Group represents a collection of lights and provides the foundation for scenes
type Group struct {
	LastAction Action   `json:"action"`
	DeviceIDs  []string `json:"devicemembership"`
	ETag       string   `json:"etag"`
	Hidden     bool     `json:"hidden"`
	ID         string   `json:"id"`
	// LightIDs contains a gateway-sorted list of all the light IDs in this group
	LightIDs []string `json:"lights"`
	// LightIDSequence contains a user-sorted list of a subset of all the light IDs in this group
	LightIDSequence []string `json:"lightsequence"`
	// MultiDeviceIDs contains the subsequent IDs of multi-device lights
	MultiDeviceIDs []string `json:"multideviceids"`
	Name           string   `json:"name"`
	Scenes         []struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		TransitionTime int    `json:"transitiontime"`
		LightCount     int    `json:"lightcount"`
	} `json:"scenes"`
	State GroupState `json:"state"`
}

// GroupState contains the fields relevant to the state of a group
type GroupState struct {
	AllOn bool `json:"all_on"`
	AnyOn bool `json:"any_on"`
}

// Action represents a state change which has occurred
type Action struct {
	On                bool      `json:"on"`
	Brightness        int       `json:"bri"`
	Hue               int       `json:"hue"`
	Saturation        int       `json:"sat"`
	ColourTemperature int       `json:"ct"`
	XY                []float64 `json:"xy"`
	Effect            string    `json:"effect"`
}

// CreateGroupRequest is used to create a new group with the specified name.
type CreateGroupRequest struct {
	Name string `json:"name"`
}

// GetGroupsResponse contains the fields returned by the 'list groups' API call
type GetGroupsResponse map[string]Group

// SetGroupConfigRequest sets the config options of the group
type SetGroupConfigRequest struct {
	Name            string   `json:"name,omitempty"`
	LightIDs        []string `json:"lights,omitempty"`
	Hidden          bool     `json:"hidden,omitempty"`
	LightIDSequence []string `json:"lightsequence,omitempty"`
	MultiDeviceIDs  []string `json:"multideviceids,omitempty"`
}

// SetGroupStateRequest sets the state of the specified group.
type SetGroupStateRequest struct {
	SetLightStateRequest
	// Toggle flips the state from on to off or vice versa.
	// This superscedes the values set directly.
	Toggle bool `json:"toggle,omitempty"`
}
