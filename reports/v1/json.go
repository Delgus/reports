package v1

import (
	"log"
)

type Report struct {
	Categories []Category `json:"categories"`
	Total
}

type Category struct {
	Name     string    `json:"name"`
	Products []Product `json:"products"`
	Total
}

type Product struct {
	Name string `json:"name"`
	Total
}

type Total struct {
	Count   int    `json:"count"`
	CostSum string `json:"cost_sum"`
	SellSum string `json:"sell_sum"`
}

type Error struct {
	Message string `json:"message"`
}

func (r *Reporter) GetJson() (Report, error) {
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
	report.Categories = []Category{newCategory(raws[0])}
	for _, raw := range raws[1:] {
		if report.Categories == nil || report.Categories[cIndex].Name != raw.Category {
			report.Categories = append(report.Categories, newCategory(raw))
			cIndex++
			continue
		}
		report.Categories[cIndex].Products = append(report.Categories[cIndex].Products, newProduct(raw))
	}
	return report, nil
}

func newCategory(r Raw) Category {
	return Category{
		Name:     r.Category,
		Total:    Total{},
		Products: []Product{newProduct(r)},
	}
}

func newProduct(r Raw) Product {
	return Product{
		Name: r.Name,
		Total: Total{
			Count:   r.Count,
			CostSum: r.CostSum,
			SellSum: r.SellSum,
		},
	}
}
