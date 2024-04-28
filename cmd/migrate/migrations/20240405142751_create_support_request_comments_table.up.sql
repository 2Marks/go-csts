CREATE TABLE IF NOT EXISTS support_request_comments (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `support_request_id` INT UNSIGNED NOT NULL,
    `comment` VARCHAR(2000) NOT NULL,
    `commenter_id` INT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(`id`),
    FOREIGN KEY(`support_request_id`) REFERENCES support_requests(`id`),
    FOREIGN KEY(`commenter_id`) REFERENCES users(`id`)
)