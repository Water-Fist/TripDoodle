CREATE TABLE sights (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        latitude REAL NOT NULL,
                        longitude REAL NOT NULL,
                        area BOOLEAN NOT NULL,
                        is_deleted BOOLEAN NOT NULL,
                        deleted_at TIMESTAMP,
                        created_at TIMESTAMP NOT NULL,
                        updated_at TIMESTAMP NOT NULL
);

CREATE TABLE posts (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       content TEXT NOT NULL,
                       image_url VARCHAR(255),
                       state BOOLEAN NOT NULL,
                       is_deleted BOOLEAN NOT NULL,
                       deleted_at TIMESTAMP,
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL,
                       sight_id INT,
                       FOREIGN KEY (sight_id) REFERENCES sights(id)
);

CREATE TABLE user (

)