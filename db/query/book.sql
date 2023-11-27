-- name: GetBook :one
SELECT * FROM book
WHERE book_id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM book
ORDER BY book_id;

-- name: CreateBook :one
INSERT INTO book (
  name, publish_date ,price ,sku ,description
) VALUES (
  $1, $2,$3,$4,$5
)
RETURNING *;


-- name: UpdateBook :one
UPDATE book SET name = $2 , publish_date = $3 , price = $4 , sku=$5 , description = $6
WHERE book_id = $1
RETURNING *;


-- name: DeleteBook :exec
DELETE FROM book WHERE book_id = $1;