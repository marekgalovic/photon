CREATE TABLE model_versions(
    `uid` VARCHAR(36) NOT NULL,
    `model_uid` VARCHAR(36) NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `is_primary` INT(1) NOT NULL DEFAULT 0,
    `is_shadow` INT(1) NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`uid`),
    FOREIGN KEY (`model_uid`) REFERENCES models(`uid`) ON DELETE CASCADE,
    UNIQUE KEY (`model_uid`, `is_primary`),
    INDEX (`model_uid`, `created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
