package fifo

type Queue[T any] interface {
	Pop() (T, bool)
	Push(...T)
	Empty() bool
}
