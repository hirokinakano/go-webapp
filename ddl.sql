CREATE DATABASE sample;
USE sample;
CREATE TABLE users(
  id int PRIMARY KEY AUTO_INCREMENT,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100) NOT NULL
);
INSERT INTO users(
  first_name,
  last_name
)
VALUES
(
  'Taro',
  'Yamada'
),
(
  'Jiro',
  'Sato'
)
