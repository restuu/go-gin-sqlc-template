CREATE TABLE IF NOT EXISTS books (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'primary key',
    title VARCHAR(255) NOT NULL COMMENT 'book title',
    author_id BIGINT NOT NULL COMMENT 'book\'s author reference ID',
    FOREIGN KEY (author_id) REFERENCES authors (id)
)