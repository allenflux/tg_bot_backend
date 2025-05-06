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

 Date: 02/04/2025 13:22:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `central_control_id` int NOT NULL,
  `tg_link` varchar(255) NOT NULL,
  `type` int NOT NULL,
  `size` int NOT NULL,
  `bot_size` int NOT NULL,
  `role_size` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of group
-- ----------------------------
BEGIN;
INSERT INTO `group` VALUES (1, '111', 1, 'aaa', 1, 11, 11, 11);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
