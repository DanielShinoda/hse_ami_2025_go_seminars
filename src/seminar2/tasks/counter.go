package tasks

type counter struct {
	value int
}

func newCounter(g int) *counter {
	return &counter{g}
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

func (c *counter) Add(g int) {
	c.value += g
}

func (c *counter) Subtract(g int) {
	c.value -= g
}
