package main

import "fmt"

func main() {
	const (
		EurInUsd float64 = 1.17
		RubInUsd float64 = 0.013
		EurInRub float64 = 1.17 / 0.013
	)
}

func getUserInputData() (string, float64, string) {
	var (
		from   string
		amount float64
		to     string
	)
	fmt.Print("Введите валюту, из которой Вы хотите конвертировать: ")
	fmt.Scan(&from)

	fmt.Print("Введите сумму конвертации: ")
	fmt.Scan(&amount)

	fmt.Print("Введите валюту, в которую Вы хотите конвертировать: ")
	fmt.Scan(&to)

	return from, amount, to
}

func calculate(from string, amount float64, to string) float64 {
	return 0.0
}
