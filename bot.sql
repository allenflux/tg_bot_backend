/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80300
 Source Host           : localhost:3306
 Source Schema         : tg_bot

 Target Server Type    : MySQL
 Target Server Version : 80300
 File Encoding         : 65001

 Date: 02/04/2025 13:21:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bot
-- ----------------------------
DROP TABLE IF EXISTS `bot`;
CREATE TABLE `bot` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `greeting` varchar(255) NOT NULL,
  `greeting_status` int NOT NULL,
  `status` int NOT NULL,
  `photo` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
