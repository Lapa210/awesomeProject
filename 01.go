package main

import (
	"fmt"
	"math"
)

func main() {
	userKg, userHeight := getUserInput()

	IMT := calculateIMT(userKg, userHeight)
	fmt.Println(outputResult(IMT))

}

func outputResult(imt float64) string {
	result := fmt.Sprintf(`----Таблица индексов-----
< 16 Значительный дефицит
< 16 - 18.5 дефицит
< 18.5 - 25 норма
< 25 - 30 лишний вес
< 30 - 35 Ожирение первой степени
< 35 - 40 Ожирение второй степени
< 40 Ожирение третей степени
-------------------------------------
Ваш индекс: %.1f`, imt)
	return result
}

func calculateIMT(userKg, userHeight float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
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
	return userKg, userHeight
}
