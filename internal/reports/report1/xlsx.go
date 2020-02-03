package report1

import (
	"github.com/tealeg/xlsx"
)

func (s *Service) GetXLSX() (*xlsx.File, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return nil, err
	}

	report, err := s.GetJSON()
	if err != nil {
		return nil, err
	}
	for _, c := range report.Categories {
		for _, p := range c.Products {
			row := sheet.AddRow()
			row.AddCell().SetString(c.Name)
			row.AddCell().SetString(p.Name)
			row.AddCell().SetInt(p.Count)
			row.AddCell().SetString(p.CostSum)
			row.AddCell().SetString(p.SellSum)
		}
		row := sheet.AddRow()
		row.AddCell().SetString(c.Name)
		row.AddCell().SetString("Total:")
		row.AddCell().SetInt(c.Count)
		row.AddCell().SetString(c.CostSum)
		row.AddCell().SetString(c.SellSum)
	}
	row := sheet.AddRow()
	row.AddCell().SetString("")
	row.AddCell().SetString("Grand Total:")
	row.AddCell().SetInt(report.Count)
	row.AddCell().SetString(report.CostSum)
	row.AddCell().SetString(report.SellSum)
	return file, nil
}