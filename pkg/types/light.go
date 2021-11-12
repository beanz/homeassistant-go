package types

type Light struct {
	Availability []Availability `json:"availability,omitempty"`
	AvailabilityMode string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	BlueTemplate string `json:"blue_template,omitempty"`
	BrightnessTemplate string `json:"brightness_template,omitempty"`
	ColorTempTemplate string `json:"color_temp_template,omitempty"`
	CommandOffTemplate string `json:"command_off_template"`
	CommandOnTemplate string `json:"command_on_template"`
	CommandTopic string `json:"command_topic"`
	Device Device `json:"device,omitempty"`
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	EntityCategory EntityCategory `json:"entity_category,omitempty"`
	EffectList []string `json:"effect_list,omitempty"`
	EffectTemplate string `json:"effect_template,omitempty"`
	GreenTemplate string `json:"green_template,omitempty"`
	Icon string `json:"icon,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic string `json:"json_attributes_topic,omitempty"`
	MaxMireds int `json:"max_mireds,omitempty"`
	MinMireds int `json:"min_mireds,omitempty"`
	Name string `json:"name,omitempty"`
	Optimistic bool `json:"optimistic,omitempty"`
	PayloadAvailable string `json:"payload_available,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	Qos int `json:"qos,omitempty"`
	RedTemplate string `json:"red_template,omitempty"`
	Schema string `json:"schema,omitempty"`
	StateTemplate string `json:"state_template,omitempty"`
	StateTopic string `json:"state_topic,omitempty"`
	UniqueID string `json:"unique_id,omitempty"`
}

