package main

import "fmt"

func main() {
	const usdEur float64 = 0.8631
	const usdRub float64 = 80.62

	fmt.Printf("100 EUR = %.2f RUB", 100/usdEur*usdRub)
}

func inputUser(message string) string {
	var result string
	fmt.Println(message)
	fmt.Scan(&result)
	return result
}

func converter(num float64, sourceCurrency string, targetCurrency string) (result float64) {
	return
}
