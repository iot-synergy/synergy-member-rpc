/*
SQLyog Community v13.1.7 (64 bit)
MySQL - 8.0.26 
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `mms_reply` (
	`id` bigint (20) not null primary key auto_increment,
	`created_at` timestamp ,
	`updated_at` timestamp ,
	`reply` varchar (2048),
	`admin_id` varchar (64),
	`admin_name` varchar (128),
	`create_time` timestamp ,
	`update_time` timestamp ,
	`comment_id` bigint (20)
); 
