package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type modeMapObject struct {
	mode       byte
	calcObject mathCalcInterfacre
}

const (
	avgMode byte = iota
	sumMode
	medMode
)

var (
	inputModeMap    *map[string]modeMapObject = getInputModeMap()
	inputModeString string                    = getInputModeString()
)

type mathCalcInterfacre interface {
	calc(*[]int) int
}

type sumCalc struct{}
type avgCalc struct{}
type medCalc struct{}

func main() {
	mode := getUserMode()
	numberSlice := getUserNumberData()

	calcObject := (*inputModeMap)[mode].calcObject

	result := calcObject.calc(numberSlice)

	fmt.Printf("Результат операции %s равен %d", mode, result)
}

func getUserMode() string {
	var input string
	for {
		fmt.Printf("Введите желаемую операцию (возможные значения: %s): ", inputModeString)

		fmt.Scanln(&input)

		if !isModeInputOk(input) {
			fmt.Printf("Вы ввели значение [%s], возможные значения: %s!\n", input, inputModeString)
			continue
		}

		return strings.ToUpper(input)
	}
}

func isModeInputOk(input string) bool {
	for k := range *inputModeMap {
		if strings.ToUpper(input) == k {
			return true
		}
	}

	return false
}

func getUserNumberData() *[]int {
	var numbers string

	for {
		fmt.Print("Введите числа через запятую: ")
		fmt.Scan(&numbers)

		numberSlice, err := transformNumberStringToNumberSlice(numbers)

		if err != nil {
			fmt.Printf("Введены неверные символы: %s.", err)
			continue
		}

		return numberSlice
	}
}

func transformNumberStringToNumberSlice(numberString string) (*[]int, error) {
	const errorPattern string = "Найдено значение [%s]"

	var result []int
	parts := strings.Split(numberString, ",")

	for _, part := range parts {
		trimmedPart := strings.TrimSpace(part)

		if trimmedPart == "" {
			continue
		}

		num, err := strconv.Atoi(trimmedPart)

		if err != nil {
			return nil, fmt.Errorf(errorPattern, part)
		}

		result = append(result, num)
	}

	return &result, nil
}

func getInputModeMap() *map[string]modeMapObject {
	result := map[string]modeMapObject{
		"AVG": {avgMode, avgCalc{}},
		"SUM": {sumMode, sumCalc{}},
		"MED": {medMode, medCalc{}},
	}

	return &result
}

func getInputModeString() string {
	const separator string = ", "

	var sb strings.Builder

	i := 1
	for k := range *inputModeMap {
		sb.WriteString(k)

		if i < len(*inputModeMap) {
			sb.WriteString(separator)
		}

		i++
	}

	return sb.String()
}

func (sc sumCalc) calc(slice *[]int) int {
	var result int = 0

	for _, val := range *slice {
		result += val
	}

	return result
}

func (ac avgCalc) calc(slice *[]int) int {
	if len(*slice) == 0 {
		return 0
	}

	sum := sumCalc{}.calc(slice)

	return sum / len(*slice)
}

func (mc medCalc) calc(slice *[]int) int {
	var result int = 0

	if len(*slice) == 0 {
		return 0
	}

	sliceCopy := append([]int(nil), *slice...)
	sort.Ints(sliceCopy)

	if len(sliceCopy)%2 == 0 {
		idx := len(sliceCopy) / 2
		result = avgCalc{}.calc(&[]int{
			sliceCopy[idx-1],
			sliceCopy[idx]})
	} else {
		idx := len(sliceCopy) / 2
		result = sliceCopy[idx]
	}

	return result
}
