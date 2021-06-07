package customport

import (
	"strconv"
	"strings"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/httpx/common/httpx"
)

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

	for _, potentialPort := range potentialPorts {
		protocol := httpx.HTTPorHTTPS
		potentialPort = strings.TrimSpace(strings.ToLower(potentialPort))
		if strings.HasPrefix(potentialPort, httpx.HTTP+":") {
			potentialPort = strings.TrimPrefix(potentialPort, httpx.HTTP+":")
			protocol = httpx.HTTP
		} else if strings.HasPrefix(potentialPort, httpx.HTTPS+":") {
			potentialPort = strings.TrimPrefix(potentialPort, httpx.HTTPS+":")
			protocol = httpx.HTTPS
		}

		potentialRange := strings.Split(potentialPort, "-")
		if len(potentialRange) < portRangeParts {
			if p, err := strconv.Atoi(potentialPort); err == nil {
				Ports[p] = protocol
			} else {
				gologger.Warning().Msgf("Could not cast port to integer, your value: %s, resulting error %s. Skipping it\n",
					potentialPort, err.Error())
			}
		} else {
			// expand range
			var lowP, highP int
			lowP, err := strconv.Atoi(potentialRange[0])
			if err != nil {
				gologger.Warning().Msgf("Could not cast first port of your port range(%s) to integer, your value: %s, resulting error %s. Skipping it\n",
					potentialPort, potentialRange[0], err.Error())
				continue
			}
			highP, err = strconv.Atoi(potentialRange[1])
			if err != nil {
				gologger.Warning().Msgf("Could not cast last port of your port range(%s) to integer, "+
					"your value: %s, resulting error %s. Skipping it\n",
					potentialPort, potentialRange[1], err.Error())
				continue
			}

			if lowP > highP {
				gologger.Warning().Msgf("first value of port range should be lower than the last part port "+
					"in that range, your range: [%d, %d]. Skipping it\n",
					lowP, highP)
				continue
			}

			for i := lowP; i <= highP; i++ {
				Ports[i] = protocol
			}
		}
	}

	*c = append(*c, value)
	return nil
}
