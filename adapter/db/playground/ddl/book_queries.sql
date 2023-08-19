-- name: FindAllBooks :many
SELECT * FROM books;

-- name: FindBooksByAuthor :many
SELECT * FROM books b
WHERE b.author_id = sqlc.arg(author_id);

-- name: InsertBook :execlastid
INSERT INTO books(title, author_id)
VALUES (?, ?)