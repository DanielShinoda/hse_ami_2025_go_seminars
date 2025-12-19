package tasks

type counter struct {
	val int
}

func newCounter(ourVl int) *counter {
	return &counter{val: ourVl}
}

func (c *counter) Increment() {
	c.val++
}

func (c *counter) Decrement() {
	c.val--
}

func (c *counter) GetValue() int {
	return c.val
}

func (c *counter) Reset() {
	c.val = 0
}

func (c *counter) Add(x int) {
	c.val += x
}

func (c *counter) Subtract(x int) {
	c.val -= x
}
