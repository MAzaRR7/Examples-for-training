package calculate

import (
	"fmt"
	"time"
)

func CalculateSomething(n int, res chan int) {
	t := time.Now()
	result := 0

	for i := 0; i <= n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}
	fmt.Printf("Время выполнения программы: %s\n", time.Since(t))

	res <- result
}

func CalcLetter(str string, res chan map[rune]int) {
	m := make(map[rune]int)
	for _, char := range str {
		m[char]++
	}
	for char, count := range m {
		fmt.Printf("Буква `%c` встречается %d раз\n", char, count)
		time.Sleep(time.Millisecond * 5)
	}
	res <- m
}

func CalcNumbers(number []int, res chan map[int][]int) {
	m := make(map[int][]int)

	for _, num := range number {
		r := num % 3
		m[r] = append(m[r], num)
		time.Sleep(time.Millisecond * 10)
	}
	for r, nums := range m {
		fmt.Printf("Остаток: %d; Числa: %v\n", r, nums)
	}

	res <- m
}
