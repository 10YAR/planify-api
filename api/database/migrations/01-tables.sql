USE `planify`;

CREATE TABLE IF NOT EXISTS `users` (
                        `id` int PRIMARY KEY AUTO_INCREMENT,
                        `firstName` varchar(255),
                        `lastName` varchar(255),
                        `email` varchar(255),
                        `password` longtext,
                        `role` ENUM ('retailer', 'customer')
);

CREATE TABLE IF NOT EXISTS `shops` (
                         `id` int PRIMARY KEY AUTO_INCREMENT,
                         `shop_name` varchar(255) UNIQUE NOT NULL,
                         `address` varchar(255),
                         `created_at` timestamp NOT NULL,
                         `user_id` int
);

CREATE TABLE IF NOT EXISTS `shop_availability` (
                                     `id` int PRIMARY KEY AUTO_INCREMENT,
                                     `shop_id` int,
                                     `day_of_week` ENUM ('monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday', 'sunday'),
                                     `duration` int,
                                     `start_time` time,
                                     `end_time` time
);

CREATE TABLE IF NOT EXISTS `appointments` (
                               `id` int PRIMARY KEY AUTO_INCREMENT,
                               `customer_name` varchar(255),
                               `appointment_date` date,
                               `appointment_time` time,
                               `appointment_date_time` datetime,
                               `shop_id` int
);

ALTER TABLE `appointments` ADD FOREIGN KEY (`shop_id`) REFERENCES `shops` (`id`);

ALTER TABLE `shop_availability` ADD FOREIGN KEY (`shop_id`) REFERENCES `shops` (`id`);

ALTER TABLE `shops` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

INSERT INTO `users` (`id`, `firstName`, `lastName`, `email`, `password`, `role`) VALUES
    (1, 'Thomas', 'Evano', 'test', 'test', 'retailer');

INSERT INTO `shops` (`id`, `shop_name`, `address`, `created_at`, `user_id`) VALUES
    (1, 'test', 'test', '2023-02-03 16:02:34', 1);

INSERT INTO `shop_availability` (`id`, `shop_id`, `day_of_week`, `duration`, `start_time`, `end_time`) VALUES
    (1, 1, 'tuesday', 30, '09:00:00', '19:00:00'),
    (2, 1, 'wednesday', 30, '09:00:00', '19:00:00'),
    (3, 1, 'thursday', 15, '09:00:00', '17:00:00'),
    (4, 1, 'friday', 15, '09:00:00', '17:00:00'),
    (5, 1, 'saturday', 30, '09:00:00', '19:00:00');

INSERT INTO `appointments` (`id`, `customer_name`, `appointment_date`, `appointment_time`, `appointment_date_time` , `shop_id`) VALUES
    (1, 'test', '2023-03-03', '09:30:00', CONCAT(appointment_date, ' ', appointment_time), 1);