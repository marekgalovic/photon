CREATE TABLE model_version_request_features(
    `model_version_uid` VARCHAR(36) NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `alias` VARCHAR(50),
    `required` INT(1) NOT NULL DEFAULT 0,
    FOREIGN KEY (`model_version_uid`) REFERENCES model_versions(`uid`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
