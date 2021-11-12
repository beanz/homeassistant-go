package types

type Lock struct {
	Availability []Availability `json:"availability,omitempty"`
	AvailabilityMode string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	CommandTopic string `json:"command_topic"`
	Device Device `json:"device,omitempty"`
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	EntityCategory EntityCategory `json:"entity_category,omitempty"`
	Icon string `json:"icon,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic string `json:"json_attributes_topic,omitempty"`
	Name string `json:"name,omitempty"`
	Optimistic bool `json:"optimistic,omitempty"`
	PayloadAvailable string `json:"payload_available,omitempty"`
	PayloadLock string `json:"payload_lock,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	PayloadUnlock string `json:"payload_unlock,omitempty"`
	Qos int `json:"qos,omitempty"`
	Retain bool `json:"retain,omitempty"`
	StateLocked string `json:"state_locked,omitempty"`
	StateTopic string `json:"state_topic,omitempty"`
	StateUnlocked string `json:"state_unlocked,omitempty"`
	UniqueID string `json:"unique_id,omitempty"`
	ValueTemplate string `json:"value_template,omitempty"`
}

