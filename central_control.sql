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

 Date: 02/04/2025 13:21:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for central_control
-- ----------------------------
DROP TABLE IF EXISTS `central_control`;
CREATE TABLE `central_control` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `domain` varchar(255) NOT NULL,
  `number_of_customers` int NOT NULL DEFAULT '0',
  `number_of_business` int NOT NULL,
  `note` varchar(255) DEFAULT NULL,
  `status` int NOT NULL DEFAULT '0',
  `secret_key` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of central_control
-- ----------------------------
BEGIN;
INSERT INTO `central_control` VALUES (1, '22', '2222', 20, 2, '2', 1, 'aaa');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
