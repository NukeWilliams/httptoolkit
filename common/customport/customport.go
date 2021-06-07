package customport

import "strings"

func init() {
	Ports = make(map[int]string)
}

const portRangeParts = 2

var Ports map[int]string

type CustomPorts []string

func (c *CustomPorts) String() string {
	return "custom ports"
}

func (c *CustomPorts) Set(value string) error {

	potentialPorts := strings.Split(value, ",")

}
