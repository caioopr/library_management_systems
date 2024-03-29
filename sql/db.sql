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