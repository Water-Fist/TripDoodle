CREATE TABLE sights (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        type VARCHAR(255) NOT NULL,
                        province VARCHAR(255) NOT NULL,
                        city_county_district VARCHAR(255) NOT NULL,
                        legal_dong VARCHAR(255) NOT NULL,
                        ri VARCHAR(255) NOT NULL,
                        street_number VARCHAR(255) NOT NULL,
                        building_number VARCHAR(255) NOT NULL,
                        latitude REAL NOT NULL,
                        longitude REAL NOT NULL,
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