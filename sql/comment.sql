/*
SQLyog Community v13.1.7 (64 bit)
MySQL - 8.0.26 
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `mms_comment` (
	`id` bigint (20) not null primary key auto_increment,
	`created_at` timestamp ,
	`updated_at` timestamp ,
	`title` varchar (768),
	`content` varchar (6144),
	`member_id` varchar (192),
	`create_time` timestamp ,
	`update_time` timestamp ,
	`is_reply` tinyint (1)
); 
