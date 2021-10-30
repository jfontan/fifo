package fifo

type Generic[T any] struct {
	items []T
	pos   int
}

func NewGeneric[T any]() *Generic[T] {
	return new(Generic[T])
}

// Push appends new item to the queue.
func (f *Generic[T]) Push(v ...T) {
	f.items = append(f.items, v...)
}

// Pop return first item in the queue. Also returns false if the queue is
// empty.
func (f *Generic[T]) Pop() (T, bool) {
	if f.Empty() {
		var zero T
		return zero, false
	}

	item := f.items[f.pos]
	f.pos++

	f.compact()
	f.recycle()

	return item, true
}

// Empty returns true if the queue is empty.
func (f *Generic[T]) Empty() bool {
	return f.pos >= len(f.items)
}

// recycle moves items to the start of the internal slice if they are all in
// the second half of the slice.
func (f *Generic[T]) recycle() {
	if f.pos < cap(f.items)/2 {
		return
	}

	old := f.items[f.pos:]
	new := f.items[:len(f.items)-f.pos]
	copy(new, old)

	f.items = new
	f.pos = 0
}

// compact creates a new slice for the data if the queue is smaller that 1/4
// the capacity of the internal slice.
func (f *Generic[T]) compact() {
	if len(f.items)-f.pos > cap(f.items)/4 {
		return
	}

	l := len(f.items) - f.pos
	new := make([]T, l, l*2)
	copy(new, f.items[f.pos:])

	f.items = new
	f.pos = 0
}
