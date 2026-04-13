package storage

import (
	"demo/jsonLoader/bins"
	"demo/jsonLoader/files"
	"encoding/json"
)

func SaveBin(bin *bins.Bin, filename string) error {
	data, err := json.Marshal(bin)
	if err == nil {
		err = files.Save(data, filename)
	}

	return err
}

func LoadListBin(filename string) ([]bins.Bin, error) {
	var listBin []bins.Bin
	data, err := files.Load(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &listBin)
	if err != nil {
		return nil, err
	}

	return listBin, nil
}

func SaveListBin(listBin *[]bins.Bin, filename string) error {
	data, err := json.Marshal(listBin)
	if err == nil {
		err = files.Save(data, filename)
	}

	return err
}
