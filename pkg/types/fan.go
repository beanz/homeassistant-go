package types

type Fan struct {
	Availability []Availability `json:"availability,omitempty"`
	AvailabilityMode string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	CommandTemplate string `json:"command_template,omitempty"`
	CommandTopic string `json:"command_topic"`
	Device Device `json:"device,omitempty"`
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	EntityCategory EntityCategory `json:"entity_category,omitempty"`
	Icon string `json:"icon,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic string `json:"json_attributes_topic,omitempty"`
	Name string `json:"name,omitempty"`
	Optimistic bool `json:"optimistic,omitempty"`
	OscillationCommandTemplate string `json:"oscillation_command_template,omitempty"`
	OscillationCommandTopic string `json:"oscillation_command_topic,omitempty"`
	OscillationStateTopic string `json:"oscillation_state_topic,omitempty"`
	OscillationValueTemplate string `json:"oscillation_value_template,omitempty"`
	PayloadAvailable string `json:"payload_available,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	PayloadOff string `json:"payload_off,omitempty"`
	PayloadOn string `json:"payload_on,omitempty"`
	PayloadOscillationOff string `json:"payload_oscillation_off,omitempty"`
	PayloadOscillationOn string `json:"payload_oscillation_on,omitempty"`
	PayloadResetPercentage string `json:"payload_reset_percentage,omitempty"`
	PayloadResetPresetMode string `json:"payload_reset_preset_mode,omitempty"`
	PercentageCommandTemplate string `json:"percentage_command_template,omitempty"`
	PercentageCommandTopic string `json:"percentage_command_topic,omitempty"`
	PercentageStateTopic string `json:"percentage_state_topic,omitempty"`
	PercentageValueTemplate string `json:"percentage_value_template,omitempty"`
	PresetModeCommandTemplate string `json:"preset_mode_command_template,omitempty"`
	PresetModeCommandTopic string `json:"preset_mode_command_topic,omitempty"`
	PresetModeStateTopic string `json:"preset_mode_state_topic,omitempty"`
	PresetModeValueTemplate string `json:"preset_mode_value_template,omitempty"`
	PresetModes []string `json:"preset_modes,omitempty"`
	Qos int `json:"qos,omitempty"`
	Retain bool `json:"retain,omitempty"`
	SpeedRangeMax int `json:"speed_range_max,omitempty"`
	SpeedRangeMin int `json:"speed_range_min,omitempty"`
	StateTopic string `json:"state_topic,omitempty"`
	StateValueTemplate string `json:"state_value_template,omitempty"`
	UniqueID string `json:"unique_id,omitempty"`
}

