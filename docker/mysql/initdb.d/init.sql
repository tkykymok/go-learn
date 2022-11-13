DROP DATABASE IF EXISTS go_app_test;
CREATE DATABASE go_app_test;
USE go_app_test;

-- todos
DROP TABLE IF EXISTS `todos`;
CREATE TABLE `todos`
(
    `id`         INT         NOT NULL AUTO_INCREMENT,
    `title`      VARCHAR(45) NULL,
    `completed`  TINYINT(1)  NOT NULL DEFAULT 0,
    `userId`     INT         NOT NULL,
    `deleted`    TINYINT(1)  NOT NULL DEFAULT 0,
    `created_at` DATETIME    NULL,
    PRIMARY KEY (`id`)
);
INSERT INTO `todos` (`id`, `title`, `completed`, `userId`, `deleted`, `created_at`)
VALUES ('1', 'テスト1', '0', '1', '0', '2022-01-01');
INSERT INTO `todos` (`id`, `title`, `completed`, `userId`, `deleted`, `created_at`)
VALUES ('2', 'テスト2', '0', '1', '0', '2022-01-01');
INSERT INTO `todos` (`id`, `title`, `completed`, `userId`, `deleted`, `created_at`)
VALUES ('3', 'テスト3', '0', '1', '0', '2022-01-01');

-- users
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`   INT         NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(45) NULL,
    PRIMARY KEY (`id`)
);
INSERT INTO `users` (`id`, `name`)
VALUES ('1', '山田');
