package reporter2

import "log"

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
	for _, raw := range raws {
		raw := raw
		if raw.RawType == grandTotal {
			report.Total = makeTotal(&raw)
			continue
		}
		if raw.RawType == categoryTotal {
			if report.Categories == nil {
				report.Categories = []Category{makeCategory(&raw)}
				continue
			}
			report.Categories = append(report.Categories, makeCategory(&raw))
			continue
		}
		if report.Categories[cIndex].Name != raw.Category {
			cIndex++
			report.Categories[cIndex].Products = []Product{makeProduct(&raw)}
		}
		report.Categories[cIndex].Products = append(report.Categories[cIndex].Products, makeProduct(&raw))
	}
	return report, nil
}

func makeProduct(r *Raw) Product {
	return Product{
		Name:  r.Name,
		Total: makeTotal(r),
	}
}

func makeCategory(r *Raw) Category {
	return Category{
		Name:  r.Category,
		Total: makeTotal(r),
	}
}

func makeTotal(r *Raw) Total {
	return Total{
		Count:   r.Count,
		CostSum: r.CostSum,
		SellSum: r.SellSum,
	}
}
