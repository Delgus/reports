package reporter2

import (
	"sort"

	"github.com/tealeg/xlsx"
)

func (r *Reporter) getXLSX() (*xlsx.File, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return nil, err
	}

	raws, err := r.getRaws()
	if err != nil {
		return nil, err
	}

	sort.Slice(raws, func(i, j int) bool {
		// общий итог опускаем вниз
		if raws[i].RawType == grandTotal {
			return false
		}
		if raws[j].RawType == grandTotal {
			return true
		}

		if raws[i].Category == raws[j].Category {
			// внутри категории тоталы опускаем вниз
			if raws[i].RawType == categoryTotal {
				return false
			}
			if raws[j].RawType == categoryTotal {
				return true
			}

			// сортировка по продуктам
			return raws[i].Name < raws[j].Name
		}
		// сортировка по категориям
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
		return "Общий итог:"
	case categoryTotal:
		return "Итого:"
	default:
		return r.Name
	}
}
