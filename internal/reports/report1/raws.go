package report1

import "github.com/shopspring/decimal"

// Raw content data for build report
type Raw struct {
	Category string          `db:"category_name"`
	Name     string          `db:"product_name"`
	Count    int             `db:"count"`
	CostSum  decimal.Decimal `db:"cost_sum"`
	SellSum  decimal.Decimal `db:"sell_sum"`
}

func (s *Service) getRaws() ([]Raw, error) {
	var raws []Raw

	query := `
select 
  product_name,
  category_name,
  sum(count) as count,
  sum(cost_sum) as cost_sum,
  sum(sell_sum) as sell_sum
from cost
join products using (product_id)
join categories using(category_id)
group by
  category_name,
  product_name
order by 
  category_name,
  product_name
`

	err := s.store.Select(&raws, query)

	return raws, err
}
