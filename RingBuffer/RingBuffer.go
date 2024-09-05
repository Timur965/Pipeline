package RingBuffer

import "sync"

type RingBuffer struct {
	buff     []int
	capacity int

	start int
	end   int

	m sync.Mutex
}

func NewRingBuffer(cap int) *RingBuffer {
	return &RingBuffer{
		buff:     make([]int, cap),
		capacity: cap,
		start:    0,
		end:      0,
	}
}

func (rb *RingBuffer) Read() int {
	rb.m.Lock()
	defer rb.m.Unlock()
	result := rb.buff[rb.start]

	rb.start = (rb.start + 1) % rb.capacity

	return result
}

func (rb *RingBuffer) Write(value int) {
	rb.m.Lock()
	defer rb.m.Unlock()
	rb.buff[rb.end] = value

	rb.end = (rb.end + 1) % rb.capacity
}

func (rb *RingBuffer) ReadAll() []int {
    rb.m.Lock()
    defer rb.m.Unlock()

    allData := make([]int, 0, rb.capacity)

    for i := rb.start; i != rb.end; i = (i + 1) % rb.capacity {
        allData = append(allData, rb.buff[i])
    }

	rb.start = 0
	rb.end = 0

    return allData
}

func (rb *RingBuffer) IsEmpty() bool {
	return rb.start == 0 && rb.end == 0
}

func (rb *RingBuffer) IsFull() bool {
	return (rb.end+1)%rb.capacity == rb.start
}