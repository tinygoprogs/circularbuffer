package circularbuffer

/*
A circular buffer.

Example usage:
```go
c := NewCircularBuffer(5)
c.Insert([]int{1, 2, 3})
fmt.Println(c)
c.Insert([]int{4})
fmt.Println(c)
c.Insert([]int{5, 6})
fmt.Println(c)
fmt.Println(c.Get())
```
*/
type CircularBuffer struct {
	content    []int
	begin, end int
	empty      bool
}

func NewCircularBuffer(maxsize int) *CircularBuffer {
	return &CircularBuffer{
		content: make([]int, maxsize),
		begin:   0,
		end:     0,
		empty:   true,
	}
}

func (c *CircularBuffer) Insert(values []int) {
	var (
		i, v, newEnd int
	)
	c.empty = false
	for i = c.end; i < len(c.content) && v < len(values); i++ {
		c.content[i] = values[v]
		v++
		newEnd = i
	}
	for i = 0; i < len(c.content) && v < len(values); i++ {
		c.content[i] = values[v]
		v++
		newEnd = i
	}
	newEnd += 1 // exclusive end of interval
	if c.end != c.begin && i > c.begin {
		c.begin = newEnd
	}
	c.end = newEnd
}

func (c *CircularBuffer) Get() []int {
	if c.empty {
		return []int{}
	}
	if c.end <= c.begin {
		first := c.content[c.begin:]
		second := c.content[:c.end]
		return append(first, second...)
	} else {
		return c.content[c.begin:c.end]
	}
}
