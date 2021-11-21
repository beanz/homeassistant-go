package mqtt

import (
	"fmt"
	"strings"
)

func ConfigTopic(prefix, component, name, variable string) string {
	return fmt.Sprintf("%s/%s/%s_%s/config",
		prefix, component, name, variable)
}

func StateTopic(prefix, name string) string {
	return fmt.Sprintf("%s/%s/state", prefix, name)
}

func AvailabilityTopic(prefix, name string) string {
	return fmt.Sprintf("%s/%s/availability", prefix, name)
}

func TopicSafe(s string) string {
	r := strings.ReplaceAll(s, "/", "_slash_")
	r = strings.ReplaceAll(r, "#", "_hash_")
	r = strings.ReplaceAll(r, "+", "_plus_")
	r = strings.ReplaceAll(r, "-", "_")
	r = strings.ReplaceAll(r, ":", "_")
	r = strings.TrimLeft(r, "_")
	r = strings.TrimRight(r, "_")
	return strings.ToLower(r)
}
