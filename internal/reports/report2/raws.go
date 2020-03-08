package report2

import "fmt"

const (
	grandTotal    = 1
	categoryTotal = 2
	product       = 3
)

// Raw content data for build report
type Raw struct {
	RawType  int    `db:"raw_type"`
	Category string `db:"category_name"`
	Name     string `db:"product_name"`
	Count    int    `db:"count"`
	CostSum  string `db:"cost_sum"`
	SellSum  string `db:"sell_sum"`
}

func (s *Service) getRaws() ([]Raw, error) {
	var raws []Raw

	query := fmt.Sprintf(`
select 
  %d as raw_type,
  product_name,
  category_name,
  sum(count) as count,
  sum(cost_sum)::numeric(16, 2) as cost_sum,
  sum(sell_sum)::numeric(16, 2) as sell_sum
from cost
join products using (product_id)
join categories using(category_id)
group by
  category_name,
  product_name
union all
select 
  %d as raw_type,
  '' as product_name,
  category_name,
  sum(count) as count,
  sum(cost_sum)::numeric(16, 2) as cost_sum,
  sum(sell_sum)::numeric(16, 2) as sell_sum
from cost
join products using (product_id)
join categories using(category_id)
group by
  category_name
union all 
select 
   %d as raw_type,
  '' as product_name,
  '' as category_name,
  sum(count) as count,
  sum(cost_sum)::numeric(16, 2) as cost_sum,
  sum(sell_sum)::numeric(16, 2) as sell_sum
from cost
join products using (product_id)
join categories using(category_id)
group by 
  raw_type
order by
  raw_type,
  category_name,
  product_name
`, product, categoryTotal, grandTotal)

	err := s.store.Select(&raws, query)

	return raws, err
}
