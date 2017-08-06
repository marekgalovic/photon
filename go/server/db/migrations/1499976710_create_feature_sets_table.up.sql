CREATE TABLE feature_sets(
    `uid` VARCHAR(36) NOT NULL,
    `name` VARCHAR(30) NOT NULL,
    `lookup_keys` TEXT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`uid`),
    INDEX (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;