CREATE TABLE IF NOT EXISTS users (
    id int PRIMARY KEY AUTO_INCREMENT,
    firstName varchar(255),
    lastName varchar(255),
    email varchar(255) UNIQUE NOT NULL,
    password longtext,
    role ENUM ('retailer', 'customer')
);

CREATE TABLE IF NOT EXISTS shops (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `shop_name` varchar(255) UNIQUE NOT NULL,
    `description` varchar(255),
    `address` varchar(255) UNIQUE NOT NULL,
    `phone_number` varchar(14),
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `user_id` int
);

CREATE TABLE IF NOT EXISTS shop_availability (
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
    `shop_id` int,
    `user_id` int DEFAULT NULL,
    `user_email` varchar(255) NOT NULL
);

ALTER TABLE `appointments`
    ADD FOREIGN KEY (`shop_id`) REFERENCES `shops` (`id`),
    ADD UNIQUE KEY `appointment_date_time` (`appointment_date_time`),
    ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `shop_availability` ADD FOREIGN KEY (`shop_id`) REFERENCES `shops` (`id`);

ALTER TABLE `shops` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

INSERT INTO `users` (`firstName`, `lastName`, `email`, `password`, `role`) VALUES
    ('testeur1', 'testeur1', 'testeur1@test.fr', 'testeur1', 'retailer'),
    ('testeur2', 'testeur2', 'testeur2@test.fr', 'testeur2', 'retailer'),
    ('testeur3', 'testeur3', 'testeur3@test.fr', 'testeur3', 'customer');

INSERT INTO `shops` (`shop_name`, `description`, `address`, `phone_number`, `created_at`, `user_id`) VALUES
    ('Dentest', 'dentiste test', '1 rue des dentistes, 00000 TestCity', '00 00 00 00 00',  '2023-02-03 16:02:34', 1),
    ('orltest', 'ORL test', '1 rue des orls, 00000 TestCity', '00 00 00 00 01', '2023-02-28 11:00:00', 2),
    ('shop3', 'shop sans appointments', '1 rue des shop3, 00000 TestCity', '00 00 00 00 02', '2023-02-28 11:00:00', 2);

INSERT INTO `shop_availability` (`shop_id`, `day_of_week`, `duration`, `start_time`, `end_time`) VALUES
    (1, 'tuesday', 15, '09:00:00', '19:00:00'),
    (1, 'wednesday', 15, '09:00:00', '19:00:00'),
    (1, 'thursday', 15, '09:00:00', '17:00:00'),
    (1, 'friday', 15, '09:00:00', '17:00:00'),
    (1, 'saturday', 15, '09:00:00', '19:00:00'),

    (2, 'monday', 30, '09:00:00', '19:00:00'),
    (2, 'tuesday', 30, '09:00:00', '19:00:00'),
    (2, 'wednesday', 30, '09:00:00', '19:00:00'),
    (2, 'thursday', 30, '09:00:00', '17:00:00'),
    (2, 'friday', 30, '09:00:00', '17:00:00'),
    (2, 'saturday', 30, '09:00:00', '19:00:00');

INSERT INTO `appointments` (`customer_name`, `appointment_date`, `appointment_time`, `appointment_date_time` , `shop_id`, `user_id`, `user_email`) VALUES
    ('testeur_sans_compte_1', '2023-03-03', '09:30:00', CONCAT(appointment_date, ' ', appointment_time), 1, null, 'testeur_sans_compte_1@test.fr'),
    ('testeur3 testeur3', '2023-03-03', '10:30:00', CONCAT(appointment_date, ' ', appointment_time), 1, 3, 'testeur3@test.fr'),
    ('test_sans_compte_2', '2023-03-03', '12:30:00', CONCAT(appointment_date, ' ', appointment_time), 1, null, 'testeur_sans_compte_2@test.fr'),
    ('testeur_sans_compte', '2023-03-03', '13:30:00', CONCAT(appointment_date, ' ', appointment_time), 2, null, 'testeur_sans_compte_@test.fr'),
    ('testeur3 testeur3', '2023-03-03', '14:30:00', CONCAT(appointment_date, ' ', appointment_time), 2, 3, 'testeur3@test.fr');
