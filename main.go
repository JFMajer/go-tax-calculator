package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"tax-calculator/prices"
)

func main() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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

	calculations := prices.NewMultipleTaxCalculations()
	taxRates := []float64{0, 0.15, 0.3, 0.5, 0.75, 1}
	for _, taxrate := range taxRates {
		taxes := prices.NewTaxAndPrices(taxrate, pricesFromFile)
		taxes.CalculateTaxes()
		calculations.AddTaxCalculation(*taxes)
	}

	calculations.WriteToFile()
}
