-- Active: 1713744276875@@127.0.0.1@5432@psql_db
-- DROP DATABASE IF EXISTS psql_db;

CREATE DATABASE psql_db;

-- \c psql_db;

CREATE TABLE nationality(
    nationality_id SERIAL PRIMARY KEY,
    nationality_name VARCHAR(50) NOT NULL,
    nationality_code CHAR(2) NOT NULL
);

INSERT INTO nationality(nationality_name, nationality_code)
VALUES 
    ('Indonesia', 'ID'), 
    ('Malaysia', 'MY'), 
    ('Singapore', 'SG'), 
    ('Saudi Arabia', 'SA'), 
    ('United States', 'US');


CREATE TABLE customer(
    cst_id SERIAL PRIMARY KEY,
    nationality_id INT NOT NULL,
    cst_name VARCHAR(50) NOT NULL,
    cst_dob DATE NOT NULL,
    cst_phonenum VARCHAR(20) NOT NULL,
    cst_email VARCHAR(50) NOT NULL,
    CONSTRAINT fk_nationality FOREIGN KEY(nationality_id) REFERENCES nationality(nationality_id)
);

INSERT INTO customer(nationality_id, cst_name, cst_dob, cst_phonenum, cst_email)
VALUES 
(1, 'joko', '1992-05-19', '62812345678', 'joko@gmail.com'),
(3, 'budi', '1998-03-28', '62898765342', 'budi@gmail.com'),
(5, 'sinta', '1990-12-01', '62818726261', 'sinta@gmail.com');


CREATE TABLE family_list(
    fl_id SERIAL PRIMARY KEY,
    cst_id INT NOT NULL,
    fl_relation VARCHAR(50) NOT NULL,
    fl_name VARCHAR(50) NOT NULL,
    fl_dob VARCHAR(50) NOT NULL,
    CONSTRAINT fk_customer FOREIGN KEY(cst_id) REFERENCES customer(cst_id)
);

INSERT INTO family_list(cst_id, fl_relation, fl_name, fl_dob)
VALUES
(1, 'ayah', 'tono', '1950-11-20'),
(2, 'ayah', 'mustopo', '1975-08-06'),
(2, 'ibu', 'sumiati', '1980-01-14'),
(1, 'ibu', 'ratih', '1952-09-16'),
(1, 'adik', 'tora', '2000-06-28'),
(3, 'kakak', 'trisna', '1990-03-27'),
(1, 'kakak', 'siska', '1970-08-09');