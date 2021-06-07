package customport

func init() {
	Ports = make(map[int]string)
}

const portRangeParts = 2

var Ports map[int]string

type CustomPorts []string

func (c *CustomPorts) String() string {
	return "custom ports"
}
