package filemanager

import (
	"log"
	"os"
)

func WriteJsonToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Failed to create file: %v", err)
		return err
	}
	defer file.Close()

	if _, err = file.Write(data); err != nil {
		log.Printf("Failed to write data to file: %v", err)
		return err
	}

	return nil
}

func ReadFileToBytes(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file %s: %v", filename, err)
		return nil, err
	}
	return data, nil
}
