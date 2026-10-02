CREATE TABLE remotes (
    id SERIAL PRIMARY KEY,
    page VARCHAR(255) NOT NULL,
    mf_name VARCHAR(255) NOT NULL,
    component_name VARCHAR(255) NOT NULL,
    entry_url VARCHAR(255) NOT NULL
)