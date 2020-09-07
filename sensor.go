package deconz

// Sensor represents a generic sensor in a Zigbee network
type Sensor struct {
	ID               string       `json:"ep"`
	Config           SensorConfig `json:"config"`
	ETag             string       `json:"etag"`
	ManufacturerName string       `json:"manufacturername"`
	ModelID          string       `json:"modelid"`
	Mode             int          `json:"mode"`
	Name             string       `json:"name"`
	State            SensorState  `json:"state"`
	SoftwareVersion  string       `json:"swversion"`
	Type             string       `json:"type"`
	UniqueID         string       `json:"uniqueid"`
}

// SensorConfig contains the settable properties of a sensor
type SensorConfig struct {
	On           bool `json:"on"`
	Reachable    bool `json:"reachable"`
	BatteryLevel int  `json:"battery"`
}

// SensorState contains the reported, immutable properties of a sensor.
// This is a generic type which contains state for all possible Zigbee sensors.
// Specific sensor types are subclassed and exposed with only their relevant fields.
type SensorState struct {
	LastUpdated string `json:"lastupdated"`
	LowBattery  bool   `json:"lowbattery"`
	Tampered    bool   `json:"tampered"`

	Alarm             bool   `json:"alarm"`
	CarbonMonoxide    bool   `json:"carbonmonoxide"`
	Consumption       int    `json:"consumption"`
	Power             int    `json:"power"`
	Fire              bool   `json:"fire"`
	Humidity          int    `json:"humidity"`
	Lux               int    `json:"lux"`
	LightLevel        int    `json:"lightlevel"`
	Dark              bool   `json:"dark"`
	Daylight          bool   `json:"daylight"`
	Open              bool   `json:"open"`
	Current           int    `json:"current"`
	Voltage           int    `json:"voltage"`
	Presence          bool   `json:"presence"`
	ButtonEvent       int    `json:"buttonevent"`
	Gesture           int    `json:"gesture"`
	EventDuration     int    `json:"eventduration"`
	X                 int    `json:"x"`
	Y                 int    `json:"y"`
	Angle             int    `json:"angle"`
	Pressure          int    `json:"pressure"`
	Temperature       int    `json:"temperature"`
	Valve             int    `json:"valve"`
	WindowOpen        string `json:"windowopen"`
	Vibration         bool   `json:"vibration"`
	OrientationX      int    `json:"orientation_x"`
	OrientationY      int    `json:"orientation_y"`
	OrientationZ      int    `json:"orientation_z"`
	TiltAngle         int    `json:"tiltangle"`
	VibrationStrength int    `json:"vibrationstrength"`
	Water             bool   `json:"water"`
}

// GetSensorsResponse contains the set of sensors in the gateway
type GetSensorsResponse map[string]Sensor

// SetSensorConfigRequest contains the fields of a sensor which can be changed.
type SetSensorConfigRequest SensorConfig
