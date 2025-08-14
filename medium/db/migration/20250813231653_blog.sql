-- +goose Up
-- +goose StatementBegin
CREATE TABLE Blogs (
    id SERIAL PRIMARY KEY NOT NULL,
    userId BIGINT UNSIGNED NOT NULL,
    blogTitle VARCHAR(255) NOT NULL,
    blogImage VARCHAR(255),
    blogContent TEXT NOT NULL,
    upVote INT DEFAULT 0,
    downVote INT DEFAULT 0 ,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (userId) REFERENCES Users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Blogs;
-- +goose StatementEnd
