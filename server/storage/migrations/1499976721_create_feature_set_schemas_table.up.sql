CREATE TABLE feature_set_schemas(
    `uid` VARCHAR(36) NOT NULL,
    `feature_set_uid` VARCHAR(36) NOT NULL,
    `schema` TEXT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`uid`),
    FOREIGN KEY (`feature_set_uid`) REFERENCES feature_sets(`uid`) ON DELETE CASCADE,
    INDEX (`feature_set_uid`, `updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
