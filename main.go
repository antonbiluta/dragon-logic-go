package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Загадываем случайное число
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(100)
	var attempts int

	fmt.Println("Выберите режим игры:")
	fmt.Println("1 - Один игрок")
	fmt.Println("2 - Два игрока")

	var mode int
	fmt.Scan(&mode)

	fmt.Println("Выберите уровень сложности:")
	fmt.Println("1 - Легкий (20 попыток)")
	fmt.Println("2 - Средний (10 попыток)")
	fmt.Println("3 - Тяжелый (5 попыток)")

	var level int
	fmt.Scan(&level)

	switch level {
	case 1:
		attempts = 20
	case 2:
		attempts = 10
	case 3:
		attempts = 5
	default:
		fmt.Println("Неверный выбор уровня. Игра окончена.")
		return
	}

	switch mode {
	case 1:
		fmt.Println("Компьютер загадает число.")
	case 2:
		fmt.Println("Первый игрок загадывает число от 1 до 100.")
		fmt.Println("Второй игрок угадывает число.")
		fmt.Printf("\nУ второго игрока есть %d попыток.\n", attempts)
		fmt.Println("Первый игрок, введите загаданное число:")
		fmt.Scan(&number)
	default:
		fmt.Println("Неверный выбор режима. Игра окончена.")
		return
	}

	fmt.Printf("Угадайте число от 1 до 100. У вас есть %d попыток.\n", attempts)

	for attempts > 0 {
		var guess int
		fmt.Scan(&guess)

		if guess < number {
			fmt.Println("Загаданное число больше")
		} else if guess > number {
			fmt.Println("Загаданное число меньше")
		} else {
			fmt.Println("Вы угадали число!")
			return
		}

		attempts--
		fmt.Printf("У вас осталось %d попыток.\n", attempts)
	}

	fmt.Println("Вы исчерпали все попытки. Игра окончена.")
	fmt.Printf("Загаданное число было %d.\n", number)
}
