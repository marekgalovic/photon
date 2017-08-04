CREATE TABLE model_version_precomputed_features(
    `model_version_uid` VARCHAR(36) NOT NULL,
    `feature_set_uid` VARCHAR(36) NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `alias` VARCHAR(50) NOT NULL,
    `required` INT(1) NOT NULL DEFAULT 0,
    FOREIGN KEY (`model_version_uid`) REFERENCES model_versions(`uid`) ON DELETE CASCADE,
    FOREIGN KEY (`feature_set_uid`) REFERENCES feature_sets(`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
