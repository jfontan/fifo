package fifo

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChunk(t *testing.T) {
	q := NewChunk[string]()
	testQueue(t, q)
}

func testQueue(t *testing.T, q Queue[string]) {
	s := NewGeneric[string]()

	var expected []string
	var result []string

	for i := 0; i < 1000; i++ {
		v := strconv.Itoa(i)
		expected = append(expected, v)
		s.Push(v)
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
	}

	for {
		v, ok := s.Pop()
		if !ok {
			break
		}
		result = append(result, v)
	}

	require.Equal(t, expected, result)
	require.Equal(t, 0, cap(s.items))
}
