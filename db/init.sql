CREATE TABLE `users` (
    `id` INT(8) unsigned NOT NULL AUTO_INCREMENT,
    `username` TEXT(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO users (username) VALUES
  ('Shunpoco'),
  ('Hoge'),
  ('Fuga');
