package files

import (
	"os"

	"github.com/fatih/color"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func WriteFile(name string, data []byte) {
	file, err := os.Create(name)
	if err != nil {
		color.Red(err.Error())
		return
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		color.Red(err.Error())
	}
}
