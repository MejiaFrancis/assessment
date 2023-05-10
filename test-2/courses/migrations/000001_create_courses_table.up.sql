-- Filename: migrations/000001_create_courses_table.up.sql
CREATE TABLE courses (
    id bigserial PRIMARY KEY, 
    course_code varchar(50),
    course_title text NOT NULL,
    course_credit BIGINT NOT NULL,
    version int NOT NULL DEFAULT 1,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
    
);