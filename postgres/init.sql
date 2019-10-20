CREATE TABLE products (
	product_id serial NOT NULL,
	product_name text NOT NULL,
	category_id int NOT NULL
);

CREATE TABLE categories (
	category_id serial NOT NULL,
	category_name varchar NOT NULL
);

CREATE TABLE cost (
	count int4 NOT NULL,
	product_id int4 NOT NULL,
	cost_sum numeric(20,6) NOT NULL,
	sell_sum numeric(20,6) NOT NULL,
	shift_date date NOT NULL
);

INSERT INTO categories
(category_name)
VALUES('Супы'),('Пиццы');

INSERT INTO products
(product_name, category_id)
VALUES('Борщ', 1),('Харчо',1),('4сыра',2),('Мясное Плато',2);

INSERT INTO cost
(product_id, count, cost_sum, sell_sum, shift_date)
VALUES
(1, 1, 30.252525, 100.292929, '2019-07-01'),
(1, 2, 60.740033, 200.000000, '2019-08-01'),
(2, 1, 30.252525, 100.292929, '2019-07-01'),
(2, 2, 30.252525, 100.292929, '2019-08-01'),
(3, 2, 300.890003, 900.504545, '2019-07-01'),
(3, 1, 150.492525, 450.292929, '2019-08-01'),
(4, 2, 300.252525, 1000.292929, '2019-07-01'),
(4, 4, 600.777777, 1850.292929, '2019-08-01');