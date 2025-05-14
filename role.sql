/*
 Navicat Premium Dump SQL

 Source Server         : local_docker_mysql
 Source Server Type    : MySQL
 Source Server Version : 90300 (9.3.0)
 Source Host           : localhost:3306
 Source Schema         : tg_bot

 Target Server Type    : MySQL
 Target Server Version : 90300 (9.3.0)
 File Encoding         : 65001

 Date: 14/05/2025 09:12:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `bot_id` int NOT NULL,
  `cmd` varchar(255) DEFAULT NULL,
  `user_size` int DEFAULT NULL,
  `bot_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` (`id`, `name`, `bot_id`, `cmd`, `user_size`, `bot_name`) VALUES (1, 'Admin', 1, 'no', 1, '测试数据');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
