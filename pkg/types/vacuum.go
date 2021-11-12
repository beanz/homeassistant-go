package types

type Vacuum struct {
	Availability []Availability `json:"availability,omitempty"`
	AvailabilityMode string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	CommandTopic string `json:"command_topic,omitempty"`
	Device Device `json:"device,omitempty"`
	FanSpeedList []string `json:"fan_speed_list,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic string `json:"json_attributes_topic,omitempty"`
	Name string `json:"name,omitempty"`
	PayloadAvailable string `json:"payload_available,omitempty"`
	PayloadCleanSpot string `json:"payload_clean_spot,omitempty"`
	PayloadLocate string `json:"payload_locate,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	PayloadPause string `json:"payload_pause,omitempty"`
	PayloadReturnToBase string `json:"payload_return_to_base,omitempty"`
	PayloadStart string `json:"payload_start,omitempty"`
	PayloadStop string `json:"payload_stop,omitempty"`
	Qos int `json:"qos,omitempty"`
	Retain bool `json:"retain,omitempty"`
	Schema string `json:"schema,omitempty"`
	SendCommandTopic string `json:"send_command_topic,omitempty"`
	SetFanSpeedTopic string `json:"set_fan_speed_topic,omitempty"`
	StateTopic string `json:"state_topic,omitempty"`
	SupportedFeatures []string `json:"supported_features,omitempty"`
	UniqueID string `json:"unique_id,omitempty"`
}

