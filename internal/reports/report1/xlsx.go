package report1

import (
	"github.com/shopspring/decimal"
	"github.com/tealeg/xlsx"
)

type xlsxBuilder struct {
	sheet           *xlsx.Sheet
	currentCategory string
	categorySellSum decimal.Decimal
	categoryCostSum decimal.Decimal
	categoryCount   int
	totalSellSum    decimal.Decimal
	totalCostSum    decimal.Decimal
	totalCount      int
}

// GetXLSX get report with xlsx
func (s *Service) GetXLSX() (*xlsx.File, error) {
	file := xlsx.NewFile()

	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return nil, err
	}

	raws, err := s.getRaws()
	if err != nil || len(raws) == 0 {
		return file, err
	}

	xb := xlsxBuilder{sheet: sheet}

	for i := range raws {
		xb.AddCategoryIfNeed(raws[i].Category, i)
		xb.AddProduct(raws[i])
	}

	xb.AddCategory()
	xb.AddTotal()

	return file, nil
}

func (xb *xlsxBuilder) AddCategoryIfNeed(category string, i int) {
	if category != xb.currentCategory && i != 0 {
		xb.AddCategory()
		xb.categorySellSum = decimal.Zero
		xb.categoryCostSum = decimal.Zero
		xb.categoryCount = 0
	}

	xb.currentCategory = category
}

func (xb *xlsxBuilder) AddCategory() {
	row := xb.sheet.AddRow()
	row.AddCell().SetString(xb.currentCategory)
	row.AddCell().SetString("Total:")
	row.AddCell().SetInt(xb.categoryCount)

	costSum, _ := xb.categoryCostSum.Round(2).Float64()
	row.AddCell().SetFloat(costSum)

	sellSum, _ := xb.categorySellSum.Round(2).Float64()
	row.AddCell().SetFloat(sellSum)
}

func (xb *xlsxBuilder) AddProduct(r Raw) {
	row := xb.sheet.AddRow()
	row.AddCell().SetString(r.Category)
	row.AddCell().SetString(r.Name)
	row.AddCell().SetInt(r.Count)

	costSum, _ := r.CostSum.Round(2).Float64()
	row.AddCell().SetFloat(costSum)

	sellSum, _ := r.SellSum.Round(2).Float64()
	row.AddCell().SetFloat(sellSum)

	xb.categorySellSum = xb.categorySellSum.Add(r.SellSum)
	xb.categoryCount += r.Count
	xb.categoryCostSum = xb.categoryCostSum.Add(r.CostSum)

	xb.totalSellSum = xb.totalSellSum.Add(r.SellSum)
	xb.totalCostSum = xb.totalCostSum.Add(r.CostSum)
	xb.totalCount += r.Count
}

func (xb *xlsxBuilder) AddTotal() {
	row := xb.sheet.AddRow()
	row.AddCell().SetString("")
	row.AddCell().SetString("Grand Total:")
	row.AddCell().SetInt(xb.totalCount)

	costSum, _ := xb.totalCostSum.Round(2).Float64()
	row.AddCell().SetFloat(costSum)

	sellSum, _ := xb.totalSellSum.Round(2).Float64()
	row.AddCell().SetFloat(sellSum)
}
