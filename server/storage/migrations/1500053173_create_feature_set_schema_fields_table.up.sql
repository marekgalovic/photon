CREATE TABLE feature_set_schema_fields(
    `feature_set_schema_uid` VARCHAR(36) NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `value_type` VARCHAR(10) NOT NULL,
    FOREIGN KEY (`feature_set_schema_uid`) REFERENCES feature_set_schemas(`uid`) ON DELETE CASCADE,
    INDEX (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
