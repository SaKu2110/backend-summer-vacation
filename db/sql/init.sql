DROP DATABASE IF EXISTS summer;
CREATE DATABASE summer;
USE summer;

---- drop ----
DROP TABLE IF EXISTS `users`;

---- create ----
CREATE TABLE IF NOT EXISTS `users`
(
 `id`       VARCHAR(50) NOT NULL,
 `password` VARCHAR(64) NOT NULL,
 PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
