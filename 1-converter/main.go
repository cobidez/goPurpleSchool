package main

import (
	"fmt"
	"strings"
)

var availableCurrenciesMap *map[string]bool = &map[string]bool{
	"eur": true,
	"usd": true,
	"rub": true,
}

const (
	EurInUsd float64 = 1.17
	RubInUsd float64 = 0.013
	EurInRub float64 = 1.17 / 0.013
)

func main() {

	from, amount, to := getUserInputData()

	fmt.Printf("В результате конверсии Вы получите %.2f %s.", calculate(from, amount, to), to)

}

// qwe
func getUserInputData() (string, float64, string) {
	var (
		from   string
		amount float64
		to     string
	)

	listOfCurrencies := getExceptCurrenciesMap(availableCurrenciesMap, nil)

	from = getUserCurrencyInput("Введите валюту, из которой Вы хотите конвертировать(%s): ", listOfCurrencies)

	amount = getUserAmountInput("Введите сумму конвертации: ")

	listOfCurrencies = getExceptCurrenciesMap(availableCurrenciesMap, &map[string]bool{from: true})

	to = getUserCurrencyInput("Введите валюту, в которую Вы хотите конвертировать(%s): ", listOfCurrencies)

	return from, amount, to
}

// 123
func getExceptCurrenciesMap(baseMap *map[string]bool, exceptMap *map[string]bool) *map[string]bool {
	if baseMap == nil || exceptMap == nil {
		return baseMap
	}

	// Создаем новый слайс и кладем в него значения
	result := make(map[string]bool)

	for k := range *baseMap {
		if _, is_exists := (*exceptMap)[k]; !is_exists {
			result[k] = true
		}
	}

	return &result
}

// 123
func getUserCurrencyInput(message string, validationData *map[string]bool) string {
	const errorText string = "Вы ввели значение \"%s\", возможные значения \"%s\"\n"

	var result string
	availableCurrencyString := getCurrencyString(validationData)

	for {
		// Пишем просьбу человеку
		fmt.Printf(message, availableCurrencyString)

		fmt.Scan(&result)

		result = strings.ToLower(result)

		if _, is_ok := (*validationData)[result]; !is_ok {
			fmt.Printf(errorText, result, availableCurrencyString)
			continue
		}

		return strings.ToLower(result)
	}
}

func getUserAmountInput(message string) float64 {
	const errorText string = "Вы ввели неверное значение, разрешены только положительные числа\n"

	var result float64

	for {
		// Пишем просьбу человеку
		fmt.Print(message)

		_, err := fmt.Scan(&result)

		if err != nil {
			fmt.Print(errorText)
			continue
		} else if result < 0 {
			fmt.Print(errorText)
			continue
		}

		return result
	}
}

// 123
func getCurrencyString(currencyMap *map[string]bool) string {
	sb := strings.Builder{}

	i := 1
	mapLen := len(*currencyMap)
	for k := range *currencyMap {
		sb.WriteString(fmt.Sprint(k))

		if i != mapLen {
			sb.WriteString(", ")
		}
		i++
	}

	return sb.String()
}

// qwe
func calculate(from string, amount float64, to string) float64 {
	var coeff float64

	switch from {
	case "eur":
		switch to {
		case "rub":
			coeff = EurInRub
		case "usd":
			coeff = EurInUsd
		}

	case "usd":
		switch to {
		case "rub":
			coeff = 1 / RubInUsd
		case "eur":
			coeff = 1 / EurInUsd
		}

	case "rub":
		switch to {
		case "usd":
			coeff = RubInUsd
		case "eur":
			coeff = 1 / EurInRub
		}
	}

	return amount * coeff

}
