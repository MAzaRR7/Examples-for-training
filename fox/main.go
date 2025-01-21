package main

import (
	"fmt"
	"fox/calculate"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("Старт: %s\n", t.Format(time.RFC3339))

	result1 := make(chan int)
	result2 := make(chan int)

	letters := make(chan map[rune]int)

	resNumbers := make(chan map[int][]int)

	go func() {
		for {
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	go calculate.CalculateSomething(1000, result1)
	go calculate.CalculateSomething(2000, result2)

	go calculate.CalcLetter("Hello world", letters)

	go calculate.CalcNumbers([]int{10, 15, 20, 25, 30, 35, 40}, resNumbers)

	fmt.Println(<-result1)
	fmt.Println(<-result2)

	fmt.Printf("Время работы программы: %s\n", time.Since(t))
}
