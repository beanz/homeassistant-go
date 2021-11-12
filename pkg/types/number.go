package types

type Number struct {
	Availability []Availability `json:"availability,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	AvailabilityMode string `json:"availability_mode,omitempty"`
	CommandTopic string `json:"command_topic,omitempty"`
	Device Device `json:"device,omitempty"`
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	EntityCategory EntityCategory `json:"entity_category,omitempty"`
	Icon string `json:"icon,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic string `json:"json_attributes_topic,omitempty"`
	Min float32 `json:"min,omitempty"`
	Max float32 `json:"max,omitempty"`
	Name string `json:"name,omitempty"`
	Optimistic bool `json:"optimistic,omitempty"`
	PayloadReset string `json:"payload_reset,omitempty"`
	Qos int `json:"qos,omitempty"`
	Retain bool `json:"retain,omitempty"`
	StateTopic string `json:"state_topic,omitempty"`
	Step float32 `json:"step,omitempty"`
	UniqueID string `json:"unique_id,omitempty"`
	UnitOfMeasurement string `json:"unit_of_measurement,omitempty"`
	ValueTemplate string `json:"value_template,omitempty"`
}

