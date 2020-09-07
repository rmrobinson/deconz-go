package deconz

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// GatewayState contains the current state of the gateway
type GatewayState struct {
	APIVersion          string              `json:"apiversion"`
	SoftwareVersion     string              `json:"swversion"`
	SoftwareUpdateState SoftwareUpdateState `json:"swupdate"`

	MACAddress    string `json:"mac"`
	ZigbeeChannel int    `json:"zigbeechannel"`
	ZigbeePANID   int    `json:"panid"`
	GatewayID     string `json:"uuid"`

	WebsocketNotifyAll bool `json:"websocketnotifyall"`
	WebsocketPort      int  `json:"websocketport"`
	LinkButtonPressed  bool `json:"linkbutton"`

	Name       string `json:"name"`
	LocalTime  string `json:"localtime"`
	UTCTime    string `json:"utc"`
	TimeFormat string `json:"timeformat"`
	Timezone   string `json:"timezone"`

	UsingDHCP bool   `json:"dhcp"`
	GatewayIP string `json:"gateway"`
	IP        string `json:"ipaddress"`
	Netmask   string `json:"netmask"`
}

// SoftwareUpdateState contains the important data about the current software update profile of the gateway
type SoftwareUpdateState struct {
	Notify      bool   `json:"notify"`
	Text        string `json:"text"`
	UpdateState int    `json:"updatestate"`
	URL         string `json:"url"`
}

// GetGatewayResponse contains the returned data from the full gateway API call
type GetGatewayResponse struct {
	GatewayState GatewayState `json:"config"`
	//Groups map[string]Group `json:"groups`
	Lights GetLightsResponse `json:"lights"`
	// Rules map[string]Rule `json:"rules"`
	// Schedules map[string]Schedule `json:"schedules"`
	Sensors map[int]Sensor `json:"sensors"`
}

// SetConfigRequest contains the set of possible gateway configuration parameters.
type SetConfigRequest struct {
	Name        string `json:"name,omitempty"`
	RFConnected bool   `json:"rfconnected,omitempty"`
	// UpdateChannel can be set to one of stable, alpha, beta
	UpdateChannel string `json:"updatechannel,omitempty"`
	// PermitJoin when set to 0 indicates no Zigbee devices can join
	// 255 means the network is open
	// 1..254 represents the time in seconds the network will be open
	// These values decrement automatically
	PermitJoin int `json:"permitjoin,omitempty"`
	// GroupDelay contains the time between two group commands, in milliseconds
	GroupDelay        int  `json:"groupdelay,omitempty"`
	OTAUActive        bool `json:"otauactive,omitempty"`
	GWDiscoveryActive bool `json:"discovery,omitempty"`
	// Unlock being set to a value > 0 (and less than 600, the max) indicates the number of seconds the gateway is open for pairing
	Unlock int `json:"unlock,omitempty"`
	// ZigbeeChannel specifies one of 11, 15, 20 or 25 (the valid Zigbee channel numbers)
	ZigbeeChannel int    `json:"zigbeechannel,omitempty"`
	Timezone      string `json:"timezone,omitempty"`
	UTC           string `json:"utc,omitempty"`
	// TimeFormat is specified as either 12h or 24h
	TimeFormat string `json:"timeformat,omitempty"`
}

// Response is a generic response returned by the API
type Response []ResponseEntry

// ResponseEntry is one of the multiple response entries returned by the API
type ResponseEntry struct {
	Success map[string]string `json:"success"`
}

// WebsocketUpdate contains the data deserialized from the async channel
type WebsocketUpdate struct {
	Meta WebsocketUpdateMetadata

	// These are conditionally filled in by parsing the State json.RawMessage field
	GroupState  *GroupState
	LightState  *LightState
	SensorState *SensorState

	// These are conditionally filled in by parsing the relevant json.RawMessage field
	Group  *Group
	Light  *Light
	Sensor *Sensor
}

// WebsocketUpdateMetadata contains the common metadata fields about the update.
type WebsocketUpdateMetadata struct {
	Type       string `json:"t"`
	Event      string `json:"e"`
	Resource   string `json:"r"`
	ResourceID string `json:"id"`
	UniqueID   string `json:"uniqueid"`

	// The following are set on `changed` events
	Config json.RawMessage `json:"config"`
	Name   string          `json:"name"`
	State  json.RawMessage `json:"state"`

	// The following fields are only set on `scene-called` events
	GroupID string `json:"gid"`
	SceneID string `json:"scid"`

	// The following fields are set on the `added` event for the relevant resource type
	Group  json.RawMessage `json:"group"`
	Light  json.RawMessage `json:"light"`
	Sensor json.RawMessage `json:"sensor"`
}

// UnmarshalJSON allows us to conditionally deserialize the websocket update
// so that only the relevant fields are available.
func (wsu *WebsocketUpdate) UnmarshalJSON(b []byte) error {
	meta := WebsocketUpdateMetadata{}
	err := json.Unmarshal(b, &meta)
	if err != nil {
		return err
	}

	wsu.Meta = meta

	if meta.Resource == "sensors" {
		if meta.Event == "changed" {
			state := &SensorState{}
			err = json.Unmarshal(meta.State, state)
			if err != nil {
				return err
			}

			wsu.SensorState = state
		} else if meta.Event == "added" {
			sensor := &Sensor{}
			err = json.Unmarshal(meta.Sensor, sensor)
			if err != nil {
				return err
			}

			wsu.Sensor = sensor
		}
	} else if meta.Resource == "lights" {
		if meta.Event == "changed" {
			state := &LightState{}
			err = json.Unmarshal(meta.State, state)
			if err != nil {
				return err
			}

			wsu.LightState = state
		} else if meta.Event == "added" {
			light := &Light{}
			err = json.Unmarshal(meta.Light, light)
			if err != nil {
				return err
			}

			wsu.Light = light
		}
	} else if meta.Resource == "groups" {
		if meta.Event == "changed" {
			state := &GroupState{}
			err = json.Unmarshal(meta.State, state)
			if err != nil {
				return err
			}

			wsu.GroupState = state
		} else if meta.Event == "added" {
			group := &Group{}
			err = json.Unmarshal(meta.Group, group)
			if err != nil {
				return err
			}

			wsu.Group = group
		}
	}

	return nil
}

// Client represents a handle to the deconz API
type Client struct {
	httpClient *http.Client
	apiKey     string
	hostname   string
	port       int
}

// NewClient creates a new deconz API client
func NewClient(httpClient *http.Client, hostname string, port int, apiKey string) *Client {
	return &Client{
		httpClient: httpClient,
		hostname:   hostname,
		port:       port,
		apiKey:     apiKey,
	}
}

// GetGateway collects the current state from the gateway and returns it
func (c *Client) GetGateway(ctx context.Context) (*GatewayState, error) {
	r, err := http.NewRequest("GET", c.getURLBase()+"config", nil)
	if err != nil {
		return nil, err
	}

	r = r.WithContext(ctx)

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gwState := &GatewayState{}

	err = json.NewDecoder(resp.Body).Decode(gwState)
	if err != nil {
		return nil, err
	}

	return gwState, nil
}

func (c *Client) getURLBase() string {
	return "http://" + c.hostname + ":" + strconv.Itoa(c.port) + "/api/" + c.apiKey + "/"
}
