package types

type Device struct {
	ConfigurationURL string `json:"configuration_url,omitempty"`
	Connections []string `json:"connections,omitempty"`
	Identifiers []string `json:"identifiers,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Model string `json:"model,omitempty"`
	Name string `json:"name,omitempty"`
	SuggestedArea string `json:"suggested_area,omitempty"`
	SwVersion string `json:"sw_version,omitempty"`
	ViaDevice string `json:"via_device,omitempty"`
}

