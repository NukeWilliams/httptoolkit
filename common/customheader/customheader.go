package customheader

type CustomHeaders []string

func (c *CustomHeaders) String() string {
	return "custom global headers"
}
