package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(filename string) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("Failed to open the file: " + filename + ", " + err.Error())
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read the file: " + filename + ", " + err.Error())
	}

	file.Close()
	return lines, nil
}
