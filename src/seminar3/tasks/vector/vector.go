package vector

import (
	"errors"
	"fmt"

	"github.com/samber/lo"
)

type Option[T any] func(*Vector[T])

type Vector[T any] struct {
	data     []T
	size     int
	capacity int
}

func WithCapacity[T any](capacity int) Option[T] {
	return func(v *Vector[T]) {
		if capacity < 0 {
			capacity = 0
		}
		v.data = make([]T, 0, capacity)
		v.size = 0
		v.capacity = capacity
	}
}

func WithValues[T any](values ...T) Option[T] {
	return func(v *Vector[T]) {
		copy := append([]T(nil), values...)
		v.data = copy
		v.size = len(copy)
		v.capacity = len(copy)
	}
}

func WithSize[T any](size int, defaultValue T) Option[T] {
	return func(v *Vector[T]) {
		if size < 0 {
			size = 0
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
			count = 0
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
		copy := append([]T(nil), slice...)
		v.data = copy
		v.size = len(copy)
		v.capacity = len(copy)
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
	var z T
	if index < 0 || index >= v.size {
		return z, errors.New("выход за пределы вектора")
	}
	return v.data[index], nil
}

func (v *Vector[T]) Front() (T, error) {
	var z T
	if v.size == 0 {
		return z, errors.New("вектор пуст")
	}
	return v.data[0], nil
}

func (v *Vector[T]) Back() (T, error) {
	var z T
	if v.size == 0 {
		return z, errors.New("вектор пуст")
	}
	return v.data[v.size-1], nil
}

func (v *Vector[T]) Data() []T {
	return v.data[:v.size]
}

func (v *Vector[T]) PushBack(value T) {
	if v.size == v.capacity {
		v.Reserve(v.growCapacity())
	}
	v.data = append(v.data, value)
	v.size++
	v.capacity = cap(v.data)
}

func (v *Vector[T]) PopBack() error {
	if v.size == 0 {
		return errors.New("вектор пуст")
	}
	v.size--
	v.data = v.data[:v.size]
	return nil
}

func (v *Vector[T]) Insert(index int, value T) error {
	if index < 0 || index > v.size {
		return errors.New("выход за пределы вектора")
	}
	if v.size == v.capacity {
		v.Reserve(v.growCapacity())
	}
	v.data = append(v.data, lo.FromPtr(new(T)))
	copy(v.data[index+1:], v.data[index:v.size])
	v.data[index] = value
	v.size++
	return nil
}

func (v *Vector[T]) Erase(index int) error {
	if index < 0 || index >= v.size {
		return errors.New("выход за пределы вектора")
	}
	copy(v.data[index:], v.data[index+1:v.size])
	v.size--
	v.data = v.data[:v.size]
	return nil
}

func (v *Vector[T]) Clear() {
	v.size = 0
	v.data = v.data[:0]
}

func (v *Vector[T]) Reserve(newCapacity int) {
	if newCapacity <= v.capacity {
		return
	}
	v.reserve(newCapacity)
}

func (v *Vector[T]) Resize(newSize int, value T) {
	if newSize < 0 {
		newSize = 0
	}
	if newSize <= v.size {
		v.size = newSize
		v.data = v.data[:newSize]
		return
	}
	if newSize > v.capacity {
		v.Reserve(newSize)
	}
	oldSize := v.size
	v.data = v.data[:newSize]
	for i := oldSize; i < newSize; i++ {
		v.data[i] = value
	}
	v.size = newSize
}

func (v *Vector[T]) Swap(other *Vector[T]) {
	v.data, other.data = other.data, v.data
	v.size, other.size = other.size, v.size
	v.capacity, other.capacity = other.capacity, v.capacity
}

func (v *Vector[T]) Assign(values ...T) {
	copy := append([]T(nil), values...)
	v.data = copy
	v.size = len(copy)
	v.capacity = len(copy)
}

func (v *Vector[T]) Begin() int {
	return 0
}

func (v *Vector[T]) End() int {
	return v.size
}

func (v *Vector[T]) String() string {
	return fmt.Sprintf("Vector%v", v.Data())
}

func (v *Vector[T]) growCapacity() int {
	if v.capacity == 0 {
		return 1
	}
	return v.capacity * 2
}

func (v *Vector[T]) reserve(newCapacity int) {
	if newCapacity <= v.capacity {
		return
	}
	newData := make([]T, v.size, newCapacity)
	copy(newData, v.data[:v.size])
	v.data = newData
	v.capacity = newCapacity
}
