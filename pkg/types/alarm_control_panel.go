package types

type AlarmControlPanel struct {
	Availability []Availability `json:"availability,omitempty"`
	AvailabilityMode string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	Code string `json:"code,omitempty"`
	CodeArmRequired bool `json:"code_arm_required,omitempty"`
	CodeDisarmRequired bool `json:"code_disarm_required,omitempty"`
	CommandTemplate string `json:"command_template,omitempty"`
	CommandTopic string `json:"command_topic"`
	Device Device `json:"device,omitempty"`
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	EntityCategory EntityCategory `json:"entity_category,omitempty"`
	Icon string `json:"icon,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic string `json:"json_attributes_topic,omitempty"`
	Name string `json:"name,omitempty"`
	PayloadArmAway string `json:"payload_arm_away,omitempty"`
	PayloadArmHome string `json:"payload_arm_home,omitempty"`
	PayloadArmNight string `json:"payload_arm_night,omitempty"`
	PayloadArmVacation string `json:"payload_arm_vacation,omitempty"`
	PayloadArmCustomBypass string `json:"payload_arm_custom_bypass,omitempty"`
	PayloadAvailable string `json:"payload_available,omitempty"`
	PayloadDisarm string `json:"payload_disarm,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	Qos int `json:"qos,omitempty"`
	Retain bool `json:"retain,omitempty"`
	StateTopic string `json:"state_topic"`
	UniqueID string `json:"unique_id,omitempty"`
	ValueTemplate string `json:"value_template,omitempty"`
}

