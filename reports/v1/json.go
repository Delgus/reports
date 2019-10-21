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

func (r *Reporter) getJson() (Report, error) {
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
	report.Total = newTotal(raws[0])
	report.Categories = []Category{newCategory(raws[0])}
	for _, raw := range raws[1:] {
		report.CostSum = sum(report.CostSum, raw.CostSum)
		report.SellSum = sum(report.SellSum, raw.SellSum)
		report.Count = report.Count + raw.Count
		if report.Categories[cIndex].Name != raw.Category {
			//при добавлении новой категории надо отформатировать данные старой
			report.Categories[cIndex].CostSum = formatCost(report.Categories[cIndex].CostSum)
			report.Categories[cIndex].SellSum = formatCost(report.Categories[cIndex].SellSum)

			report.Categories = append(report.Categories, newCategory(raw))
			cIndex++
			continue
		}
		report.Categories[cIndex].Products = append(report.Categories[cIndex].Products, newProduct(raw))
		report.Categories[cIndex].CostSum = sum(report.Categories[cIndex].CostSum, raw.CostSum)
		report.Categories[cIndex].SellSum = sum(report.Categories[cIndex].SellSum, raw.SellSum)
		report.Categories[cIndex].Count = report.Categories[cIndex].Count + raw.Count
	}
	//отформатировать последнюю категорию
	report.Categories[cIndex].CostSum = formatCost(report.Categories[cIndex].CostSum)
	report.Categories[cIndex].SellSum = formatCost(report.Categories[cIndex].SellSum)
	//отформатировать общие тоталы
	report.CostSum = formatCost(report.CostSum)
	report.SellSum = formatCost(report.SellSum)

	return report, nil
}

//возвращает новую категорию
func newCategory(r Raw) Category {
	return Category{
		Name:     r.Category,
		Total:    newTotal(r),
		Products: []Product{newProduct(r)},
	}
}

//возвращает новый продукт с сразу отформатированными данными
func newProduct(r Raw) Product {
	return Product{
		Name: r.Name,
		Total: Total{
			Count:   r.Count,
			CostSum: formatCost(r.CostSum),
			SellSum: formatCost(r.SellSum),
		},
	}
}

func newTotal(r Raw) Total {
	return Total{
		Count:   r.Count,
		CostSum: r.CostSum,
		SellSum: r.SellSum,
	}
}
