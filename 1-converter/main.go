package main

import (
	"fmt"
	"strconv"
	"strings"
)

var arrCurrency = [3]string{"USD", "EUR", "RUB"}

func main() {
	var currencyArray = make(map[string]map[string]float64, 3)
	var sourceCurrency string
	var value float64
	var targetCurrency string
	var message string

	initCurrencyArray(&currencyArray)

	message = "Введите исходную валюту (" + strings.Join(arrCurrency[:], ", ") + "): "
	for {
		sourceCurrency = inputUser(message)

		if isValidCurrency(sourceCurrency) {
			break
		}

		fmt.Println("Валюта указана не верно, повторите ввод.")
	}

	message = "Введите сумму для конвертации: "
	for {
		var ok bool
		result := inputUser(message)

		if value, ok = isValidFloat(result); ok && value > 0 {
			break
		}

		fmt.Println("Введите корректное число больше 0.")
	}

	message = "Введите целевую валюту (" + strings.Join(arrCurrency[:], ", ") + "): "
	for {
		targetCurrency = inputUser(message)

		if isValidCurrency(targetCurrency) && targetCurrency != sourceCurrency {
			break
		}

		fmt.Println("Валюта указана не верно, повторите ввод.")
	}

	result := converterCurrency(value, sourceCurrency, targetCurrency, &currencyArray)
	fmt.Printf("Итого: %.2f%s", result, targetCurrency)

}

func inputUser(message string) string {
	var result string
	fmt.Print(message)
	fmt.Scan(&result)
	return result
}

func isValidCurrency(currency string) bool {
	isCurrent := false

	for _, curr := range arrCurrency {
		if curr == currency {
			isCurrent = true
			break
		}
	}

	return isCurrent
}

func isValidFloat(value string) (float64, bool) {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, false
	}
	return result, true
}

func converterCurrency(value float64, sourceCurrency string, targetCurrency string, currencyArray *map[string]map[string]float64) (result float64) {
	return (*currencyArray)[sourceCurrency][targetCurrency] * value
}

func initCurrencyArray(currencyArray *map[string]map[string]float64) {
	for _, currency := range arrCurrency {
		(*currencyArray)[currency] = make(map[string]float64, 2)
	}

	(*currencyArray)["EUR"]["USD"] = 1.15
	(*currencyArray)["EUR"]["RUB"] = 92.19
	(*currencyArray)["USD"]["EUR"] = 0.87
	(*currencyArray)["USD"]["RUB"] = 79.73
	(*currencyArray)["RUB"]["EUR"] = 0.010847
	(*currencyArray)["RUB"]["USD"] = 0.012542
}
