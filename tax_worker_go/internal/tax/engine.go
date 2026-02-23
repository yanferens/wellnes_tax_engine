package tax

import "tax_worker/internal/models"

func Calculate(subtotal float64, breakdown models.TaxBreakdown) models.ProcessedOrder {
	compositeRate := breakdown.StateRate + breakdown.CountyRate + breakdown.CityRate + breakdown.SpecialRates
	taxAmount := subtotal * compositeRate
	totalAmount := subtotal + taxAmount

	return models.ProcessedOrder{
		CompositeTaxRate: compositeRate,
		TaxAmount:        taxAmount,
		TotalAmount:      totalAmount,
		Breakdown:        breakdown,
	}
}
