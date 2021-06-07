package customlist

type CustomList []string

func (c *CustomList) String() string {
	return "custom global list"
}

func (c *CustomList) Set(value string) error {
	*c = append(*c, value)
	return nil
}
