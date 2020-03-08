package report2

import (
	"sort"

	"github.com/tealeg/xlsx"
)

// GetXLSX get xlsx
func (s *Service) GetXLSX() (*xlsx.File, error) {
	file := xlsx.NewFile()

	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return nil, err
	}

	raws, err := s.getRaws()
	if err != nil {
		return nil, err
	}

	sort.Slice(raws, func(i, j int) bool {
		// down grand total
		if raws[i].RawType == grandTotal {
			return false
		}
		if raws[j].RawType == grandTotal {
			return true
		}

		if raws[i].Category == raws[j].Category {
			// down category total
			if raws[i].RawType == categoryTotal {
				return false
			}
			if raws[j].RawType == categoryTotal {
				return true
			}

			// sort by product name
			return raws[i].Name < raws[j].Name
		}
		// sort by category
		return raws[i].Category < raws[j].Category
	})

	for _, raw := range raws {
		raw := raw
		row := sheet.AddRow()
		row.AddCell().SetString(raw.Category)
		row.AddCell().SetString(productNameOrTotal(&raw))
		row.AddCell().SetInt(raw.Count)
		row.AddCell().SetString(raw.CostSum)
		row.AddCell().SetString(raw.SellSum)
	}

	return file, nil
}

func productNameOrTotal(r *Raw) string {
	switch r.RawType {
	case grandTotal:
		return "Grand total:"
	case categoryTotal:
		return "Total:"
	default:
		return r.Name
	}
}
