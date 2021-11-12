package types

type Availability struct {
	PayloadAvailable string `json:"payload_available,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	Topic string `json:"topic"`
}

