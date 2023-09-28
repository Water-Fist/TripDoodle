-- 트리거 제거
DROP TRIGGER update_user_updated_at ON users;
DROP TRIGGER update_sight_updated_at ON sights;
DROP TRIGGER update_post_updated_at ON posts;
DROP TRIGGER update_comment_updated_at ON comments;

-- 함수 제거
DROP FUNCTION update_updated_at_column;

-- 테이블 제거 (참조하는 순서에 따라)
DROP TABLE comments;
DROP TABLE posts;
DROP TABLE sights;
DROP TABLE users;
