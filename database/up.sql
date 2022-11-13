DROP TABLE IF EXISTS students;

CREATE TABLE students(
    id VARCHAR(32) primary key,
    name VARCHAR(255) not null,
    age INTEGER not null
);