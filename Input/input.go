package Input

import (
	"fmt"
	"strconv"
)

func InputData() (<-chan int, <-chan bool) {
	data := make(chan int)
	done := make(chan bool)

	go func() {
		defer close(done)
		defer close(data)

		text := ""
		for {
			fmt.Print("Введите число: ")
			fmt.Scan(&text)

			if text == "exit" {
				return
			}

			num, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println("Вводить можно только числа")
				continue
			}

			data <- num
		}
	}()

	return data, done
}
