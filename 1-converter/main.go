package main

import "fmt"

func main() {
	const usdEur float64 = 0.8631
	const usdRub float64 = 80.62

	fmt.Printf("100 EUR = %.2f RUB", 100/usdEur*usdRub)
}
