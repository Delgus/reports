package report2

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
	raws, err := s.getRaws()
	if err != nil || len(raws) == 0 {
		return Report{}, err
	}

	var cIndex int

	var report Report

	for i := range raws {
		if raws[i].RawType == grandTotal {
			report.Total = makeTotal(&raws[i])
			continue
		}

		if raws[i].RawType == categoryTotal {
			if report.Categories == nil {
				report.Categories = []Category{makeCategory(&raws[i])}
				continue
			}

			report.Categories = append(report.Categories, makeCategory(&raws[i]))

			continue
		}

		if report.Categories[cIndex].Name != raws[i].Category {
			cIndex++

			report.Categories[cIndex].Products = []Product{makeProduct(&raws[i])}
		} else {
			report.Categories[cIndex].Products = append(report.Categories[cIndex].Products, makeProduct(&raws[i]))
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
