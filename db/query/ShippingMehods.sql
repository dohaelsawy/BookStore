-- name: GetShippingMehod :one
SELECT * FROM shipping_method
WHERE shipping_method_id = $1 LIMIT 1;

-- name: ListShippingMthods :many
SELECT * FROM shipping_method
ORDER BY shipping_method_id;

-- name: CreateShippingMehod :one
INSERT INTO shipping_method (
  name, cost
) VALUES (
  $1, $2
)
RETURNING *;


-- name: UpdateShippingMehod :one
UPDATE shipping_method SET name = $2 , cost = $3
WHERE shipping_method_id = $1
RETURNING *;


-- name: DeleteShippingMethod :exec
DELETE FROM shipping_method WHERE shipping_method_id = $1;