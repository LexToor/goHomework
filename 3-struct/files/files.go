package files

import (
	"os"
	"strings"
)

func CheckExtJson(filename string) bool {
	idx := strings.LastIndex(filename, ".")
	if idx == -1 {
		return false
	}

	if filename[idx+1:] == "json" {
		return true
	}

	return false
}

func Load(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Save(data []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return nil
	}

	return nil
}
