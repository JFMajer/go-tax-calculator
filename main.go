package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"

	"tax-calculator/filemanager"
	"tax-calculator/prices"
)

func main() {
	ioManager := filemanager.NewFileManager("prices.txt", "tax_and_prices.json")
	calculations := prices.NewMultipleTaxCalculations(ioManager)
	byteData, err := calculations.IOManager.ReadFileToBytes()
	if err != nil {
		log.Fatalf("Failed to read data: %v", err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(byteData))

	var pricesFromFile []float64

	for scanner.Scan() {
		str := scanner.Text()
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting string %q to float64: %v\n", str, err)
			continue
		}
		pricesFromFile = append(pricesFromFile, f)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while reading file: %v\n", err)
	}

	taxRates := []float64{0, 0.15, 0.3, 0.5, 0.75, 1}
	for _, taxrate := range taxRates {
		taxes := prices.NewTaxAndPrices(taxrate, pricesFromFile)
		taxes.CalculateTaxes()
		calculations.AddTaxCalculation(*taxes)
	}

	calculations.WriteToFile()
}
