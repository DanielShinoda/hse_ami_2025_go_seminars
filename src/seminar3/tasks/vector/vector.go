package vector

import (
	"errors"
	"fmt"
)

// Option is a functional option type for configuring vector creation
type Option[T any] func(*Vector[T])

// Vector is a generic dynamic array implementation similar to C++ std::vector
type Vector[T any] struct {
	data     []T
	size     int
	capacity int
}

// WithCapacity returns an option to set initial capacity
func WithCapacity[T any](capacity int) Option[T] {
	return func(v *Vector[T]) {
		if capacity > 0 {
			v.data = make([]T, capacity)
			v.capacity = capacity
		}
	}
}

// WithValues returns an option to initialize with values
func WithValues[T any](values ...T) Option[T] {
	return func(v *Vector[T]) {
		v.data = make([]T, len(values))
		copy(v.data, values)
		v.size = len(values)
		v.capacity = len(values)
	}
}

// WithSize returns an option to set initial size with default value
func WithSize[T any](size int, defaultValue T) Option[T] {
	return func(v *Vector[T]) {
		if size <= 0 {
			return
		}
		v.data = make([]T, size)
		for i := range v.data {
			v.data[i] = defaultValue
		}
		v.size = size
		v.capacity = size
	}
}

// WithFill returns an option to fill the vector with n copies of a value
func WithFill[T any](count int, value T) Option[T] {
	return func(v *Vector[T]) {
		if count <= 0 {
			return

		}
		v.data = make([]T, count)
		for i := range v.data {
			v.data[i] = value
		}
		v.size = count
		v.capacity = count
	}
}

// FromSlice returns an option to initialize from an existing slice
func FromSlice[T any](slice []T) Option[T] {
	return func(v *Vector[T]) {
		if len(slice) == 0 {
			return
		}
		v.data = make([]T, len(slice))
		copy(v.data, slice)
		v.size = len(slice)
		v.capacity = len(slice)

	}
}

// New creates a new vector with the given options
func New[T any](options ...Option[T]) *Vector[T] {
	v := &Vector[T]{
		data:     make([]T, 0),
		size:     0,
		capacity: 0,
	}

	// Apply all options
	for _, option := range options {
		option(v)
	}

	return v
}

// NewInt creates a new vector of integers with optional configuration
// This is a convenience function for common types
func NewInt(options ...Option[int]) *Vector[int] {
	return New[int](options...)
}

// NewString creates a new vector of strings with optional configuration
func NewString(options ...Option[string]) *Vector[string] {
	return New[string](options...)
}

// NewFloat64 creates a new vector of float64 with optional configuration
func NewFloat64(options ...Option[float64]) *Vector[float64] {
	return New[float64](options...)
}

// Size returns the number of elements in the vector
func (v *Vector[T]) Size() int {
	return v.size
}

// Capacity returns the capacity of the vector
func (v *Vector[T]) Capacity() int {
	return v.capacity
}

// Empty returns true if the vector is empty
func (v *Vector[T]) Empty() bool {
	return v.size == 0
}

// At returns the element at the specified index with bounds checking
func (v *Vector[T]) At(index int) (T, error) {
	if index < 0 || index >= v.size {
		var t T
		return t, errors.New("ind out of range")
	}
	return v.data[index], nil
}

// Front returns the first element
func (v *Vector[T]) Front() (T, error) {
	if v.size == 0 {
		var t T
		return t, errors.New("empty")
	}
	return v.data[0], nil
}

// Back returns the last element
func (v *Vector[T]) Back() (T, error) {
	if v.size == 0 {
		var t T
		return t, errors.New("empty")
	}
	return v.data[v.size-1], nil
}

// Data returns the underlying slice
func (v *Vector[T]) Data() []T {
	if v.size == 0 {
		return []T{}
	}

	return v.data[:v.size]
}

// PushBack adds an element to the end of the vector
func (v *Vector[T]) PushBack(value T) {
	if v.size >= v.capacity {
		newCap := v.growCapacity()
		v.reserve(newCap)

	}
	v.data[v.size] = value
	v.size++
}

// PopBack removes the last element from the vector
func (v *Vector[T]) PopBack() error {
	if v.size == 0 {
		return errors.New("popback error")
	}
	v.size--
	return nil
}

// Insert inserts an element at the specified position
func (v *Vector[T]) Insert(index int, value T) error {
	if index < 0 || index >= v.size {
		return errors.New("out of range")
	}
	if v.size >= v.capacity {
		v.reserve(v.growCapacity())
	}
	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[index] = value
	v.size++
	return nil
}

// Erase removes the element at the specified position
func (v *Vector[T]) Erase(index int) error {
	if index < 0 || index >= v.size {
		return errors.New("out of range")
	}

	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.size--
	return nil
}

// Clear removes all elements from the vector
func (v *Vector[T]) Clear() {
	for i := 0; i < v.size; i++ {
		var t T
		v.data[i] = t
	}

	v.size = 0
}

// Reserve increases the capacity of the vector
func (v *Vector[T]) Reserve(newCapacity int) {
	if newCapacity > v.capacity {
		v.reserve(newCapacity)
	}
}

// Resize changes the size of the vector
func (v *Vector[T]) Resize(newSize int, value T) {
	if newSize > v.size {
		for v.size < newSize {
			v.PushBack(value)
		}
	} else if newSize < v.size {
		v.size = newSize
	}
}

// Swap exchanges the contents of the vector with another vector
func (v *Vector[T]) Swap(other *Vector[T]) {
	v.data, other.data = other.data, v.data
	v.size, other.size = other.size, v.size
	v.capacity, other.capacity = other.capacity, v.capacity
}

// Assign replaces the contents of the vector with new values
func (v *Vector[T]) Assign(values ...T) {
	newSize := len(values)
	if newSize > v.capacity {
		v.reserve(newSize)
	}
	copy(v.data, values)
	for i := newSize; i < v.size; i++ {
		var t T
		v.data[i] = t
	}

	v.size = newSize
}

// Begin returns the starting index for iteration
func (v *Vector[T]) Begin() int {
	return 0
}

// End returns the ending index for iteration
func (v *Vector[T]) End() int {
	return v.size
}

// String returns a string representation of the vector as Vector[...]
func (v *Vector[T]) String() string {
	return fmt.Sprintf("Vector%v", v.Data())
}

// growCapacity calculates the new capacity when resizing is needed
// returns new capacity
func (v *Vector[T]) growCapacity() int {
	if v.capacity == 0 {
		return 1
	}

	return v.capacity * 2
}

// reserve internal method to handle capacity changes
func (v *Vector[T]) reserve(newCapacity int) {
	if newCapacity <= v.capacity {
		return
	}
	c := make([]T, newCapacity)
	copy(c, v.data[:v.size])
	for i := 0; i < v.size; i++ {
		c[i] = v.data[i]
	}
	v.data = c
	v.capacity = newCapacity

}
