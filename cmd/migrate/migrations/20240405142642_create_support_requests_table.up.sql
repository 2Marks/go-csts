CREATE TABLE IF NOT EXISTS support_requests (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `subject` VARCHAR(500) NOT NULL,
    `description` VARCHAR(2000) NOT NULL,
    `status` ENUM('pending', 'processing', 'closed') NOT NULL DEFAULT 'pending',
    `creator_id` INT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    `closed_at` TIMESTAMP, 

    PRIMARY KEY(`id`),
    FOREIGN KEY (`creator_id`) REFERENCES users(`id`)
)