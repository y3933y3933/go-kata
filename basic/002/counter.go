package counter

type Counter struct {
	Count int
}

func (c *Counter) Increment() {
	c.Count++
}
func (c *Counter) Decrement() {
	c.Count--
}
func (c *Counter) Reset() {
	c.Count = 0
}
func (c Counter) Value() int {
	return c.Count
}
