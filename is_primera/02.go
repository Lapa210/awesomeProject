package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	for {
		userKg, userHeight := getUserInput()

		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			panic("Не заданны параметры для расчета")
		}

		fmt.Println(outputResult(IMT))

		isRepeateCalculation := checkRepeatCalculation()

		if !isRepeateCalculation {
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

func calculateIMT(userKg, userHeight float64) (float64, error) {
	const IMTPower = 2
	if userKg <= 0 || userKg <= 0 {
		return 0, errors.New("no params error")
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
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

func checkRepeatCalculation() bool {
	var userString string
	fmt.Print("Для завершения нажмите n, для продолжения нажмите y:")
	fmt.Scan(&userString)

	if userString == "y" || userString == "Y" {
		return true
	}
	return false
}
