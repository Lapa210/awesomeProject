package main

import (
	"fmt"
	"math"
)

func main() {
	var i string
	for i != "n" {
		userKg, userHeight := getUserInput()

		IMT := calculateIMT(userKg, userHeight)

		fmt.Println(outputResult(IMT))

		fmt.Print("Для завершения нажмите n Enter, для продолжения нажмите y Enter:")
		fmt.Scan(&i)
		if i == "y" {
			continue
		} else if i == "n" {
			break
		} else {
			fmt.Println("Неверный ввод, завершение программы!")
			break
		}
	}
}

func outputResult(IMT float64) string {
	var res string

	switch {
	case IMT < 16:
		res = fmt.Sprintln("Значительный дефицит!")
	case IMT < 18.5:
		res = fmt.Sprintln("Дефицит")
	case IMT < 25:
		res = fmt.Sprintln("Норма")
	case IMT < 30:
		res = fmt.Sprint("Лишний вес")
	case IMT < 35:
		res = fmt.Sprintln("Ожирение первой степени")
	case IMT < 40:
		res = fmt.Sprintln("Ожирение второй степени")
	case IMT >= 40:
		res = fmt.Sprintln("Ожирение третей степени")
	}
	result := fmt.Sprintf("Ваш индекс: %.1f, %s", IMT, res)
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
	fmt.Println("----------------------------------")
	fmt.Print("Введите ваш рост в сантиметрах:")
	fmt.Scan(&userHeight)
	fmt.Println("----------------------------------")
	fmt.Print("Введите ваш вес в киллограмах:")
	fmt.Scan(&userKg)
	fmt.Println("----------------------------------")
	return userKg, userHeight
}
