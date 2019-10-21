package reporter1

type Raw struct {
	Category string `db:"category_name"`
	Name     string `db:"product_name"`
	Count    int    `db:"count"`
	CostSum  string `db:"cost_sum"`
	SellSum  string `db:"sell_sum"`
}

func (r *Reporter) getRaws() ([]Raw, error) {
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
	err := r.store.Select(&raws, query)
	return raws, err
}
