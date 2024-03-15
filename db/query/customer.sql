-- name: GetCustomer :one
SELECT * FROM customer
WHERE customer_id = $1 LIMIT 1;


-- name: GetCustomerByEmail :one
SELECT * FROM customer
WHERE email = $1 LIMIT 1;

-- name: ListCustomers :many
SELECT * FROM customer
ORDER BY customer_id;

-- name: Createcustomer :one
INSERT INTO customer (
  first_name, last_name ,email ,password ,city,phone_number
) VALUES (
  $1, $2,$3,$4,$5,$6
)
RETURNING *;


-- name: Updatecustomer :one
UPDATE customer SET first_name = $2 , last_name = $3 , email = $4 , password=$5 , city = $6 , phone_number = $7
WHERE customer_id = $1
RETURNING *;


-- name: Deletecustomer :exec
DELETE FROM customer WHERE customer_id = $1;