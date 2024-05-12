package prices

import (
	"encoding/json"
	"log"
	"os"
)

type TaxAndPrices struct {
	TaxRate       float64   `json:"tax_rate"`
	InputPrices   []float64 `json:"no_tax_prices"`
	PricesWithTax []float64 `json:"prices_including_tax"`
}

type MultipleTaxCalculations struct {
	Calculations []TaxAndPrices
}

func NewTaxAndPrices(taxrate float64, inputprices []float64) *TaxAndPrices {
	return &TaxAndPrices{
		TaxRate:       taxrate,
		InputPrices:   inputprices,
		PricesWithTax: make([]float64, 0),
	}
}

func NewMultipleTaxCalculations() *MultipleTaxCalculations {
	return &MultipleTaxCalculations{}
}

func (TaP *TaxAndPrices) CalculateTaxes() error {
	for _, p := range TaP.InputPrices {
		priceWithTax := p * (1 + TaP.TaxRate)
		TaP.PricesWithTax = append(TaP.PricesWithTax, priceWithTax)
	}
	return nil
}

func (MTC *MultipleTaxCalculations) AddTaxCalculation(tap TaxAndPrices) {
	MTC.Calculations = append(MTC.Calculations, tap)
}

func (MTC *MultipleTaxCalculations) WriteToFile() error {
	file, err := os.Create("tax_and_prices.json")
	if err != nil {
		log.Fatalf("Failed to create file")
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(MTC); err != nil {
		log.Fatalf("Error encoding data to JSON: %v", err)
	}

	return nil

}
