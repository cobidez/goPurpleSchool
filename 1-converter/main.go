package main

import (
	"errors"
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

	result, err := calculate(from, amount, to)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("В результате конверсии Вы получите %.2f %s.", result, to)

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
	result := make(map[string]bool)

	if baseMap == nil || exceptMap == nil {
		for k, v := range *baseMap {
			result[k] = v
		}

		return &result
	}

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
func calculate(from string, amount float64, to string) (float64, error) {
	if from == to {
		return amount, errors.New("валюты конвертации одинаковые")
	}

	var coeff float64
	var is_default bool = false

	switch from {
	case "eur":
		switch to {
		case "rub":
			coeff = EurInRub
		case "usd":
			coeff = EurInUsd
		default:
			coeff = 1
			is_default = true
		}

	case "usd":
		switch to {
		case "rub":
			coeff = 1 / RubInUsd
		case "eur":
			coeff = 1 / EurInUsd
		default:
			coeff = 1
			is_default = true
		}

	case "rub":
		switch to {
		case "usd":
			coeff = RubInUsd
		case "eur":
			coeff = 1 / EurInRub
		default:
			coeff = 1
			is_default = true
		}

	default:
		coeff = 1
		is_default = true
	}

	var err error
	if is_default {
		err = errors.New("не найдена валюта. коэффициент стал равен 1")
	}
	return amount * coeff, err

}
