-- Filename: migrations/000001_create_courses_table.up.sql
CREATE TABLE courses (
    course_code varchar(50),
    course_title text NOT NULL,
    course_credits BIGINT NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    version int NOT NULL DEFAULT 1
);