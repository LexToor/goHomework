package files

import (
	"os"
	"strings"
)

type FileManagerImpl struct{}

func NewFileManager() *FileManagerImpl {
	return &FileManagerImpl{}
}

type FileManager interface {
	Load(filename string) ([]byte, error)
	Save(data []byte, filename string) error
	CheckExtJson(filename string) bool
}

func (fm *FileManagerImpl) CheckExtJson(filename string) bool {
	idx := strings.LastIndex(filename, ".")
	if idx == -1 {
		return false
	}

	if filename[idx+1:] == "json" {
		return true
	}

	return false
}

func (fm *FileManagerImpl) Load(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (fm *FileManagerImpl) Save(data []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
