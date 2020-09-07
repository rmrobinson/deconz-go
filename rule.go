package deconz

import "encoding/json"

// Rule contains the fields of a rule
type Rule struct {
	// ID contains the bridge-specified ID of this rule.
	ID              string
	Actions         []RuleAction    `json:"actions"`
	Conditions      []RuleCondition `json:"conditions"`
	CreatedAt       string          `json:"created"`
	ETag            string          `json:"etag"`
	LastTriggeredAt string          `json:"lasttriggered"`
	Name            string          `json:"name"`
	Owner           string          `json:"owner"`
	Periodic        int             `json:"periodic"`
	// Status can be one of "enabled", "disabled"
	Status         string `json:"status"`
	TriggeredCount int    `json:"timestriggered"`
}

// RuleAction contains a single action in a rule.
// The Body is a JSON object serialized outside this format.
type RuleAction struct {
	Address string          `json:"address,omitempty"`
	Body    json.RawMessage `json:"body,omitempty"`
	// Method can be one of PUT, POST, DELETE, BIND
	Method string `json:"method,omitempty"`
}

// RuleCondition contains a single condition in a rule.
type RuleCondition struct {
	Address string `json:"address,omitempty"`
	// Operator can be one of eq (equals), gt (greater than), lt (less than), dx (on change)
	Operator string `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`
}

// GetRulesResponse contains the data returned by a call to list the rules.
type GetRulesResponse map[string]Rule

// SetRulesRequest contains the fields which can be set on a rule.
type SetRulesRequest struct {
	Actions    []RuleAction    `json:"actions,omitempty"`
	Conditions []RuleCondition `json:"conditions,omitempty"`
	Name       string          `json:"name,omitempty"`
	Perodic    int             `json:"perodic,omitempty"`
	Status     string          `json:"status,omitempty"`
}
