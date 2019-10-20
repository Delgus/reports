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
  product_name;


select 
  category_name,
  sum(count) as count,
  sum(cost_sum) as cost_sum,
  sum(sell_sum) as sell_sum
from cost
join products using (product_id)
join categories using(category_id)
group by
  category_name;

select 
  sum(count) as count,
  sum(cost_sum) as cost_sum,
  sum(sell_sum) as sell_sum
from cost
join products using (product_id)
join categories using(category_id);