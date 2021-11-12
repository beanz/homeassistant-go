package types

type Cover struct {
	Availability []Availability `json:"availability,omitempty"`
	AvailabilityMode string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	CommandTopic string `json:"command_topic,omitempty"`
	Device Device `json:"device,omitempty"`
	DeviceClass string `json:"device_class,omitempty"`
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	EntityCategory EntityCategory `json:"entity_category,omitempty"`
	Icon string `json:"icon,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic string `json:"json_attributes_topic,omitempty"`
	Name string `json:"name,omitempty"`
	Optimistic bool `json:"optimistic,omitempty"`
	PayloadAvailable string `json:"payload_available,omitempty"`
	PayloadClose string `json:"payload_close,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	PayloadOpen string `json:"payload_open,omitempty"`
	PayloadStop string `json:"payload_stop,omitempty"`
	PositionClosed int `json:"position_closed,omitempty"`
	PositionOpen int `json:"position_open,omitempty"`
	PositionTemplate string `json:"position_template,omitempty"`
	PositionTopic string `json:"position_topic,omitempty"`
	Qos int `json:"qos,omitempty"`
	Retain bool `json:"retain,omitempty"`
	SetPositionTemplate string `json:"set_position_template,omitempty"`
	SetPositionTopic string `json:"set_position_topic,omitempty"`
	StateClosed string `json:"state_closed,omitempty"`
	StateClosing string `json:"state_closing,omitempty"`
	StateOpen string `json:"state_open,omitempty"`
	StateOpening string `json:"state_opening,omitempty"`
	StateStopped string `json:"state_stopped,omitempty"`
	StateTopic string `json:"state_topic,omitempty"`
	TiltClosedValue int `json:"tilt_closed_value,omitempty"`
	TiltCommandTemplate string `json:"tilt_command_template,omitempty"`
	TiltCommandTopic string `json:"tilt_command_topic,omitempty"`
	TiltMax int `json:"tilt_max,omitempty"`
	TiltMin int `json:"tilt_min,omitempty"`
	TiltOpenedValue int `json:"tilt_opened_value,omitempty"`
	TiltOptimistic bool `json:"tilt_optimistic,omitempty"`
	TiltStatusTemplate string `json:"tilt_status_template,omitempty"`
	TiltStatusTopic string `json:"tilt_status_topic,omitempty"`
	UniqueID string `json:"unique_id,omitempty"`
	ValueTemplate string `json:"value_template,omitempty"`
}

