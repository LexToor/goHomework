package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var OPERATIONS = [3]string{"AVG", "SUM", "MED"}

func main() {
	var message string
	var operation string
	var numsArr []float64
	functions := map[string]func([]float64){
		"AVG": func(numsArr []float64) { fmt.Println("Среднее:", sum(numsArr)/float64(len(numsArr))) },
		"SUM": func(numsArr []float64) { fmt.Println("Сумма:", sum(numsArr)) },
		"MED": func(numsArr []float64) { fmt.Println(median(numsArr)) },
	}

	message = "Введите операция над числами (" + strings.Join(OPERATIONS[:], ", ") + "): "
	for {
		operation = inputUser(message)

		if chechOper(operation) {
			break
		}

		fmt.Println("Выберите опеацию из предложенных: ", strings.Join(OPERATIONS[:], ", "))
	}

	message = "Введите числа через запятую: "
	for {
		var ok bool
		stringNums := inputUser(message)

		numsArr, ok = stringToFloatSlice(stringNums)
		if ok {
			break
		}
	}

	functions[operation](numsArr)
	// switch operation {
	// case "AVG":
	// 	fmt.Println("Среднее:", sum(numsArr)/float64(len(numsArr)))
	// case "SUM":
	// 	fmt.Println("Сумма:", sum(numsArr))
	// case "MED":
	// 	fmt.Println(median(numsArr))
	// }

}

func inputUser(message string) string {
	var result string
	fmt.Print(message)
	fmt.Scan(&result)
	return result
}

func chechOper(operation string) bool {
	isCorrectOperation := false
	for _, oper := range OPERATIONS {
		if oper == operation {
			isCorrectOperation = true
			break
		}
	}
	return isCorrectOperation
}

func stringToFloatSlice(input string) ([]float64, bool) {
	stringParts := strings.Split(input, ",")

	floatSlice := make([]float64, 0, len(stringParts))

	for _, str := range stringParts {
		str = strings.TrimSpace(str)
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Printf("не удалось преобразовать '%s' в число.\n", str)
			return nil, false
		}

		floatSlice = append(floatSlice, num)
	}
	return floatSlice, true
}

func sum(numsArr []float64) float64 {
	sum := 0.0
	for _, num := range numsArr {
		sum += num
	}
	return sum
}

func median(numsArr []float64) float64 {
	var result float64
	size := len(numsArr)
	sort.Float64s(numsArr)

	if size%2 == 0 {
		result = (numsArr[size/2] + numsArr[size/2-1]) / 2
	} else {
		result = numsArr[(size+1)/2-1]
	}

	return result
}
