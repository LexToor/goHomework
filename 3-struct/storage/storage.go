package storage

import (
	"demo/jsonLoader/bins"
	"demo/jsonLoader/files"
	"encoding/json"
)

type StorageImpl struct {
	fileManager files.FileManager
}

func NewStorage(fileManager files.FileManager) *StorageImpl {
	return &StorageImpl{
		fileManager: fileManager,
	}
}

type Storage interface {
	SaveBin(bin bins.Bin, filename string) error
	LoadListBin(filename string) ([]bins.Bin, error)
	SaveListBin(listBin []bins.Bin, filename string) error
}

func (s *StorageImpl) SaveBin(bin *bins.Bin, filename string) error {
	data, err := json.Marshal(bin)
	if err == nil {
		err = s.fileManager.Save(data, filename)
	}

	return err
}

func (s *StorageImpl) LoadListBin(filename string) ([]bins.Bin, error) {
	var listBin []bins.Bin
	data, err := s.fileManager.Load(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &listBin)
	if err != nil {
		return nil, err
	}

	return listBin, nil
}

func (s *StorageImpl) SaveListBin(listBin *[]bins.Bin, filename string) error {
	data, err := json.Marshal(listBin)
	if err == nil {
		err = s.fileManager.Save(data, filename)
	}

	return err
}
