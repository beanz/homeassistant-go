package types

type DeviceTrigger struct {
	AutomationType string `json:"automation_type"`
	Payload string `json:"payload,omitempty"`
	Qos int `json:"qos,omitempty"`
	Topic string `json:"topic"`
	Type string `json:"type"`
	Subtype string `json:"subtype"`
	Device Device `json:"device"`
}

