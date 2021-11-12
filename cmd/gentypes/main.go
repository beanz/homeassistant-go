package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
	"os"
)

const (
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdin io.Reader, stdout, stderr io.Writer) error {
	structs := map[string]*Struct{}
	for _, html := range args[1:] {
		f, err := os.Open(html)
		if err != nil {
			return fmt.Errorf("failed to open html file %s: %+v", html, err)
		}
		defer f.Close()
		err = processPage(f, structs)
		if err != nil {
			return fmt.Errorf("failed to parse html file %s: %+v", html, err)
		}
	}
	for _, s := range structs {
		filename := "../../pkg/types/" + s.SnakeName + ".go"
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file %s: %+v", filename, err)
		}
		defer f.Close()
		fmt.Fprintf(f, "package types\n\ntype %s struct {\n", s.Name)
		for _, field := range s.Fields {
			fmt.Fprintf(f, "\t%s %s %s\n", field.Name, field.Type, field.Json)
		}
		fmt.Fprintf(f, "}\n\n")
	}
	err := writeEntityCategory("../../pkg/types/entity_category.go")
	if err != nil {
		return fmt.Errorf("failed to write file entity_category.go: %+v", err)
	}
	return writeDeviceClass("../../pkg/types/device_class.go")
}

func writeEntityCategory(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %+v", filename, err)
	}
	defer f.Close()
	fmt.Fprintf(f, `package types

type EntityCategory string

const (
	DiagnosticEntity EntityCategory = "diagnostic"
	SystemEntity     EntityCategory = "system"
	ConfigEntity     EntityCategory = "config"
	PrimaryEntity    EntityCategory = ""
)
`)
	return nil
}

func writeDeviceClass(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %+v", filename, err)
	}
	defer f.Close()
	fmt.Fprintf(f, `package types

type DeviceClass string

const (
	DeviceClassAqi                      DeviceClass = "aqi"
	DeviceClassBattery                  DeviceClass = "battery"
	DeviceClassCarbonDioxide            DeviceClass = "carbon_dioxide"
	DeviceClassCarbonMonoxide           DeviceClass = "carbon_monoxide"
	DeviceClassCurrent                  DeviceClass = "current"
	DeviceClassDate                     DeviceClass = "date"
	DeviceClassEnergy                   DeviceClass = "energy"
	DeviceClassGas                      DeviceClass = "gas"
	DeviceClassHumidity                 DeviceClass = "humidity"
	DeviceClassIlluminance              DeviceClass = "illuminance"
	DeviceClassMonetary                 DeviceClass = "monetary"
	DeviceClassNitrogenDioxide          DeviceClass = "nitrogen_dioxide"
	DeviceClassNitrogenMonoxide         DeviceClass = "nitrogen_monoxide"
	DeviceClassNitrousOxide             DeviceClass = "nitrous_oxide"
	DeviceClassOzone                    DeviceClass = "ozone"
	DeviceClassPm1                      DeviceClass = "pm1"
	DeviceClassPm10                     DeviceClass = "pm10"
	DeviceClassPm25                     DeviceClass = "pm25"
	DeviceClassPowerFactor              DeviceClass = "power_factor"
	DeviceClassPower                    DeviceClass = "power"
	DeviceClassPressure                 DeviceClass = "pressure"
	DeviceClassSignalStrength           DeviceClass = "signal_strength"
	DeviceClassSulphurDioxide           DeviceClass = "sulphur_dioxide"
	DeviceClassTemperature              DeviceClass = "temperature"
	DeviceClassTimestamp                DeviceClass = "timestamp"
	DeviceClassVolatileOrganicCompounds DeviceClass = "volatile_organic_compounds"
	DeviceClassVoltage                  DeviceClass = "voltage"
)
`)
	return nil
}

func processPage(r io.Reader, structs map[string]*Struct) error {
	doc, err := html.Parse(r)
	if err != nil {
		return fmt.Errorf("HTML Parse error: %+v", err)
	}
	name := "Unknown"
	snakeName := "unknown"
	var cfgNode *html.Node
	var findCfgNode func(*html.Node)
	findCfgNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			s := strings.TrimPrefix(
				strings.TrimSuffix(textContent(n), " - Home Assistant"),
				"MQTT ")
			name = strings.ReplaceAll(s, " ", "")
			snakeName = strings.ToLower(strings.ReplaceAll(s, " ", "_"))
		}
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "config-vars" {
					cfgNode = n
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findCfgNode(c)
		}
	}
	findCfgNode(doc)
	var parseCfg func(string, string, *html.Node)
	parseCfg = func(name string, snakeName string, n *html.Node) {
		class := ""
		if n.Type == html.ElementNode && n.Data == "div" {
			class = findClass(n)
		}
		switch class {
		case "config-vars-label":
			cn := map[string]string{}
			findChildren("span", n, cn)
			omit := ""
			if strings.Contains(cn["config-vars-required"], "Optional") {
				omit = ",omitempty"
			}
			if _, ok := structs[name]; !ok {
				structs[name] = &Struct{
					Name:      name,
					SnakeName: snakeName,
					Fields:    []*Field{},
				}
			}

			if cn["config-vars-label-name"] == "entity_category" {
				cn["config-vars-type"] = "entity_category"
			}
			field := &Field{
				Name:      toCamel(cn["config-vars-label-name"]),
				Type:      toType(cn["config-vars-type"]),
				SnakeName: cn["config-vars-label-name"],
				Json: fmt.Sprintf("`json:\"%s%s\"`",
					cn["config-vars-label-name"],
					omit,
				),
			}
			if field.Name == "Identifiers" {
				field.Type = "[]string"
			}
			if field.Name == "Connections" && field.Type == "list" {
				field.Type = "[][]string"
			}
			structs[name].Fields = append(structs[name].Fields, field)
		case "nested":
			fields := structs[name].Fields
			lastField := fields[len(fields)-1]
			if lastField.Type == "map" {
				lastField.Type = lastField.Name
			} else if lastField.Type == "[]string" {
				lastField.Type = "[]" + lastField.Name
			}
			name = lastField.Name
			snakeName = lastField.SnakeName
			if _, ok := structs[name]; ok {
				return
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseCfg(name, snakeName, c)
		}
	}
	if cfgNode == nil {
		return fmt.Errorf("no config found")
	}
	parseCfg(name, snakeName, cfgNode)
	return nil
}

type Struct struct {
	Name      string
	SnakeName string
	Fields    []*Field
}

type Field struct {
	Name      string
	SnakeName string
	Type      string
	Json      string
}

func findClass(n *html.Node) string {
	for _, a := range n.Attr {
		if a.Key == "class" {
			return a.Val
		}
	}
	return ""
}

func findChildren(kind string, n *html.Node, cn map[string]string) {
	if n.Type == html.ElementNode && n.Data == "span" {
		class := findClass(n)
		cn[class] = textContent(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findChildren("span", c, cn)
	}
}

func textContent(n *html.Node) string {
	s := ""
	if n.Type == html.TextNode {
		s += n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s += textContent(c)
	}
	return strings.Join(strings.Fields(s), " ")
}

func toType(s string) string {
	switch s {
	case "boolean":
		return "bool"
	case "integer":
		return "int"
	case "float":
		return "float32"
	case "template", "icon":
		return "string"
	case "string | list", "list":
		return "[]string"
	case "device_class":
		return "DeviceClass"
	case "entity_category":
		return "EntityCategory"
	}
	return s
}

func toCamel(s string) string {
	res := ""
	upper := true
	for _, v := range s {
		if upper {
			res += strings.ToUpper(string(v))
			upper = false
		} else if v == '_' {
			upper = true
		} else {
			res += string(v)
		}
	}
	if res[len(res)-2:] == "Id" {
		res = res[:len(res)-2] + "ID"
	}
	return strings.ReplaceAll(strings.ReplaceAll(res,
		"Json", "JSON"),
		"Url", "URL")
}
