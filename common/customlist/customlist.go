package customlist

type CustomList []string

func (c *CustomList) String() string {
	return "custom global list"
}
