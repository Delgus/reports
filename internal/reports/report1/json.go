package report1

import "github.com/shopspring/decimal"

// Report is main root struct for Report
type Report struct {
	Categories []*Category `json:"categories"`
	Total
}

// Category is struct for category of products
type Category struct {
	Name     string     `json:"name"`
	Products []*Product `json:"products"`
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

type reportBuilder struct {
	Report
	currentCategory *Category
	categorySellSum decimal.Decimal
	totalSellSum    decimal.Decimal
	categoryCostSum decimal.Decimal
	totalCostSum    decimal.Decimal
}

// GetJSON get json
func (s *Service) GetJSON() (Report, error) {
	raws, err := s.getRaws()
	if err != nil || len(raws) == 0 {
		return Report{}, err
	}

	var rb reportBuilder

	for i := range raws {
		if rb.isNewCategory(raws[i].Category) {
			rb.AddCategory(raws[i])
		}

		rb.AddProduct(raws[i])
	}

	rb.AddTotal()

	return rb.Report, nil
}

func (rb *reportBuilder) isNewCategory(category string) bool {
	return len(rb.Categories) == 0 || category != rb.currentCategory.Name
}

func (rb *reportBuilder) AddCategory(r Raw) {
	rb.currentCategory = &Category{
		Name:     r.Category,
		Products: make([]*Product, 0),
	}

	rb.categoryCostSum = decimal.Zero
	rb.categorySellSum = decimal.Zero
	rb.Categories = append(rb.Categories, rb.currentCategory)
}

func (rb *reportBuilder) AddProduct(r Raw) {
	rb.currentCategory.Products = append(rb.currentCategory.Products, &Product{
		Name: r.Name,
		Total: Total{
			Count:   r.Count,
			CostSum: r.CostSum.StringFixed(2),
			SellSum: r.SellSum.StringFixed(2),
		},
	})

	rb.categorySellSum = rb.categorySellSum.Add(r.SellSum)
	rb.currentCategory.SellSum = rb.categorySellSum.StringFixed(2)

	rb.currentCategory.Count += r.Count

	rb.categoryCostSum = rb.categoryCostSum.Add(r.CostSum)
	rb.currentCategory.CostSum = rb.categoryCostSum.StringFixed(2)

	rb.totalSellSum = rb.totalSellSum.Add(r.SellSum)
	rb.totalCostSum = rb.totalCostSum.Add(r.CostSum)
	rb.Count += r.Count
}

func (rb *reportBuilder) AddTotal() {
	rb.CostSum = rb.totalCostSum.StringFixed(2)
	rb.SellSum = rb.totalSellSum.StringFixed(2)
}
