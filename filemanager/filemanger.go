package filemanager

import (
	"log"
	"os"
)

type FileManager struct {
	InputFile  string
	OutputFile string
}

func (fm FileManager) WriteJsonToFile(data []byte) error {
	file, err := os.Create(fm.OutputFile)
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

func (fm FileManager) ReadFileToBytes() ([]byte, error) {
	data, err := os.ReadFile(fm.InputFile)
	if err != nil {
		log.Printf("Error reading file %s: %v", fm.InputFile, err)
		return nil, err
	}
	return data, nil
}

func NewFileManager(inputfile string, outputfile string) FileManager {
	return FileManager{
		InputFile:  inputfile,
		OutputFile: outputfile,
	}
}
