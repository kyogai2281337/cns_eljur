create database if not exists journal;
use journal;
CREATE TABLE IF NOT EXISTS `roles` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`name` varchar(255) NOT NULL DEFAULT '50',
    PRIMARY KEY (`id`)
);
insert into roles (name) values ("superuser");
insert into roles (name) values ("user");
CREATE TABLE IF NOT EXISTS `users` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`email` varchar(255) NOT NULL DEFAULT '40' UNIQUE,
	`encrypted_password` varchar(255) NOT NULL DEFAULT '100',
	`is_active` bool DEFAULT '1',
	`role_id` int,
	`first_name` varchar(255) DEFAULT '40',
    `last_name` varchar(255) DEFAULT '40',
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `cabinets` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`name` varchar(255) NOT NULL UNIQUE,
	`type` int NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `groups` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`name` varchar(10) NOT NULL UNIQUE,
	`spec_id` int NOT NULL,
	`max_pairs` int NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `specializations` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`name` varchar(50) NOT NULL UNIQUE,
	`course` int NOT NULL,
	`plan_id` varchar(40) NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `subjects` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`name` varchar(50) NOT NULL UNIQUE,
	`type` int NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `teachers` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`name` varchar(100) NOT NULL UNIQUE,
	`links_id` varchar(255) NOT NULL,
	`capacity` int NOT NULL,
	PRIMARY KEY (`id`)
);

ALTER TABLE `groups` ADD CONSTRAINT `groups_fk2` FOREIGN KEY (`spec_id`) REFERENCES `specializations`(`id`);
ALTER TABLE `users` ADD CONSTRAINT `users_fk4` FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`);