package deconz

import "encoding/json"

// CreateScheduleRequest specifies the fields to create a new schedule.
type CreateScheduleRequest struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Command     ScheduleCommand `json:"command"`
	// Status must be one of "enabled" or "disabled"
	Status     string `json:"status"`
	AutoDelete bool   `json:"autodelete"`
	// Time represents when this scheduled action will occur
	// Supported formats include:
	// - a specific date (yyyy-MM-ddThh:mm:ss)
	// - a repeated date (W[0..127]/Thh:mm:ss)
	// for repeated dates, the field is a bitmask formed by 0MTWTFSS
	// For example: 01111100 = 124 is all weekdays, 00000011 = 3 is the weekend.
	// - a timer (PThh:mm:ss)
	// - a recurring timer (R[0..99]/PThh:mm:ss)
	// the number after the R is the number of repetitions of the timer. No number means infinite.
	Time string `json:"time"`
}

// Schedule contains the fields of a single specified scheduled event.
type Schedule struct {
	// ID contains the bridge-specified ID of this schedule.
	ID          string
	AutoDelete  bool            `json:"autodelete"`
	Command     ScheduleCommand `json:"command"`
	Description string          `json:"description"`
	ETag        string          `json:"etag"`
	Name        string          `json:"name"`
	// Status is specified as one of "enabled" or "disabled"
	Status string `json:"status"`
	// Time follows the rules in the CreateScheduleRequest documentation
	Time string `json:"time"`
}

// ScheduleCommand contains the command a scheduled action will run when triggered.
type ScheduleCommand struct {
	Address string `json:"address"`
	// Method must be specified as "PUT"
	Method string          `json:"method"`
	Body   json.RawMessage `json:"body"`
}

// SetScheduleConfigRequest contains the config fields of a schedule which can be edited.
type SetScheduleConfigRequest struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Command     ScheduleCommand `json:"command"`
	Status      string          `json:"status"`
	AutoDelete  bool            `json:"autodelete"`
	Time        string          `json:"time"`
}

// GetSchedulesResponse contains the set of schedules.
type GetSchedulesResponse map[string]Schedule
