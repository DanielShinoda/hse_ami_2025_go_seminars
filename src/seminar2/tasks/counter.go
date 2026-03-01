package tasks

type counter struct {
    value int
}

func newCounter(initial int) *counter {
    return &counter{value: initial}
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

func (c *counter) Add(x int) {
    c.value += x
}

func (c *counter) Subtract(x int) {
    c.value -= x
}