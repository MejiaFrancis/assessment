-- Filename: migrations/000001_create_courses_table.up.sql
CREATE TABLE courses (
    course_id bigserial PRIMARY KEY,
    course_code text NOT NULL,
    course_numbercredits BIGINT NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);