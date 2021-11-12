package types

type Scene struct {
	Availability []Availability `json:"availability,omitempty"`
	AvailabilityMode string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	CommandTopic string `json:"command_topic,omitempty"`
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	EntityCategory EntityCategory `json:"entity_category,omitempty"`
	Icon string `json:"icon,omitempty"`
	Name string `json:"name,omitempty"`
	PayloadAvailable string `json:"payload_available,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	PayloadOn string `json:"payload_on,omitempty"`
	Qos int `json:"qos,omitempty"`
	Retain bool `json:"retain,omitempty"`
	UniqueID string `json:"unique_id,omitempty"`
}

