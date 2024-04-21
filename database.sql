-- Active: 1713612460274@@127.0.0.1@5432@psql_db
DROP DATABASE IF EXISTS psql_db;

CREATE DATABASE psql_db;

\c psql_db;

CREATE TABLE nationality(
    nationality_id SERIAL PRIMARY KEY,
    nationality_name VARCHAR(50) NOT NULL,
    nationality_code CHAR(2) NOT NULL
);

INSERT INTO nationality(nationality_name, nationality_code)
VALUES ('Indonesia', 'ID');

CREATE TABLE customer(
    cst_id SERIAL PRIMARY KEY,
    nationality_id INT NOT NULL,
    cst_name CHAR(50) NOT NULL,
    cst_dob DATE NOT NULL,
    cst_phonenum VARCHAR(20) NOT NULL,
    cst_email varchar(50) NOT NULL,
    CONSTRAINT fk_nationality FOREIGN KEY(nationality_id) REFERENCES nationality(nationality_id)
);

INSERT INTO customer(nationality_id, cst_name, cst_dob, cst_phonenum, cst_email)
VALUES (1, 'joko', '1999-12-23', '62813829728', 'joko@gmail.com')

CREATE TABLE family_list(
    fl_id SERIAL PRIMARY KEY,
    cst_id INT NOT NULL,
    fl_relation VARCHAR(50) NOT NULL,
    fl_name VARCHAR(50) NOT NULL,
    fl_dob VARCHAR(50) NOT NULL,
    CONSTRAINT fk_customer FOREIGN KEY(cst_id) REFERENCES customer(cst_id)
);

INSERT INTO family_list(cst_id, fl_relation, fl_name, fl_dob)
VALUES (1, 'ayah', 'budi', '1970-4-16')