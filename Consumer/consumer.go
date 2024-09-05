package Consumer

import "fmt"

func Output(done <-chan bool, data <-chan int) {
	go func() {
		for {
			select {
				case num := <-data:
					fmt.Println("Обработанные данные: ", num)
				case <-done:
					return
			}
		}
	}()
}