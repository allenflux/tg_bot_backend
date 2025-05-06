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

 Date: 01/04/2025 13:51:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `role` int NOT NULL DEFAULT '0',
  `status` int NOT NULL DEFAULT '1',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, '11', '1122', 10, 1, '2025-03-11 12:56:13', 'aa');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
