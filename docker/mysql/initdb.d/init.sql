DROP DATABASE IF EXISTS go_app_test;
CREATE DATABASE go_app_test;
USE go_app_test;

DROP TABLE IF EXISTS `todos`;
CREATE TABLE `todos`
(
    `id`         INT         NOT NULL,
    `title`      VARCHAR(45) NULL,
    `completed`  TINYINT(1)  NOT NULL DEFAULT 0,
    `userId`     INT         NOT NULL,
    `created_at` DATETIME    NULL,
    PRIMARY KEY (`id`)
);
INSERT INTO `todos` (`id`, `title`, `completed`, `userId`, `created_at`)
VALUES ('1', 'テスト1', '0', '1', '2022-01-01');
INSERT INTO `todos` (`id`, `title`, `completed`, `userId`, `created_at`)
VALUES ('2', 'テスト2', '0', '1', '2022-01-01');
INSERT INTO `todos` (`id`, `title`, `completed`, `userId`, `created_at`)
VALUES ('3', 'テスト3', '0', '1', '2022-01-01');

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`   INT         NOT NULL,
    `name` VARCHAR(45) NULL,
    PRIMARY KEY (`id`)
);
INSERT INTO `users` (`id`, `name`)
VALUES ('1', '山田');
