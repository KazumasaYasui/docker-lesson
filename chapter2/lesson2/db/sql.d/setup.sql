CREATE USER 'sampleapp'@'%' IDENTIFIED BY 'sampleapp_pw';
GRANT ALL PRIVILEGES ON sampleapp_development.* TO 'sampleapp'@'%';
FLUSH PRIVILEGES;

DROP DATABASE IF EXISTS sampleapp_development;
CREATE DATABASE sampleapp_development;

USE sampleapp_development;

SET CHARACTER_SET_CLIENT = utf8mb4;
SET CHARACTER_SET_CONNECTION = utf8mb4;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id bigint(20) NOT NULL AUTO_INCREMENT,
  name varchar(255) DEFAULT NULL,
  email varchar(255) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

INSERT INTO sampleapp_development.users (
  id, name, email, created_at, updated_at
) VALUES
  (1, '前田くん', 'sample1@example.com', '2021-04-12 12:34:56', '2021-04-12 12:34:56'),
  (2, '井草くん', 'sample2@example.com', '2021-04-12 12:34:56', '2021-04-12 12:34:56'),
  (3, '岡島くん', 'sample3@example.com', '2021-04-12 12:34:56', '2021-04-12 12:34:56'),
  (4, '関根くん', 'sample4@example.com', '2021-04-12 12:34:56', '2021-04-12 12:34:56');
