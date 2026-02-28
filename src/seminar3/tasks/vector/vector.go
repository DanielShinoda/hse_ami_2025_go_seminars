package vector

import (
	"errors"
	"fmt"
)

type Option[T any] func(*Vector[T])

type Vector[T any] struct {
	data     []T
	size     int
	capacity int
}

func WithCapacity[T any](capacity int) Option[T] {
	return func(v *Vector[T]) {
		if capacity > 0 {
			v.Reserve(capacity)
		}
	}
}

func WithValues[T any](values ...T) Option[T] {
	return func(v *Vector[T]) {
		v.data = make([]T, len(values))
		copy(v.data, values)
		v.size = len(values)
		v.capacity = len(values)
	}
}

func WithSize[T any](size int, defaultValue T) Option[T] {
	return func(v *Vector[T]) {
		if size < 0 {
			return
		}
		v.data = make([]T, size)
		for i := 0; i < size; i++ {
			v.data[i] = defaultValue
		}
		v.size = size
		v.capacity = size
	}
}

func WithFill[T any](count int, value T) Option[T] {
	return func(v *Vector[T]) {
		if count < 0 {
			return
		}
		v.data = make([]T, count)
		for i := 0; i < count; i++ {
			v.data[i] = value
		}
		v.size = count
		v.capacity = count
	}
}

func FromSlice[T any](slice []T) Option[T] {
	return func(v *Vector[T]) {
		v.data = make([]T, len(slice))
		copy(v.data, slice)
		v.size = len(slice)
		v.capacity = len(slice)
	}
}

func New[T any](options ...Option[T]) *Vector[T] {
	v := &Vector[T]{
		data:     make([]T, 0),
		size:     0,
		capacity: 0,
	}

	for _, option := range options {
		option(v)
	}

	return v
}

func NewInt(options ...Option[int]) *Vector[int] {
	return New[int](options...)
}

func NewString(options ...Option[string]) *Vector[string] {
	return New[string](options...)
}

func NewFloat64(options ...Option[float64]) *Vector[float64] {
	return New[float64](options...)
}

func (v *Vector[T]) Size() int {
	return v.size
}

func (v *Vector[T]) Capacity() int {
	return v.capacity
}

func (v *Vector[T]) Empty() bool {
	return v.size == 0
}

func (v *Vector[T]) At(index int) (T, error) {
	if index < 0 || index >= v.size {
		var zero T
		return zero, errors.New("index out of range")
	}
	return v.data[index], nil
}

func (v *Vector[T]) Front() (T, error) {
	if v.size == 0 {
		var zero T
		return zero, errors.New("vector is empty")
	}
	return v.data[0], nil
}

func (v *Vector[T]) Back() (T, error) {
	if v.size == 0 {
		var zero T
		return zero, errors.New("vector is empty")
	}
	return v.data[v.size-1], nil
}

func (v *Vector[T]) Data() []T {
	return v.data
}

func (v *Vector[T]) PushBack(value T) {
	if v.size == v.capacity {
		newCap := v.growCapacity()
		v.reserve(newCap)
	}
	v.data = append(v.data, value)
	v.size++
	v.capacity = cap(v.data)
}

func (v *Vector[T]) PopBack() error {
	if v.size == 0 {
		return errors.New("vector is empty")
	}

	var zero T
	v.data[v.size-1] = zero
	v.data = v.data[:v.size-1]
	v.size--
	return nil
}

func (v *Vector[T]) Insert(index int, value T) error {
	if index < 0 || index > v.size {
		return errors.New("index out of range")
	}

	if v.size == v.capacity {
		newCap := v.growCapacity()
		v.reserve(newCap)
	}

	v.data = append(v.data, value)
	copy(v.data[index+1:], v.data[index:])
	v.data[index] = value
	v.size++
	v.capacity = cap(v.data)

	return nil
}

func (v *Vector[T]) Erase(index int) error {
	if index < 0 || index >= v.size {
		return errors.New("index out of range")
	}

	var zero T
	copy(v.data[index:], v.data[index+1:])
	v.data[v.size-1] = zero
	v.data = v.data[:v.size-1]
	v.size--

	return nil
}

func (v *Vector[T]) Clear() {
	var zero T
	for i := 0; i < v.size; i++ {
		v.data[i] = zero
	}
	v.data = v.data[:0]
	v.size = 0
}

func (v *Vector[T]) Reserve(newCapacity int) {
	if newCapacity > v.capacity {
		v.reserve(newCapacity)
	}
}

func (v *Vector[T]) Resize(newSize int, value T) {
	if newSize < 0 {
		return
	}

	if newSize < v.size {
		var zero T
		for i := newSize; i < v.size; i++ {
			v.data[i] = zero
		}
		v.data = v.data[:newSize]
		v.size = newSize
	} else if newSize > v.size {
		if newSize > v.capacity {
			v.Reserve(newSize)
		}
		for i := v.size; i < newSize; i++ {
			v.data = append(v.data, value)
		}
		v.size = newSize
		v.capacity = cap(v.data)
	}
}

func (v *Vector[T]) Swap(other *Vector[T]) {
	v.data, other.data = other.data, v.data
	v.size, other.size = other.size, v.size
	v.capacity, other.capacity = other.capacity, v.capacity
}

func (v *Vector[T]) Assign(values ...T) {
	v.Clear()
	for _, val := range values {
		v.PushBack(val)
	}
}

func (v *Vector[T]) Begin() int {
	return 0
}

func (v *Vector[T]) End() int {
	return v.size
}

func (v *Vector[T]) String() string {
	return fmt.Sprintf("Vector%v", v.data)
}

func (v *Vector[T]) growCapacity() int {
	if v.capacity == 0 {
		return 1
	}
	return v.capacity * 2
}

func (v *Vector[T]) reserve(newCapacity int) {
	temp := make([]T, v.size, newCapacity)
	copy(temp, v.data)
	v.data = temp
	v.capacity = newCapacity
}
