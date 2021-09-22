CREATE TABLE IF NOT EXISTS `cards` (
    `id` INT(8) unsigned NOT NULL AUTO_INCREMENT,
    `title` TEXT(255) NOT NULL,
    `describe` TEXT(255) NOT NULL,
    `reference` TEXT(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `cards` (`title`, `describe`, `reference`) VALUES
  ('test', 'test-describe', 'test-reference'),
  ('test', 'test-describe', 'test-reference');
