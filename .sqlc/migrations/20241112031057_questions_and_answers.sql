-- +goose Up
-- +goose StatementBegin
CREATE TABLE questions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    question_text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE company_question_answers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    company_id UUID NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    question_id UUID NOT NULL REFERENCES questions(id),
    answer_text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(company_id, question_id)
);

CREATE INDEX idx_company_answers_company ON company_question_answers(company_id);
CREATE INDEX idx_company_answers_question ON company_question_answers(question_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE company_question_answers;
DROP TABLE questions;
-- +goose StatementEnd