package Stage

import (
	"Pipeline/RingBuffer"
	"time"
)

const BufferSize = 10
const BufferTimeout = 10 * time.Second

func FilterNegativeNumber(done <-chan bool, data <-chan int) <-chan int {
	filtred := make(chan int)

	go func() {
		for{
			select{
				case num := <-data:
					if num > 0 {
						select {
							case filtred <- num:
							case <-done:
								return
						}
					}
				case <-done:
					return
			}
		}
	}()

	return filtred
}

func FilterNotMultiples3(done <-chan bool, data <-chan int) <-chan int {
	filtred := make(chan int)

	go func() {
		for{
			select{
				case num := <-data:
					if num % 3 == 0 {
						select {
						case filtred <- num:
						case <-done:
							return
					}
					}
				case <-done:
					return
			}
		}
	}()

	return filtred
}

func Buffering(done <-chan bool, data <-chan int) <-chan int{
	rb := RingBuffer.NewRingBuffer(BufferSize)
	result := make(chan int)

	go func() {
		for{
			select{
				case num := <- data:
					rb.Write(num)
				case <-done:
					return
			}
		}
	}()

	go func(){
		for{
			select{
			case <-time.After(BufferTimeout):
					dataRB := rb.ReadAll()
					for _, num := range dataRB{
						select {
							case result <- num:
							case <-done: 
								return
						}
					}
				case <-done:
					return
			}
		}
	}()

	return result
}