package types

type Sensor struct {
	Availability []Availability `json:"availability,omitempty"`
	AvailabilityMode string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	Device Device `json:"device,omitempty"`
	DeviceClass DeviceClass `json:"device_class,omitempty"`
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	EntityCategory EntityCategory `json:"entity_category,omitempty"`
	ExpireAfter int `json:"expire_after,omitempty"`
	ForceUpdate bool `json:"force_update,omitempty"`
	Icon string `json:"icon,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic string `json:"json_attributes_topic,omitempty"`
	LastResetValueTemplate string `json:"last_reset_value_template,omitempty"`
	Name string `json:"name,omitempty"`
	PayloadAvailable string `json:"payload_available,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	Qos int `json:"qos,omitempty"`
	StateClass string `json:"state_class,omitempty"`
	StateTopic string `json:"state_topic"`
	UniqueID string `json:"unique_id,omitempty"`
	UnitOfMeasurement string `json:"unit_of_measurement,omitempty"`
	ValueTemplate string `json:"value_template,omitempty"`
}

