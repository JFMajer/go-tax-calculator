package prices

import "fmt"

type TaxAndPrices struct {
	TaxRate       float64              `json:"tax_rate"`
	InputPrices   []float64            `json:"no_tax_proces"`
	PricesWithTax map[string][]float64 `json:"prices_including_tax"`
}

func NewTaxAndPrices(taxrate float64, inputprices []float64) *TaxAndPrices {
	return &TaxAndPrices{
		TaxRate:       taxrate,
		InputPrices:   inputprices,
		PricesWithTax: make(map[string][]float64),
	}
}

func (TaP TaxAndPrices) CalculateTaxes() error {
	for _, p := range TaP.InputPrices {
		priceWithTax := p * (1 + TaP.TaxRate)
		TaP.PricesWithTax[fmt.Sprint(TaP.TaxRate)] = append(TaP.PricesWithTax[fmt.Sprint(TaP.TaxRate)], priceWithTax)
	}
	return nil
}
