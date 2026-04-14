package main

import (
	"demo/jsonLoader/bins"
	"demo/jsonLoader/files"
	"demo/jsonLoader/storage"
	"fmt"
)

func main() {
	fileManager := files.NewFileManager()
	storage := storage.NewStorage(fileManager)
	binList, err := storage.LoadListBin("storage.json")
	if err != nil {
		fmt.Println("Ошибка чтения списка Bin-ов:", err)
	}

	bin := bins.NewBin("1", false, "Test")
	binList = append(binList, *bin)

	err = storage.SaveListBin(&binList, "storage.json")
	if err != nil {
		fmt.Println("Ошибка записи списка Bin-ов:", err)
	}

	fmt.Println("binList:", binList)
	fmt.Println("bin:", bin)

	err = storage.SaveBin(bin, "bin.json")
	if err != nil {
		fmt.Println("Ошибка записи bin-a:", err)
	}

}
