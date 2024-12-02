-- CREATE SCHEMA IF NOT EXISTS LIB_SYSTEM;

-- SET search_path TO LIB_SYSTEM;

-- Cria a tabela Usuarios (com a chave estrangeira ON DELETE CASCADE)
-- CREATE TABLE IF NOT EXISTS users (
--   id SERIAL PRIMARY KEY,
--   nome VARCHAR(100) NOT NULL,
--   nickname VARCHAR(100) NOT NULL,
--   email VARCHAR(100) NOT NULL UNIQUE,
--   password VARCHAR(100) NOT NULL,
--   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
-- );
CREATE DATABASE IF NOT EXISTS lib_system;
USE lib_system;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nickname varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    created_at timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE authors(
    id int auto_increment primary key,
    name varchar(50) not null,
);

CREATE TABLE books(
    id int auto_increment primary key,
    title varchar(150) not null,
	avarage_rating DECIMAL(4,2),
	isbn             varchar(10) not null,
	isbn13           varchar(13) not null,
	language_code     varchar(10) not null,
	num_pages         int not null default 0,
	ratings_count     int not null default 0,
	text_reviews_count int not null default 0,
	publication_date timestamp,
	publisher   varchar(150) not null     
) ENGINE=INNODB;

CREATE TABLE books_authors(
    id int auto_increment primary key,
    author_id int,
    book_id int,
    FOREIGN KEY (author_id) REFERENCES authors (id),
    FOREIGN KEY (book_id) REFERENCES users (id)
);

