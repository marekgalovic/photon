CREATE TABLE model_versions(
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `model_id` BIGINT NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `file_name` VARCHAR(100) NOT NULL,
    `is_primary` INT(1) NOT NULL DEFAULT 0,
    `is_shadow` INT(1) NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`model_id`) REFERENCES models(`id`) ON DELETE CASCADE,
    INDEX (`model_id`, `is_primary`),
    INDEX (`model_id`, `created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
