/*
SQLyog Community v13.1.7 (64 bit)
MySQL - 8.0.26 
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `mms_members` (
	`id` char (36) not null primary key,
	`created_at` timestamp ,
	`updated_at` timestamp ,
	`status` tinyint (3),
	`username` varchar (255),
	`password` varchar (255),
	`nickname` varchar (255),
	`mobile` varchar (255),
	`email` varchar (255),
	`avatar` varchar (512),
	`wechat_open_id` varchar (255),
	`expired_at` timestamp ,
	`rank_id` bigint (20)
); 
