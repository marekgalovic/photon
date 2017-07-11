CREATE TABLE models(
    `uid` VARCHAR(36) NOT NULL,
    `name` VARCHAR(30) NOT NULL,
    `owner` VARCHAR(50) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`uid`),
    INDEX (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
