package httptoolkit

type Filter interface {
	Filter(response *Response) (bool, error)
}

type FilterString struct {
	Keywords []string
}
