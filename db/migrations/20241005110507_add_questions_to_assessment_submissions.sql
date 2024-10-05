-- +goose Up
-- +goose StatementBegin
ALTER TABLE assessment_submissions ADD COLUMN question_id bigint NOT NULL REFERENCES questions(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP COLUMN question_id FROM assessment_submissions;
-- +goose StatementEnd
