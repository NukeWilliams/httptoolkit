package customheader

type CustomHeaders []string

func (c *CustomHeaders) String() string {
	return "custom global headers"
}

func (c *CustomHeaders) Set(value string) error {
	*c = append(*c, value)
	return nil
}
