package main

import (
	"demo/jsonLoader/bins"
	"fmt"
)

func main() {
	binList := []bins.Bin{}
	bin := bins.NewBin("1", false, "Test")
	fmt.Println(bin)

	binList = append(binList, *bin)
	fmt.Println(binList)
}
