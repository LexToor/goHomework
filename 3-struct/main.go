package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func main() {
	binList := []Bin{}
	bin := NewBin("1", false, "Test")
	fmt.Println(bin)

	binList = append(binList, *bin)
	fmt.Println(binList)
}
