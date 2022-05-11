CREATE TABLE IF NOT EXISTS scopes(
    `id` BINARY(16) NOT NULL,
    `resource_domain_name` VARCHAR(127) NOT NULL,
    `resource_name` VARCHAR(127) NOT NULL,
    `name` VARCHAR(127) NOT NULL,
    `type` ENUM('public', 'priavte') NOT NULL,
    `ctime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `mtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    UNIQUE(`name`),
    INDEX(`resource_domain_name`),
    INDEX(`type`)
) ENGINE InnoDB COLLATE 'utf8mb4_unicode_ci' CHARACTER SET 'utf8mb4';
