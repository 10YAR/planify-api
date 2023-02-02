USE `planify`;

CREATE TABLE IF NOT EXISTS `retailer` (
                            `id` int PRIMARY KEY AUTO_INCREMENT,
                            `retailer_name` varchar(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS `shops` (
                         `id` int PRIMARY KEY AUTO_INCREMENT,
                         `shop_name` varchar(255) UNIQUE NOT NULL,
                         `created_at` timestamp NOT NULL,
                         `retailer_id` int
);

CREATE TABLE IF NOT EXISTS `shop_availability` (
                                     `id` int PRIMARY KEY AUTO_INCREMENT,
                                     `shop_id` int,
                                     `availability_date` date,
                                     `start_time` datetime,
                                     `end_time` datetime
);

CREATE TABLE IF NOT EXISTS `appointment` (
                               `id` int PRIMARY KEY AUTO_INCREMENT,
                               `customer_name` varchar(255),
                               `appointment_date` date,
                               `start_time` datetime,
                               `end_time` datetime,
                               `status` boolean,
                               `shop_id` int
);

ALTER TABLE `appointment` ADD FOREIGN KEY (`shop_id`) REFERENCES `shops` (`id`);

ALTER TABLE `shop_availability` ADD FOREIGN KEY (`shop_id`) REFERENCES `shops` (`id`);

ALTER TABLE `shops` ADD FOREIGN KEY (`retailer_id`) REFERENCES `retailer` (`id`);