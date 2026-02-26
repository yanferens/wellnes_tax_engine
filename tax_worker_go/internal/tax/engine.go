package tax

import (
	"math"
	"tax_worker/internal/models"
)

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func Calculate(subtotal float64, breakdown models.TaxBreakdown) models.ProcessedOrder {

	compositeRate := breakdown.StateRate + breakdown.CountyRate + breakdown.CityRate + breakdown.SpecialRates

	compositeRate = RoundFloat(compositeRate, 4)

	taxAmount := subtotal * compositeRate

	taxAmount = RoundFloat(taxAmount, 2)

	totalAmount := subtotal + taxAmount
	totalAmount = RoundFloat(totalAmount, 2)

	return models.ProcessedOrder{
		CompositeTaxRate: compositeRate,
		TaxAmount:        taxAmount,
		TotalAmount:      totalAmount,
		Breakdown:        breakdown,
	}
}
