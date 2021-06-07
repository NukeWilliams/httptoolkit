package httptoolkit

import (
	"regexp"
	"strings"
)

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

type FilterRegex struct {
	Regexs []string
}

func (f FilterRegex) Filter(response *Response) (bool, error) {
	for _, regex := range f.Regexs {
		matched, err := regexp.MatchString(regex, response.Raw)
		if err != nil {
			return false, err
		}
		if matched {
			return true, nil
		}
	}

	return false, nil
}

type CustomCallback func(response *Response) (bool, error)
