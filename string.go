package fifo

// String implements a fifo queue for strings.
type String struct {
	items []string
	pos   int
}

// NewString creates a new String fifo queue.
func NewString() *String {
	return new(String)
}

// Push appends new strings to the queue.
func (s *String) Push(v ...string) {
	s.items = append(s.items, v...)
}

// Pop return first string in the queue. Also returns false if the queue is
// empty.
func (s *String) Pop() (string, bool) {
	if s.Empty() {
		return "", false
	}

	item := s.items[s.pos]
	s.pos++

	s.compact()
	s.recycle()

	return item, true
}

// Empty returns true if the queue is empty.
func (s *String) Empty() bool {
	return s.pos >= len(s.items)
}

// recycle moves items to the start of the internal slice if they are all in
// the second half of the slice.
func (s *String) recycle() {
	if s.pos < cap(s.items)/2 {
		return
	}

	old := s.items[s.pos:]
	new := s.items[:len(s.items)-s.pos]
	copy(new, old)

	s.items = new
	s.pos = 0
}

// compact creates a new slice for the data if the queue is smaller that 1/4
// the capacity of the internal slice.
func (s *String) compact() {
	if len(s.items)-s.pos > cap(s.items)/4 {
		return
	}

	l := len(s.items) - s.pos
	new := make([]string, l, l*2)
	copy(new, s.items[s.pos:])

	s.items = new
	s.pos = 0
}
