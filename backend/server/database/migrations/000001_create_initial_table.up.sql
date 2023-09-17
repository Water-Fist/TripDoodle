CREATE TABLE "user" (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        nickname VARCHAR(255) NOT NULL UNIQUE,
                        email VARCHAR(255) NOT NULL UNIQUE,
                        password VARCHAR(255) NOT NULL,
                        is_deleted BOOLEAN NOT NULL,
                        deleted_at TIMESTAMP,
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sight (
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
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE post (
                      id SERIAL PRIMARY KEY,
                      title VARCHAR(255) NOT NULL,
                      content TEXT NOT NULL,
                      image_url VARCHAR(255),
                      state BOOLEAN NOT NULL,
                      is_deleted BOOLEAN NOT NULL,
                      deleted_at TIMESTAMP,
                      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      sight_id INT,
                      user_id INT,
                      FOREIGN KEY (sight_id) REFERENCES sight(id),
                      FOREIGN KEY (user_id) REFERENCES "user"(id)
);

CREATE TABLE comment (
                         id SERIAL PRIMARY KEY,
                         content TEXT NOT NULL,
                         is_deleted BOOLEAN NOT NULL,
                         deleted_at TIMESTAMP,
                         created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         post_id INT,
                         user_id INT,
                         parent_comment_id INT,
                         FOREIGN KEY (post_id) REFERENCES post(id),
                         FOREIGN KEY (parent_comment_id) REFERENCES comment(id),
                         FOREIGN KEY (user_id) REFERENCES "user"(id)
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user_updated_at
    BEFORE UPDATE ON "user"
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_sight_updated_at
    BEFORE UPDATE ON sight
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_post_updated_at
    BEFORE UPDATE ON post
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_comment_updated_at
    BEFORE UPDATE ON comment
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
