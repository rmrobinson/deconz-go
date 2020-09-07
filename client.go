package deconz

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var (
	// ErrMalformedResponse is returned if the deconz response JSON isn't formatted as expected
	ErrMalformedResponse = errors.New("malformed deconz response")
)

// Response is a generic response returned by the API
type Response []ResponseEntry

// ResponseEntry is one of the multiple response entries returned by the API
type ResponseEntry struct {
	Success map[string]interface{} `json:"success"`
	Error   ResponseError          `json:"error"`
}

// ResponseError contains a general error which was detected.
type ResponseError struct {
	Type        int    `json:"type"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

// Error allows the response error to be returned as an Error compatible type.
func (re ResponseError) Error() string {
	return fmt.Sprintf("%s: %d (%s)", re.Address, re.Type, re.Description)
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

func (c *Client) getURLBase() string {
	return "http://" + c.hostname + ":" + strconv.Itoa(c.port) + "/api/" + c.apiKey + "/"
}

func (c *Client) get(ctx context.Context, path string, respType interface{}) error {
	r, err := http.NewRequest(http.MethodGet, c.getURLBase()+path, nil)
	if err != nil {
		return err
	}

	r = r.WithContext(ctx)

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return json.NewDecoder(resp.Body).Decode(respType)
	}

	deconzResp := Response{}
	err = json.NewDecoder(resp.Body).Decode(&deconzResp)
	if err != nil {
		return err
	}

	if len(deconzResp) < 1 {
		return ErrMalformedResponse
	}

	return deconzResp[0].Error
}

func (c *Client) put(ctx context.Context, path string, reqType interface{}) error {
	req, err := json.Marshal(reqType)
	if err != nil {
		return err
	}

	r, err := http.NewRequest(http.MethodPut, c.getURLBase()+path, bytes.NewBuffer(req))
	if err != nil {
		return err
	}

	r = r.WithContext(ctx)

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	deconzResp := Response{}
	err = json.NewDecoder(resp.Body).Decode(&deconzResp)
	if err != nil {
		return err
	}

	if len(deconzResp) < 1 {
		return ErrMalformedResponse
	}
	for _, deconsRespEntry := range deconzResp {
		if len(deconsRespEntry.Success) < 1 {
			return deconsRespEntry.Error
		}
	}

	return nil
}

func (c *Client) delete(ctx context.Context, path string) error {
	r, err := http.NewRequest(http.MethodDelete, c.getURLBase()+path, nil)
	if err != nil {
		return err
	}

	r = r.WithContext(ctx)

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	deconzResp := Response{}
	err = json.NewDecoder(resp.Body).Decode(&deconzResp)
	if err != nil {
		return err
	}

	if len(deconzResp) < 1 {
		return ErrMalformedResponse
	}
	for _, deconsRespEntry := range deconzResp {
		if len(deconsRespEntry.Success) < 1 {
			return deconsRespEntry.Error
		}
	}

	return nil
}

// GetGateway collects the current state from the gateway and returns it
func (c *Client) GetGateway(ctx context.Context) (*GatewayState, error) {
	gwState := &GatewayState{}

	err := c.get(ctx, "config", gwState)
	if err != nil {
		return nil, err
	}

	return gwState, nil
}

// GetLights retrieves all the lights available on the gatway
func (c *Client) GetLights(ctx context.Context) (*GetLightsResponse, error) {
	lightsResp := &GetLightsResponse{}

	err := c.get(ctx, "lights", lightsResp)
	if err != nil {
		return nil, err
	}

	return lightsResp, nil
}

// GetLight retrieves the specified light
func (c *Client) GetLight(ctx context.Context, id int) (*Light, error) {
	light := &Light{}

	err := c.get(ctx, "lights/"+strconv.Itoa(id), light)
	if err != nil {
		return nil, err
	}

	return light, nil
}

// SetLightState specifies the new state of a light
func (c *Client) SetLightState(ctx context.Context, id int, newState *SetLightStateRequest) error {
	return c.put(ctx, "lights/"+strconv.Itoa(id)+"/state", newState)
}

// SetLightConfig specifies the new config of a light
func (c *Client) SetLightConfig(ctx context.Context, id int, newConfig *SetLightConfigRequest) error {
	return c.put(ctx, "lights/"+strconv.Itoa(id), newConfig)
}

// DeleteLight removes the specified light from the gateway
func (c *Client) DeleteLight(ctx context.Context, id int) error {
	return c.delete(ctx, "lights/"+strconv.Itoa(id))
}

// DeleteLightGroups removes the light from all its groups
func (c *Client) DeleteLightGroups(ctx context.Context, id int) error {
	return c.delete(ctx, "lights/"+strconv.Itoa(id)+"/groups")
}

// DeleteLightScenes removes the light from all its scenes
func (c *Client) DeleteLightScenes(ctx context.Context, id int) error {
	return c.delete(ctx, "lights/"+strconv.Itoa(id)+"/scenes")
}
