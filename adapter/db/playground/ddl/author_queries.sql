-- name: FindAllAuthors :many
SELECT * FROM authors;

-- name: InsertAuthors :execlastid
INSERT INTO authors(name)
VALUES (?);

-- name: FindAuthorsWithName :many
SELECT * FROM authors
WHERE name LIKE CONCAT('%', ?, '%');

-- name: FindAuthorByID :one
SELECT * FROM authors
WHERE id = ?;