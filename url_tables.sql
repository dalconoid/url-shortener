CREATE TABLE urls (
                      id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                      slug TEXT CONSTRAINT unique_slug UNIQUE,
                      url TEXT CONSTRAINT unique_url UNIQUE,
                      created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);