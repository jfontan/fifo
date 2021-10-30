package fifo

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFIFO(t *testing.T) {
	s := New[string]()

	var expected []string
	var result []string

	var max int
	for i := 0; i < 1000; i++ {
		v := strconv.Itoa(i)
		expected = append(expected, v)
		s.Push(v)
		if cap(s.items) > max {
			max = cap(s.items)
		}
	}

	for i := 0; i < 750; i++ {
		v, ok := s.Pop()
		require.True(t, ok)
		result = append(result, v)
	}

	for i := 0; i < 750; i++ {
		v := strconv.Itoa(i)
		expected = append(expected, v)
		s.Push(v)
		if cap(s.items) > max {
			max = cap(s.items)
		}
	}

	for {
		v, ok := s.Pop()
		if !ok {
			break
		}
		result = append(result, v)
	}

	// test adding items after emptying the queue
	for i := 0; i < 256; i++ {
		v := strconv.Itoa(i)
		expected = append(expected, v)
		s.Push(v)
		if cap(s.items) > max {
			max = cap(s.items)
		}
	}

	for {
		v, ok := s.Pop()
		if !ok {
			break
		}
		result = append(result, v)
	}

	require.Equal(t, expected, result)
	require.Less(t, max, 2048)
	require.Equal(t, 0, cap(s.items))
}
