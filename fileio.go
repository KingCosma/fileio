package fileio

import (
	"os"
	"strings"
)

var Version string = "1.0"

func WriteString(filename string, data string) error {
	err := os.WriteFile(filename, []byte(data), 0644)
	return err
}

func ReadString(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	return string(data), err
}

func ReadAndSplit(filename string, sep string) ([]string, error) {
	data, err := ReadString(filename)
	if err != nil {
		return nil, err
	}
	arr := strings.Split(data, sep)
	return arr, err
}
