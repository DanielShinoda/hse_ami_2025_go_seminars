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
		if capacity < 0 {
			capacity = 0
		}
		v.data = make([]T, 0, capacity)
		v.capacity = capacity
	}
}

// WithValues returns an option to initialize with values
func WithValues[T any](values ...T) Option[T] {
	return func(v *Vector[T]) {
		v.data = make([]T, len(values))
		copy(v.data, values)
		v.capacity = len(values)
		v.size = len(values)
	}
}

// WithSize returns an option to set initial size with default value
func WithSize[T any](size int, defaultValue T) Option[T] {
	return func(v *Vector[T]) {
		if size < 0 {
			return
		}
		v.data = make([]T, size)
		v.capacity = size
		v.size = size

		for i := range v.data {
			v.data[i] = defaultValue
		}
	}
}

// WithFill returns an option to fill the vector with n copies of a value
func WithFill[T any](count int, value T) Option[T] {
	return func(v *Vector[T]) {
		v.size = count
		v.capacity = count
		v.data = make([]T, count)
		for i := range v.data {
			v.data[i] = value
		}
	}
}

// FromSlice returns an option to initialize from an existing slice
func FromSlice[T any](slice []T) Option[T] {
	return func(v *Vector[T]) {
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
	return len(v.data)
}

// Capacity returns the capacity of the vector
func (v *Vector[T]) Capacity() int {
	return v.capacity
}

// Empty returns true if the vector is empty
func (v *Vector[T]) Empty() bool {
	return len(v.data) == 0
}

// At returns the element at the specified index with bounds checking
func (v *Vector[T]) At(index int) (T, error) {
	if index < 0 || index > len(v.data) {
		var empty T
		return empty, errors.New("out of bound index")
	}
	return v.data[index], nil
}

// Front returns the first element
func (v *Vector[T]) Front() (T, error) {
	if len(v.data) == 0 {
		var empty T
		return empty, errors.New("empty vector")
	}
	return v.data[0], nil
}

// Back returns the last element
func (v *Vector[T]) Back() (T, error) {
	if len(v.data) == 0 {
		var empty T
		return empty, errors.New("empty vector")
	}
	return v.data[len(v.data)-1], nil
}

// Data returns the underlying slice
func (v *Vector[T]) Data() []T {
	return v.data
}

// PushBack adds an element to the end of the vector
func (v *Vector[T]) PushBack(value T) {
	if v.capacity <= v.size {
		v.capacity = v.growCapacity()
	}
	v.data = append(v.data, value)
	v.size++
}

// PopBack removes the last element from the vector
func (v *Vector[T]) PopBack() error {
	if len(v.data) == 0 {
		return errors.New("empty vector")
	}
	v.data = v.data[:len(v.data)-1]
	v.size--
	return nil
}

// Insert inserts an element at the specified position
func (v *Vector[T]) Insert(index int, value T) error {
	if index < 0 || index > v.size {
		return errors.New("out of bound index")
	}

	if v.size == v.capacity {
		newCapacity := v.growCapacity()
		v.capacity = newCapacity
	}

	v.size++
	v.data = append(v.data, value)
	copy(v.data[index+1:], v.data[index:v.size])
	v.data[index] = value
	return nil
}

// Erase removes the element at the specified position
func (v *Vector[T]) Erase(index int) error {
	if index < 0 || index > v.size {
		return errors.New("out of bound index")
	}

	v.data = append(v.data[:index], v.data[index+1:]...)
	v.size--

	return nil
}

// Clear removes all elements from the vector
func (v *Vector[T]) Clear() {
	v.size = 0
	v.data = make([]T, 0)
}

// Reserve increases the capacity of the vector
func (v *Vector[T]) Reserve(newCapacity int) {
	if v.capacity < newCapacity {
		v.reserve(newCapacity)
	}
}

// Resize changes the size of the vector
func (v *Vector[T]) Resize(newSize int, value T) {
	if newSize < 0 {
		return
	}

	if newSize > v.capacity {
		v.reserve(newSize)
	}

	if newSize > v.size {
		for i := v.size; i < newSize; i++ {
			if i >= len(v.data) {
				v.data = append(v.data, value)
			} else {
				v.data[i] = value
			}
		}
	}
	v.size = newSize
	v.data = v.data[:newSize]
}

// Swap exchanges the contents of the vector with another vector
func (v *Vector[T]) Swap(other *Vector[T]) {
	v.data, other.data = other.data, v.data
	v.size, other.size = other.size, v.size
	v.capacity, other.capacity = other.capacity, v.capacity
}

// Assign replaces the contents of the vector with new values
func (v *Vector[T]) Assign(values ...T) {
	if len(values) >= v.size {
		v.data = values
	} else {
		for i, n := range values {
			v.data[i] = n
		}
	}
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
	if v.size == 0 {
		return "Vector[]"
	}

	base := "Vector["
	for i := 0; i < v.size; i++ {
		if i > 0 {
			base += " "
		}
		base += fmt.Sprintf("%v", v.data[i])
	}
	base += "]"
	return base
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
	newData := make([]T, v.size, newCapacity)
	copy(newData, v.data[:v.size])
	v.data = newData
	v.capacity = newCapacity
}
