-- name: GetPaymentTransaction :one
SELECT * FROM transaction_payment
WHERE transaction_id = $1 LIMIT 1;

-- name: ListPaymentTranscation :many
SELECT * FROM transaction_payment
ORDER BY transaction_id;

-- name: CreatePaymentTranscation :one
INSERT INTO transaction_payment (
  amount, line_id ,shipping_method_id ,card_number ,expiration_date ,cvv ,status ,customer_id
) VALUES (
  $1, $2,$3,$4,$5,$6,$7,$8
)
RETURNING *;



-- name: DeletePaymentTranscation :exec
DELETE FROM transaction_payment WHERE transaction_id = $1;