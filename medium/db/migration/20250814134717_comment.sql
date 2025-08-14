-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Comments(
    id SERIAL PRIMARY KEY NOT NULL,
    blogId BIGINT UNSIGNED NOT NULL,
    userId BIGINT UNSIGNED NOT NULL,
    comment TEXT NOT NULL,
    upVote INT DEFAULT 0,
    downVote INT DEFAULT 0,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (blogId) REFERENCES Blogs(id) ON DELETE CASCADE,
    FOREIGN KEY (userId) REFERENCES Users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Comments;
-- +goose StatementEnd
