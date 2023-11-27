-- name: GetOrder :one
SELECT * FROM customer_order
WHERE customer_order_id = $1 LIMIT 1;

-- name: ListCustomerOrder :many
SELECT * FROM customer_order
ORDER BY customer_order_id;

-- name: CreateCustomerOrder :one
INSERT INTO customer_order (
  order_date, customer_id ,shipping_method_id
) VALUES (
  $1, $2,$3
)
RETURNING *;


-- name: UpdateCustomerOrder :one
UPDATE customer_order SET order_date = $2 , customer_id = $3 , shipping_method_id = $4 
WHERE customer_order_id = $1
RETURNING *;


-- name: DeleteCustomerOrder :exec
DELETE FROM customer_order WHERE customer_order_id = $1;