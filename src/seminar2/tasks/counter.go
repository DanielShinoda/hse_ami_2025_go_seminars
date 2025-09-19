package tasks

type counter struct {
	value int
}

func newCounter(a int) *counter {
	return &counter{value: a}
}

func (c *counter) Increment() {
	c.value++
}

func (c *counter) Decrement() {
	c.value--
}

func (c *counter) GetValue() int {
	return c.value
}

func (c *counter) Reset() {
	c.value = 0
}

func (c *counter) Add(a int) {
	c.value += a
}

func (c *counter) Subtract(a int) {
	c.value -= a
}
