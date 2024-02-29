/*
SQLyog Community v13.1.7 (64 bit)
MySQL - 8.0.26 
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `mms_members` (
	`id` char (108),
	`created_at` timestamp ,
	`updated_at` timestamp ,
	`status` tinyint (3),
	`username` varchar (765),
	`password` varchar (765),
	`nickname` varchar (765),
	`mobile` varchar (765),
	`email` varchar (765),
	`avatar` varchar (1536),
	`wechat_open_id` varchar (765),
	`expired_at` timestamp ,
	`rank_id` bigint (20)
); 
