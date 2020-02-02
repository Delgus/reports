package reporter1

import (
	"log"
)

// Report is main root struct for Report
type Report struct {
	Categories []Category `json:"categories"`
	Total
}

// Category is struct for category of products
type Category struct {
	Name     string    `json:"name"`
	Products []Product `json:"products"`
	Total
}

// Product is struct for product
type Product struct {
	Name string `json:"name"`
	Total
}

// Total is struct for displaing parameters
type Total struct {
	Count   int    `json:"count"`
	CostSum string `json:"cost_sum"`
	SellSum string `json:"sell_sum"`
}

// Error is struct for displaing error
type Error struct {
	Message string `json:"message"`
}

func (r *Reporter) getJSON() (Report, error) {
	var report Report

	raws, err := r.getRaws()
	if err != nil {
		log.Println(err)
		return report, err
	}
	if len(raws) == 0 {
		return report, nil
	}

	var cIndex int
	report.Total = makeTotal(raws[0])
	report.Categories = []Category{makeCategory(raws[0])}
	for _, raw := range raws[1:] {
		report.CostSum = sum(report.CostSum, raw.CostSum)
		report.SellSum = sum(report.SellSum, raw.SellSum)
		report.Count += raw.Count
		if report.Categories[cIndex].Name != raw.Category {
			// before adding new category we need format current category
			report.Categories[cIndex].CostSum = formatCost(report.Categories[cIndex].CostSum)
			report.Categories[cIndex].SellSum = formatCost(report.Categories[cIndex].SellSum)

			report.Categories = append(report.Categories, makeCategory(raw))
			cIndex++
			continue
		}
		report.Categories[cIndex].Products = append(report.Categories[cIndex].Products, makeProduct(raw))
		report.Categories[cIndex].CostSum = sum(report.Categories[cIndex].CostSum, raw.CostSum)
		report.Categories[cIndex].SellSum = sum(report.Categories[cIndex].SellSum, raw.SellSum)
		report.Categories[cIndex].Count += raw.Count
	}
	// format last category
	report.Categories[cIndex].CostSum = formatCost(report.Categories[cIndex].CostSum)
	report.Categories[cIndex].SellSum = formatCost(report.Categories[cIndex].SellSum)
	// format grand total
	report.CostSum = formatCost(report.CostSum)
	report.SellSum = formatCost(report.SellSum)

	return report, nil
}

func makeCategory(r Raw) Category {
	return Category{
		Name:     r.Category,
		Total:    makeTotal(r),
		Products: []Product{makeProduct(r)},
	}
}

func makeProduct(r Raw) Product {
	return Product{
		Name: r.Name,
		Total: Total{
			Count:   r.Count,
			CostSum: formatCost(r.CostSum),
			SellSum: formatCost(r.SellSum),
		},
	}
}

func makeTotal(r Raw) Total {
	return Total{
		Count:   r.Count,
		CostSum: r.CostSum,
		SellSum: r.SellSum,
	}
}
