package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64
	fmt.Println("__Калькулятор индекса массы тела__")
	fmt.Println()
	fmt.Println("----------------------")
	fmt.Print("Введите ваш рост в сантиметрах:")
	fmt.Scan(&userHeight)
	fmt.Println("----------------------")
	fmt.Print("Введите ваш вес в киллограмах:")
	fmt.Scan(&userKg)
	fmt.Println("----------------------")
	fmt.Println()

	IMT := userKg / math.Pow(userHeight/100, IMTPower)

	fmt.Printf(`----Таблица индексов-----
< 16 Значительный дефицит
< 16 - 18.5 дефицит
< 18.5 - 25 норма
< 25 - 30 лишний вес
< 30 - 35 Ожирение первой степени
< 35 - 40 Ожирение второй степени
< 40 Ожирение третей степени
-------------------------------------
Ваш индекс: %.1f`, IMT)
}
