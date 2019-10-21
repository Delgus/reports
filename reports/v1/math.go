package v1

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

const (
	delimiter         = "."
	priceFormat       = "%.2f"
	quantityPrecision = 6
)

//суммирует строки
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

//убирает незначащие нули
func truncateZeros(s string) string {
	if !strings.Contains(s, delimiter) {
		return s
	}
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, delimiter)
	return s
}

//форматирует в формат цены
func formatCost(s string) string {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return "error"
	}
	return fmt.Sprintf(priceFormat, f)
}
