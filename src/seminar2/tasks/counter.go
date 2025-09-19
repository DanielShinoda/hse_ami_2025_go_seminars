package tasks

import "sync"

type counter struct {
	mx     sync.Mutex
	amount int
}

func newCounter(_ int) *counter {
	return &counter{}
}

func (c *counter) Increment() {
	c.mx.Lock()
	c.amount++
	c.mx.Unlock()
}

func (c *counter) Decrement() {
	c.mx.Lock()
	c.amount--
	c.mx.Unlock()
}

func (c *counter) GetValue() int {
	c.mx.Lock()
	res := c.amount
	c.mx.Unlock()
	return res
}

func (c *counter) Reset() {
	c.mx.Lock()
	c.amount = 0
	c.mx.Unlock()
}

func (c *counter) Add(a int) {
	c.mx.Lock()
	c.amount += a
	c.mx.Unlock()
}

func (c *counter) Subtract(a int) {
	c.mx.Lock()
	c.amount += a
	c.mx.Unlock()
}
