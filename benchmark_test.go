package fifo

import "testing"

func BenchmarkGeneric(b *testing.B) {
	q := NewGeneric[string]()
	for n := 0; n < b.N; n++ {
		benchmark(q)
	}
}

func BenchmarkString(b *testing.B) {
	q := NewString()
	for n := 0; n < b.N; n++ {
		benchmark(q)
	}
}

var benchCases = []struct {
	counter int
	add     bool
}{
	{
		counter: 1024 * 1024,
		add:     true,
	},
	{
		counter: 1024 * 1024,
		add:     false,
	},
	{
		counter: 10 * 1024 * 1024,
		add:     true,
	},
	{
		counter: 9 * 1024 * 1024,
		add:     false,
	},
	{
		counter: 10 * 1024 * 1024,
		add:     true,
	},
	{
		counter: 9 * 1024 * 1024,
		add:     false,
	},
}

func benchmark(q Queue[string]) {
	for _, step := range benchCases {
		for c := 0; c < step.counter; c++ {
			if step.add {
				q.Push("string")
			} else {
				_, _ = q.Pop()
			}
		}
	}

}
