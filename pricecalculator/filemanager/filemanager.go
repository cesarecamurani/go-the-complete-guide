package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	FilePath   string
	OutputPath string
}

func New(filePath, outputPath string) FileManager {
	return FileManager{
		FilePath:   filePath,
		OutputPath: outputPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	time.Sleep(3 * time.Second) // Simulate a delay for demonstration purposes

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err = encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to write JSON to file: %v", err)
	}

	return nil
}
