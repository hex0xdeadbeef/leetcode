package dequeue

import (
	"fmt"
	"strings"
)

type Dequeue[T any] struct {
	size int
	body []T

	l, r int
}

func New[T any](size int) *Dequeue[T] {
	body := make([]T, size)

	dequeue := Dequeue[T]{
		size: size,
		body: body,
		l:    0,
		r:    0,
	}

	return &dequeue
}

func (d *Dequeue[T]) Append(elem T) {
	d.body[d.r] = elem
	d.r++
}

func (d *Dequeue[T]) Pop() T {
	elem := d.body[d.l]
	d.l++

	if d.l == d.r {
		d.l, d.r = 0, 0
	}

	return elem
}

func (d *Dequeue[T]) String() string {
	var (
		buf strings.Builder
	)

	for i := d.l; i < d.r; i++ {
		buf.WriteString(fmt.Sprintf(" %v ", d.body[i]))
	}

	return buf.String()
}
