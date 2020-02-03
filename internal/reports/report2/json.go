package report2

import "log"

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

// GetJSON get json
func (s *Service) GetJSON() (Report, error) {
	var report Report

	raws, err := s.getRaws()
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
		} else {
			report.Categories[cIndex].Products = append(report.Categories[cIndex].Products, makeProduct(&raw))
		}
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
