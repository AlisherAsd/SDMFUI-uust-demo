CREATE TABLE pages (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL UNIQUE,
    value VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE widgets (
    id SERIAL PRIMARY KEY,
    mf_name VARCHAR(255) NOT NULL,
    component_name VARCHAR(255) NOT NULL,
    entry_url VARCHAR(255) NOT NULL
);


CREATE TABLE pages_widgets (
    id SERIAL PRIMARY KEY,
    page_id INT NOT NULL,
    widget_id INT NOT NULL,
    level  REAL NOT NULL,

    FOREIGN KEY (page_id) REFERENCES pages (id)
        ON DELETE CASCADE,

    FOREIGN KEY (widget_id) REFERENCES widgets (id)
        ON DELETE CASCADE
);