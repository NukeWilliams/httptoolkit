package httptoolkit

import (
	"net/http"
	"strings"

	"github.com/krishpranav/httptoolkit/common/slice"
)

var CSPHeaders = []string{
	"Content-Security-Policy",               // standard
	"Content-Security-Policy-Report-Only",   // standard
	"X-Content-Security-Policy-Report-Only", // non - standard
	"X-Webkit-Csp-Report-Only",              // non - standard
}

type CSPData struct {
	Domains []string `json:"domains,omitempty"`
}

func (h *HTTPTOOLKIT) CSPGrab(r *http.Response) *CSPData {
	domains := make(map[string]struct{})
	for _, cspHeader := range CSPHeaders {
		cspRaw := r.Header.Get(cspHeader)
		if cspRaw != "" {
			rules := strings.Split(cspRaw, ";")
			for _, rule := range rules {
				tokens := strings.Split(rule, " ")
				for _, t := range tokens {
					if isPotentialDomain(t) {
						domains[t] = struct{}{}
					}
				}
			}
		}
	}

	if len(domains) > 0 {
		return &CSPData{Domains: slice.ToSlice(domains)}
	}
	return nil
}

func isPotentialDomain(s string) bool {
	return strings.Contains(s, ".") || strings.HasPrefix(s, "http")
}
