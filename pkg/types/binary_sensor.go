package types

type BinarySensor struct {
	Availability []Availability `json:"availability,omitempty"`
	AvailabilityMode string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	Device Device `json:"device,omitempty"`
	DeviceClass string `json:"device_class,omitempty"`
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	EntityCategory EntityCategory `json:"entity_category,omitempty"`
	ExpireAfter int `json:"expire_after,omitempty"`
	ForceUpdate bool `json:"force_update,omitempty"`
	Icon string `json:"icon,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic string `json:"json_attributes_topic,omitempty"`
	Name string `json:"name,omitempty"`
	OffDelay int `json:"off_delay,omitempty"`
	PayloadAvailable string `json:"payload_available,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	PayloadOff string `json:"payload_off,omitempty"`
	PayloadOn string `json:"payload_on,omitempty"`
	Qos int `json:"qos,omitempty"`
	StateTopic string `json:"state_topic"`
	UniqueID string `json:"unique_id,omitempty"`
	ValueTemplate string `json:"value_template,omitempty"`
}

