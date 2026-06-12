-- You’re given a slow query in a PostgreSQL database for an e-commerce platform. The query fetches order details for a specific user,
-- but it’s taking too long with millions of orders. Here’s the query:
SELECT
  o.*, -- Select only needed columns, reduce data duplication
  p.*
FROM
  orders o -- Should filter here by sub-query first.
  JOIN order_items oi ON o.order_id = oi.order_id
  JOIN products p ON oi.product_id = p.product_id
WHERE
  o.user_id = 12345 -- Too late, filter after joining, waste of merging rows.
ORDER BY
  o.order_date DESC
  -- Missing LIMIT
;

-- Constraints: The orders table has 10M rows, order_items has 50M rows, and products has 1M rows.
--
-- Strategy 1: Critical indexes:
-- For sorting, limiting, and filtering by user.
CREATE INDEX idx_orders_user_date ON orders(user_id, order_date DESC);

-- These 3 indexes speed up the JOINs.
CREATE INDEX idx_order_items_order ON order_items(order_id);

CREATE INDEX idx_order_items_product ON order_items(product_id);

CREATE INDEX idx_products_id ON products(product_id);

-- Strategy 2: Sub-query to filter orders first to lower dataset.
