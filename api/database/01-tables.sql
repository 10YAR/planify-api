-- Last update: 2021-02-03 12:49

USE `planify`;

CREATE TABLE IF NOT EXISTS `users` (
                        `id` int PRIMARY KEY AUTO_INCREMENT,
                        `firstName` varchar(255),
                        `lastName` varchar(255),
                        `email` varchar(255) UNIQUE NOT NULL,
                        `password` longtext,
                        `role` ENUM ('retailer', 'customer')
);

CREATE TABLE IF NOT EXISTS `shops` (
                         `id` int PRIMARY KEY AUTO_INCREMENT,
                         `shop_name` varchar(255) UNIQUE NOT NULL,
                         `description` varchar(255),
                         `address` varchar(255) UNIQUE NOT NULL,
                         `phone_number` varchar(14),
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

ALTER TABLE `appointments`
    ADD FOREIGN KEY (`shop_id`) REFERENCES `shops` (`id`),
    ADD UNIQUE KEY `appointment_date_time` (`appointment_date_time`);

ALTER TABLE `shop_availability` ADD FOREIGN KEY (`shop_id`) REFERENCES `shops` (`id`);

ALTER TABLE `shops` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

INSERT INTO `users` (`firstName`, `lastName`, `email`, `password`, `role`) VALUES
    ('Thomas', 'Evano', 'thomas@mail.fr', 'thomas', 'retailer'),
    ('Diyar', 'Bayrakli', 'dyia@mail.fr', 'dyiar', 'retailer');

INSERT INTO `shops` (`shop_name`, `description`, `address`, `phone_number`, `created_at`, `user_id`) VALUES
    ('Docteur Thomas', 'Médecin généraliste', '2 rue des médecins, 75000 Paris', '01 01 01 01 01', '2023-02-03 16:00:00', 1),
    ('Docteur Diyar', 'Opthalmo', '1 rue des ophtalmo, 75000 Paris', '01 01 01 01 01', '2023-02-03 16:02:34', 1);

INSERT INTO `shop_availability` (`id`, `shop_id`, `day_of_week`, `duration`, `start_time`, `end_time`) VALUES
    (1, 1, 'tuesday', 30, '09:00:00', '19:00:00'),
    (2, 1, 'wednesday', 30, '09:00:00', '19:00:00'),
    (3, 1, 'thursday', 15, '09:00:00', '17:00:00'),
    (4, 1, 'friday', 15, '09:00:00', '17:00:00'),
    (5, 1, 'saturday', 30, '09:00:00', '19:00:00');

INSERT INTO `appointments` (`customer_name`, `appointment_date`, `appointment_time`, `appointment_date_time` , `shop_id`) VALUES
    ('Diyar Bayrakli', '2023-03-03', '09:30:00', CONCAT(appointment_date, ' ', appointment_time), 1);