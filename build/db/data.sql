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
CREATE TABLE IF NOT EXISTS `permission` (
    `id` int AUTO_INCREMENT NOT NULL UNIQUE,
    `name` varchar(255) DEFAULT '40',
    PRIMARY KEY (`id`)
);
insert into permission (name) values ("profile");

CREATE TABLE IF NOT EXISTS `usr_perms` (
     `id_user` int DEFAULT '40',
     `id_perm` varchar(255) DEFAULT '40'
);
insert into usr_perms (id_user,id_perm) values ("1","1");


CREATE TABLE IF NOT EXISTS `specializations` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`name` varchar(255) NOT NULL UNIQUE,
	`prioweight` tinyint NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `auditory` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`name` varchar(255) NOT NULL,
	`spec_ids` int NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `subject` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`name` varchar(255) NOT NULL UNIQUE,
	`spec_ids` int NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `groups` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`name` varchar(255) NOT NULL,
	`year_id` int NOT NULL,
	`prevyeargrouplink` int NOT NULL,
	`tutor_id` int NOT NULL,
	`specialization_id` int NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `years` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`number` int NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `canonic_weeks` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`group_id` int NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `current_weeks` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`group_id` int NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `canonic_days` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`week_id` int NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `current_days` (
	`id` int AUTO_INCREMENT NOT NULL UNIQUE,
	`week_id` int NOT NULL,
	PRIMARY KEY (`id`)
);


ALTER TABLE `users` ADD CONSTRAINT `users_fk4` FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`);

ALTER TABLE `auditory` ADD CONSTRAINT `auditory_fk2` FOREIGN KEY (`spec_ids`) REFERENCES `specializations`(`id`);
ALTER TABLE `subject` ADD CONSTRAINT `subject_fk2` FOREIGN KEY (`spec_ids`) REFERENCES `specializations`(`id`);
ALTER TABLE `groups` ADD CONSTRAINT `groups_fk2` FOREIGN KEY (`year_id`) REFERENCES `years`(`id`);

ALTER TABLE `groups` ADD CONSTRAINT `groups_fk4` FOREIGN KEY (`tutor_id`) REFERENCES `users`(`id`);

ALTER TABLE `groups` ADD CONSTRAINT `groups_fk5` FOREIGN KEY (`specialization_id`) REFERENCES `specializations`(`id`);

ALTER TABLE `canonic_weeks` ADD CONSTRAINT `canonic_weeks_fk1` FOREIGN KEY (`group_id`) REFERENCES `groups`(`id`);
ALTER TABLE `current_weeks` ADD CONSTRAINT `current_weeks_fk1` FOREIGN KEY (`group_id`) REFERENCES `groups`(`id`);
ALTER TABLE `canonic_days` ADD CONSTRAINT `canonic_days_fk1` FOREIGN KEY (`week_id`) REFERENCES `canonic_weeks`(`id`);
ALTER TABLE `current_days` ADD CONSTRAINT `current_days_fk1` FOREIGN KEY (`week_id`) REFERENCES `current_weeks`(`id`);