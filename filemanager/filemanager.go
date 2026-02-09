package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("Failed to open the file: " + fm.InputFilePath + ", " + err.Error())
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read the file: " + fm.InputFilePath + ", " + err.Error())
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteJson(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Failed to write to file: " + fm.OutputFilePath + ", " + err.Error())
	}

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to JSON, " + err.Error())
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
