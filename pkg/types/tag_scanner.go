package types

type TagScanner struct {
	Topic string `json:"topic"`
	ValueTemplate string `json:"value_template,omitempty"`
	Device Device `json:"device"`
}

