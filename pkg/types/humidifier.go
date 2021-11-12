package types

type Humidifier struct {
	Availability []Availability `json:"availability,omitempty"`
	AvailabilityMode string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	CommandTemplate string `json:"command_template,omitempty"`
	CommandTopic string `json:"command_topic"`
	Device Device `json:"device,omitempty"`
	DeviceClass string `json:"device_class,omitempty"`
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	EntityCategory EntityCategory `json:"entity_category,omitempty"`
	Icon string `json:"icon,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic string `json:"json_attributes_topic,omitempty"`
	MaxHumidity int `json:"max_humidity,omitempty"`
	MinHumidity int `json:"min_humidity,omitempty"`
	Name string `json:"name,omitempty"`
	Optimistic bool `json:"optimistic,omitempty"`
	PayloadAvailable string `json:"payload_available,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	PayloadOff string `json:"payload_off,omitempty"`
	PayloadOn string `json:"payload_on,omitempty"`
	PayloadResetHumidity string `json:"payload_reset_humidity,omitempty"`
	PayloadResetMode string `json:"payload_reset_mode,omitempty"`
	TargetHumidityCommandTemplate string `json:"target_humidity_command_template,omitempty"`
	TargetHumidityCommandTopic string `json:"target_humidity_command_topic"`
	TargetHumidityStateTopic string `json:"target_humidity_state_topic,omitempty"`
	TargetHumidityStateTemplate string `json:"target_humidity_state_template,omitempty"`
	ModeCommandTemplate string `json:"mode_command_template,omitempty"`
	ModeCommandTopic string `json:"mode_command_topic,omitempty"`
	ModeStateTopic string `json:"mode_state_topic,omitempty"`
	ModeStateTemplate string `json:"mode_state_template,omitempty"`
	Modes []string `json:"modes,omitempty"`
	Qos int `json:"qos,omitempty"`
	Retain bool `json:"retain,omitempty"`
	StateTopic string `json:"state_topic,omitempty"`
	StateValueTemplate string `json:"state_value_template,omitempty"`
	UniqueID string `json:"unique_id,omitempty"`
}

