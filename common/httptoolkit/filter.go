package httptoolkit

import "strings"

type Filter interface {
	Filter(response *Response) (bool, error)
}

type FilterString struct {
	Keywords []string
}

func (f FilterString) Filter(response *Response) (bool, error) {
	for _, keyword := range f.Keywords {
		if strings.Contains(response.Raw, keyword) {
			return true, nil
		}
	}

	return false, nil
}
