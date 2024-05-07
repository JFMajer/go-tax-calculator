package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var prices []float64

	for scanner.Scan() {
		str := scanner.Text()
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting string %q to float64: %v\n", str, err)
			continue
		}
		prices = append(prices, f)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while reading file: %v\n", err)
	}

	taxRates := []float64{0, 0.07, 0.15, 0.3}
	result := make(map[float64][]float64)

	for _, p := range prices {
		result[p] = []float64{}
		for _, t := range taxRates {
			result[p] = append(result[p], p*t)
		}
	}

	fmt.Println(result)

}