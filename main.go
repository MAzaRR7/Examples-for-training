package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int
var mutex sync.Mutex

func Deposit(dep int, wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	balance += dep
	fmt.Printf("Поступление: %d, Ваш баланс: %d\n", dep, balance)
	mutex.Unlock()
}

func Withraw(wit int, wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	if balance < wit {
		fmt.Printf("Снятие невозможно! Недостаточно средств! Ваш баланс: %d\n", balance)
		mutex.Unlock()
		return
	}
	balance -= wit
	fmt.Printf("Вы сняли %d, Ваш баланс: %d\n", wit, balance)
	mutex.Unlock()
}

func CheckBalance(wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	fmt.Printf("Ваш баланс: %d\n", balance)
	mutex.Unlock()
}

func Procent(proc int, wg *sync.WaitGroup) {
	defer wg.Done()

	res := balance * proc / 100

	for i := 0; i < 5; i++ {
		mutex.Lock()
		balance += res
		fmt.Printf("Процент %d, Баланс увеличен на %d, Ваш баланс: %d\n", proc, res, balance)
		mutex.Unlock()
	}
}
func main() {
	var wg sync.WaitGroup

	t := time.Now()
	fmt.Printf("Старт: %s\n", t.Format(time.RFC3339))

	wg.Add(1)
	go Deposit(1000, &wg)

	wg.Add(1)
	go Withraw(500, &wg)

	wg.Add(1)
	go Procent(1, &wg)

	wg.Add(1)
	go CheckBalance(&wg)

	wg.Wait()
	fmt.Println("ВСЕ ОПЕРАЦИИ ВЫПОЛНЕНЫ!")
}
