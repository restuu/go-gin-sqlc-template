-- name: FindAllAuthors :many
SELECT * FROM authors;

-- name: InsertAuthors :execlastid
INSERT INTO authors(name)
VALUES (?)