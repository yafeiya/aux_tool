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

 Date: 08/08/2024 18:03:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for datatable
-- ----------------------------
DROP TABLE IF EXISTS `datatable`;
CREATE TABLE `datatable`  (
  `Id` int(0) NOT NULL AUTO_INCREMENT,
  `Type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Task` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Dataset_Id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Table_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Header_num` int(0) DEFAULT NULL,
  `Data_len` int(0) DEFAULT NULL,
  `Data_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Csv_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 83 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of datatable
-- ----------------------------
INSERT INTO `datatable` VALUES (84, '数值数据集', '任务1', '33', 'loss.csv', 4, 56, 'float', '数值数据集/任务1/33/loss.csv', '1723109845680');
INSERT INTO `datatable` VALUES (85, '图像数据集', '自然图像', '45', '车辆检测数据', 300000, 1024, 'jpg', '图像数据集/自然图像/45/loss.csv', '1723111228220');

SET FOREIGN_KEY_CHECKS = 1;
