# fifo

This package implements a fifo queue for strings.

It is implemented on a slice. It recycles the slice in case the whole queue resides in the second half of the allocated slice. It also compacts the underlying slice to a smaller one in case the queue is smaller than 1/4 the capacity of the underlying slice.

## Usage

Get the module:

```
$ go get github.com/jfontan/fifo
```

Example:

```go
import "github.com/jfontan/fifo"

...

    s := NewString()

    s.Push("data")
    v, ok := s.Pop()
    println("Empty?", ok)
    println("Value", v)
```
    