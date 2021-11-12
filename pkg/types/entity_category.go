package types

type EntityCategory string

const (
	DiagnosticEntity EntityCategory = "diagnostic"
	SystemEntity     EntityCategory = "system"
	ConfigEntity     EntityCategory = "config"
	PrimaryEntity    EntityCategory = ""
)
