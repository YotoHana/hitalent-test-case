-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS answer(
    id SERIAL PRIMARY KEY,
    question_id INTEGER NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_question
        FOREIGN KEY (question_id)
        REFERENCES question(id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS answer;
-- +goose StatementEnd
