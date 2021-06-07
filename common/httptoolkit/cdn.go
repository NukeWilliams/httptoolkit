package httptoolkit

import (
	"fmt"
	"net"
)

func (h *HTTPTOOLKIT) CdnCheck(ip string) (bool, error) {
	if h.cdn == nil {
		return false, fmt.Errorf("cdn cleint not configured")
	}

	return h.cdn.Check(net.ParseIP((ip)))
}
