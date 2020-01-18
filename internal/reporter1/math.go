package reporter1

import (
	"fmt"
	"math/big"
	"strconv"
)

const (
	priceFormat       = "%.2f"
	quantityPrecision = 6
)

// суммирует строки
func sum(s1, s2 string) string {
	f1, is := new(big.Float).SetString(s1)
	if !is {
		return "error"
	}
	f2, is := new(big.Float).SetString(s2)
	if !is {
		return "error"
	}
	return new(big.Float).Add(f1, f2).Text('f', quantityPrecision)
}

// форматирует в формат цены
func formatCost(s string) string {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return "error"
	}
	return fmt.Sprintf(priceFormat, f)
}
