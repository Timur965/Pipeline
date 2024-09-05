package Consumer

import (
	"fmt"
	"log"
)

func Output(done <-chan bool, data <-chan int) {
	log.Println("\nВывод обработаных данных")
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