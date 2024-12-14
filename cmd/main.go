package main

import (
	"fmt"
	"time"
)

func main() {
	// Выводим несколько строк
	fmt.Println("Строка 1")
	fmt.Println("Строка 2")
	fmt.Println("Строка 3")

	// Ждем 2 секунды
	time.Sleep(2 * time.Second)

	// Удаляем последнюю строку
	fmt.Printf("\033[1A\033[K")
	fmt.Printf("\033[1A\033[K")
	fmt.Printf("\033[1A\033[K")

	// Выводим новую строку вместо удаленной
	fmt.Println("Новая строка")
}
