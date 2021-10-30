package fifo

const chunkSize = 1024

type chunk[T any] struct {
	values    [chunkSize]T
	pos, size int
	next      *chunk[T]
}

// Chunk is a fifo queue that stores items in queues.
type Chunk[T any] struct {
	head, tail *chunk[T]
}

var _ Queue[string] = new(Chunk[string])

// NewChunk creates a new Chunk fifo queue.
func NewChunk[T any]() *Chunk[T] {
	return new(Chunk[T])
}

// Empty is true if there are no items in the queue.
func (c *Chunk[T]) Empty() bool {
	return c.head == nil || c.head.pos >= c.head.size
}

// Push adds items to the queue.
func (c *Chunk[T]) Push(items ...T) {
	if c.Empty() {
		var chunk chunk[T]
		c.head, c.tail = &chunk, &chunk
	}

	for _, i := range items {
		c.push(i)
	}
}

func (c *Chunk[T]) push(item T) {
	if c.tail.size >= chunkSize {
		var chunk chunk[T]
		c.tail.next = &chunk
		c.tail = &chunk
	}

	cur := c.tail
	cur.values[cur.size] = item
	cur.size++
}

// Pop returns the first item in the queue. Also returns false in case the
// queue is empty.
func (c *Chunk[T]) Pop() (T, bool) {
	if c.Empty() {
		var zero T
		return zero, false
	}

	cur := c.head
	value := cur.values[cur.pos]

	cur.pos++
	if cur.pos >= cur.size {
		c.head = cur.next
	}

	return value, true
}
