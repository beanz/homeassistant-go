package types

type DeviceTracker struct {
	Availability []Availability `json:"availability,omitempty"`
	AvailabilityMode string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	Device Device `json:"device,omitempty"`
	Icon string `json:"icon,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic string `json:"json_attributes_topic,omitempty"`
	Name string `json:"name,omitempty"`
	PayloadAvailable string `json:"payload_available,omitempty"`
	PayloadHome string `json:"payload_home,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	PayloadNotHome string `json:"payload_not_home,omitempty"`
	Qos int `json:"qos,omitempty"`
	SourceType string `json:"source_type,omitempty"`
	StateTopic string `json:"state_topic"`
	UniqueID string `json:"unique_id,omitempty"`
	ValueTemplate string `json:"value_template,omitempty"`
}

