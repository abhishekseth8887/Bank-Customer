-- +migrate Up

CREATE TABLE `customers` (
    `id` bigint AUTO_INCREMENT NOT NULL,
    `first_name` varchar(255) NOT NULL,
    `middle_name` varchar(255),
    `last_name` varchar(255) NOT NULL,
    `dob` varchar(255) NOT NULL,
    `mobile_number` varchar(255) NOT NULL,
    `gender` varchar(255) NOT NULL,
    `customer_number` varchar(255) NOT NULL,
    `country_of_birth` varchar(255) NOT NULL,
    `country_of_residence` varchar(255) NOT NULL,
    `customer_segment` varchar(255) NOT NULL,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`),
    KEY `index_name` (`first_name`, `middle_name`, `last_name`),
    KEY `index_mobile` (`mobile_number`),
    KEY `index_created_at` (`created_at`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- +migrate Down

DROP TABLE IF EXISTS `customers`;
