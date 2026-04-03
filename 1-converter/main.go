package main

import (
	"fmt"
	"strconv"
	"strings"
)

const USDEUR float64 = 0.8631
const USDRUB float64 = 80.62

var arrCurrency = [3]string{"USD", "EUR", "RUB"}

func main() {

	var sourceCurrency string
	var value float64
	var targetCurrency string
	var message string

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

	result := converterCurrency(value, sourceCurrency, targetCurrency)
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

func converterCurrency(value float64, sourceCurrency string, targetCurrency string) (result float64) {
	switch {
	case sourceCurrency == "EUR" && targetCurrency == "USD":
		result = value / USDEUR
	case sourceCurrency == "USD" && targetCurrency == "EUR":
		result = value * USDEUR
	case sourceCurrency == "RUB" && targetCurrency == "USD":
		result = value / USDRUB
	case sourceCurrency == "USD" && targetCurrency == "RUB":
		result = value * USDRUB
	case sourceCurrency == "EUR" && targetCurrency == "RUB":
		result = value / USDEUR * USDRUB
	case sourceCurrency == "RUB" && targetCurrency == "EUR":
		result = value / USDRUB * USDEUR
	}
	return
}
