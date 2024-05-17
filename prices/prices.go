package prices

import (
	"encoding/json"
	"log"
	"math"

	"tax-calculator/filemanager"
)

type TaxAndPrices struct {
	TaxRate       float64   `json:"tax_rate"`
	InputPrices   []float64 `json:"no_tax_prices"`
	PricesWithTax []float64 `json:"prices_including_tax"`
}

type MultipleTaxCalculations struct {
	Calculations []TaxAndPrices
	IOManager    filemanager.FileManager
}

func NewTaxAndPrices(taxrate float64, inputprices []float64) *TaxAndPrices {
	return &TaxAndPrices{
		TaxRate:       taxrate,
		InputPrices:   inputprices,
		PricesWithTax: make([]float64, 0),
	}
}

func NewMultipleTaxCalculations(inputfile string, outputfile string) *MultipleTaxCalculations {
	return &MultipleTaxCalculations{
		IOManager: filemanager.NewFileManager(inputfile, outputfile),
	}
}

func (TaP *TaxAndPrices) CalculateTaxes() error {
	for _, p := range TaP.InputPrices {
		priceWithTax := p * (1 + TaP.TaxRate)
		roundedPriceWithTax := math.Round(priceWithTax*100) / 100
		TaP.PricesWithTax = append(TaP.PricesWithTax, roundedPriceWithTax)
	}
	return nil
}

func (MTC *MultipleTaxCalculations) AddTaxCalculation(tap TaxAndPrices) {
	MTC.Calculations = append(MTC.Calculations, tap)
}

func (MTC *MultipleTaxCalculations) WriteToFile() error {
	jsonData, err := json.MarshalIndent(MTC, "", "  ")
	if err != nil {
		log.Printf("Error marshaling data: %v", err)
	}
	err = MTC.IOManager.WriteJsonToFile(jsonData)
	if err != nil {
		log.Printf("Error writing to file: %v", err)
		return err
	}

	return nil

}
