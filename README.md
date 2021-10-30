# fifo

*Note:* it uses generics so it needs go 1.18 (or current go tip).

This package implements fifo queues:

* `Generic`: big slice generic implementation
* `String`: non generic version of big slice implementation
* `Chunk`: chunk storage implementation

`Generic` and `String` are implemented on a slice. It recycles the slice in case the whole queue resides in the second half of the allocated slice. It also compacts the underlying slice to a smaller one in case the queue is smaller than 1/4 the capacity of the underlying slice.

`Chunk` implementation is based on https://github.com/foize/go.fifo. Values are stored in chunks of size 1024. Old chunks are freed when empty after retrieving values.

`Chunk` implementation is faster and consumes less memory. Here are the results from the provided benchmark:

```
BenchmarkGeneric
BenchmarkGeneric-16    	       1	1356704075 ns/op	2474543744 B/op	22020262 allocs/op
BenchmarkString
BenchmarkString-16     	       1	1392281938 ns/op	2474523512 B/op	22020217 allocs/op
BenchmarkChunk
BenchmarkChunk-16      	       2	 778262590 ns/op	748683516 B/op	22041603 allocs/op
```

## Usage

Get the module:

```
$ go get github.com/jfontan/fifo
```

Example for `Chunk` implementation:

```go
import "github.com/jfontan/fifo"

...

    s := NewChunk[string]()

    s.Push("data")
    v, ok := s.Pop()
    println("Empty?", ok)
    println("Value", v)
```