package sync

// Counter represents a that can be incremented and read out.
type Counter struct {
	value int
}

// Inc increments the counter by 1
func (c *Counter) Inc() {
	c.value++
}

// Value returns the counter's current value
func (c *Counter) Value() int {
	return c.value
}
