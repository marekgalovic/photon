CREATE TABLE model_precomputed_features(
    `model_id` BIGINT NOT NULL,
    `feature_set_id` BIGINT NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `alias` VARCHAR(50) NOT NULL,
    `required` INT(1) NOT NULL DEFAULT 0,
    FOREIGN KEY (`model_id`) REFERENCES models(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`feature_set_id`) REFERENCES feature_sets(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
