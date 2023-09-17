-- 트리거 제거
DROP TRIGGER update_user_updated_at ON "user";
DROP TRIGGER update_sight_updated_at ON sight;
DROP TRIGGER update_post_updated_at ON post;
DROP TRIGGER update_comment_updated_at ON comment;

-- 함수 제거
DROP FUNCTION update_updated_at_column;

-- 테이블 제거 (참조하는 순서에 따라)
DROP TABLE comment;
DROP TABLE post;
DROP TABLE sight;
DROP TABLE "user";
