package tasks

type counter struct {
	count int
}

func newCounter(n int) *counter {
	return &counter{n}
}

func (c *counter) Increment() {
	c.count++
}

func (c *counter) Decrement() {
	c.count--
}

func (c *counter) GetValue() int {
	return c.count
}

func (c *counter) Reset() {
	c.count = 0
}

func (c *counter) Add(n int) {
	c.count += n
}

func (c *counter) Subtract(n int) {
	c.count -= n
}
