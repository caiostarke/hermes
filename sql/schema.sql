CREATE TABLE IF NOT EXISTS study_case (
    id BIGSERIAL PRIMARY KEY, 
    name VARCHAR(255) NOT NULL,
    tags JSONB,
    comment TEXT,
    description TEXT,
    next_review DATE,
    created_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS flash_cards (
    id BIGSERIAL PRIMARY KEY, 
    front TEXT,
    back TEXT,
    next_review DATE,
    study_case_id BIGSERIAL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    FOREIGN KEY (study_case_id) REFERENCES study_case(id)
);

CREATE TABLE IF NOT EXISTS ingested_texts {
    id biserial primary key, 
    text TEXT not null,
    created_at TIMESTAMP DEFAULT current_timestamp
};