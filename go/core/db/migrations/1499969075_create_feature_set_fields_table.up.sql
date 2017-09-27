CREATE TABLE feature_set_fields(
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `feature_set_id` BIGINT NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `value_type` VARCHAR(10) NOT NULL,
    `nullable` INT(1) NOT NULL DEFAULT 1, 
    FOREIGN KEY (`feature_set_id`) REFERENCES feature_sets(`id`) ON DELETE CASCADE,
    INDEX (`feature_set_id`, `name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
