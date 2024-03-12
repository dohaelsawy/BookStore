-- name: GetOrderLine :one
SELECT * FROM order_line
WHERE line_id = $1 LIMIT 1;

-- name: ListOrderLines :many
SELECT * FROM order_line
ORDER BY line_id;

-- name: CreateOrderLine :one
INSERT INTO order_line (
  order_id, book_id ,price
) VALUES (
  $1, $2,$3
)
RETURNING *;


-- name: UpdateOrderLine :one
UPDATE order_line SET order_id = $2 , book_id = $3 , price = $4 
WHERE line_id = $1
RETURNING *;


-- name: DeleteOrderLine :exec
DELETE FROM order_line WHERE line_id = $1;