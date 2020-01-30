package chaptersync

import "sync"

// Counter represents a that can be incremented and read out.
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc increments the counter by 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the counter's current value
func (c *Counter) Value() int {
	return c.value
}
