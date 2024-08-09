/*
 Navicat Premium Data Transfer

 Source Server         : local_mysql
 Source Server Type    : MySQL
 Source Server Version : 80039
 Source Host           : localhost:3306
 Source Schema         : auxtool

 Target Server Type    : MySQL
 Target Server Version : 80039
 File Encoding         : 65001

 Date: 08/08/2024 18:03:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for database
-- ----------------------------
DROP TABLE IF EXISTS `database`;
CREATE TABLE `database`  (
  `Id` bigint(0) NOT NULL AUTO_INCREMENT,
  `Released` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Dataset_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Rank` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Character_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Header` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Data_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `Task` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 45 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of database
-- ----------------------------
INSERT INTO `database` VALUES (33, '11', '123', '数值数据集', '2级', '2', '有', './任务1/数值数据集/123/', '2', '任务1');
INSERT INTO `database` VALUES (42, '11', '22', '数值数据集', '3级', '22', '有', './任务1/数值数据集/22/', '2', '任务1');
INSERT INTO `database` VALUES (44, '11', '1012', '数值数据集', '3级', '11', '有', './任务2/数值数据集/1012/', '11', '任务2');
INSERT INTO `database` VALUES (45, '11', '车辆数据', '图像数据集', '2级', 'int,png,jpg', '有', './自然图像/图像数据集/车辆数据/', '车辆检测数据集', '自然图像');

SET FOREIGN_KEY_CHECKS = 1;
