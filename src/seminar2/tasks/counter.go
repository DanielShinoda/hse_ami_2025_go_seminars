package tasks

type counter struct {
	cnt int
}

func newCounter(a int) *counter {
	return &counter{a}
}

func (c *counter) Increment() {
	c.cnt++
}

func (c *counter) Decrement() {
	c.cnt--
}

func (c *counter) GetValue() int {
	return c.cnt
}

func (c *counter) Reset() {
	c.cnt = 0
}

func (c *counter) Add(x int) {
	c.cnt += x
}

func (c *counter) Subtract(x int) {
	c.cnt -= x
}
