CREATE TABLE model_features(
    `model_id` BIGINT NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `alias` VARCHAR(50) NOT NULL,
    `required` INT(1) NOT NULL DEFAULT 0,
    FOREIGN KEY (`model_id`) REFERENCES models(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
