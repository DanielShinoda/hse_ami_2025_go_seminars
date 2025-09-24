package tasks

type counter struct {
	n int
}

func newCounter(init_n int) *counter {
	return &counter{n: init_n}
}

func (c *counter) Increment() {
	c.n++
}

func (c *counter) Decrement() {
	c.n--
}

func (c *counter) GetValue() int {
	return c.n
}

func (c *counter) Reset() {
	c.n = 0
}

func (c *counter) Add(x int) {
	c.n += x
}

func (c *counter) Subtract(x int) {
	c.n -= x
}
