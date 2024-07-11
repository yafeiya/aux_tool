/*
 Navicat Premium Data Transfer

 Source Server         : server
 Source Server Type    : MySQL
 Source Server Version : 80029 (8.0.29)
 Source Host           : 49.234.4.144:3306
 Source Schema         : auxtool

 Target Server Type    : MySQL
 Target Server Version : 80029 (8.0.29)
 File Encoding         : 65001

 Date: 11/07/2024 15:38:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for database
-- ----------------------------
DROP TABLE IF EXISTS `database`;
CREATE TABLE `database`  (
  `Id` bigint NOT NULL AUTO_INCREMENT,
  `Released` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Dataset_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Rank` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Character_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Header` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Data_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `Task` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 45 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of database
-- ----------------------------
INSERT INTO `database` VALUES (33, '11', '123', '数值数据集', '2级', '2', '有', './任务1/数值数据集/123/', '2', '任务1');
INSERT INTO `database` VALUES (42, '11', '22', '数值数据集', '3级', '22', '有', './任务1/数值数据集/22/', '2', '任务1');
INSERT INTO `database` VALUES (44, '11', '1012', '数值数据集', '3级', '11', '有', './任务2/数值数据集/1012/', '11', '任务2');

-- ----------------------------
-- Table structure for datatable
-- ----------------------------
DROP TABLE IF EXISTS `datatable`;
CREATE TABLE `datatable`  (
  `Id` int NOT NULL AUTO_INCREMENT,
  `Type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Task` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Dataset_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Table_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Header_num` int NULL DEFAULT NULL,
  `Data_len` int NULL DEFAULT NULL,
  `Data_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Csv_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 81 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of datatable
-- ----------------------------
INSERT INTO `datatable` VALUES (73, '数值数据集', '任务1', '33', '888.csv', 4, 56, 'float', '数值数据集/任务1/33/888.csv', '1697360644668');
INSERT INTO `datatable` VALUES (78, '数值数据集', '任务1', '33', '444.csv', 4, 56, 'float', '数值数据集/任务1/33/444.csv', '1697365482093');
INSERT INTO `datatable` VALUES (80, '数值数据集', '任务1', '33', '111.csv', 4, 56, 'float', '数值数据集/任务1/33/111.csv', '1698133150550');

-- ----------------------------
-- Table structure for example
-- ----------------------------
DROP TABLE IF EXISTS `example`;
CREATE TABLE `example`  (
  `Id` int NOT NULL AUTO_INCREMENT,
  `Example_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Rank` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `State` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Cpu_num` int NULL DEFAULT NULL,
  `Gpu_num` int NULL DEFAULT NULL,
  `Post_data` bigint NULL DEFAULT NULL,
  `Dataset_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Model_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Epoch_num` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Model_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Loss` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Optimizer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Decay` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Evaluation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Model_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Memory` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Start_time` bigint NULL DEFAULT NULL,
  `End_time` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of example
-- ----------------------------
INSERT INTO `example` VALUES (31, '2', '2级', '运行中', 4, 1, 0, 'http://127.0.0.1:5173/A领域/领域1/2', '', '200e', '', 'loss', 'optimizer', 'decay', 'evaluation', 'model_url', '2000M', 1698487875903, 0);
INSERT INTO `example` VALUES (32, '1', '1级', '运行中', 4, 1, 0, 'http://127.0.0.1:5173/A领域/领域1/1', '', '200e', '', 'loss', 'optimizer', 'decay', 'evaluation', 'model_url', '2000M', 1698495197411, 0);

-- ----------------------------
-- Table structure for modelbase
-- ----------------------------
DROP TABLE IF EXISTS `modelbase`;
CREATE TABLE `modelbase`  (
  `Id` int NOT NULL AUTO_INCREMENT,
  `Dataset_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Rank` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Lan` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Data_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `Code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Task` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Released` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Image_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of modelbase
-- ----------------------------
INSERT INTO `modelbase` VALUES (10, '22222', '机器学习', '1级', 'python', '机器学习/回归/10/444.csv', '1', '1', '回归', '11', '机器学习/回归/10/1.png');
INSERT INTO `modelbase` VALUES (12, '222', '机器学习', '1级', 'C++', '机器学习/回归/12/20.png', '11', '1', '回归', '11', '机器学习/回归/12/20.png');
INSERT INTO `modelbase` VALUES (13, '12222', '机器学习', '2级', 'C++', '机器学习/回归/13/123.csv', '1', '1', '回归', '11', '机器学习/回归/13/截图 2023-10-08 20-49-13.png');
INSERT INTO `modelbase` VALUES (16, '2222222', '机器学习', '1级', 'python', '机器学习/回归/16/mega耿鬼白底.png', '1', '1', '回归', '11', '');
INSERT INTO `modelbase` VALUES (17, '1', '强化学习', '2级', 'python', '', '无', '无', '值迭代', '11', '');
INSERT INTO `modelbase` VALUES (18, '111111111', '机器学习', '1级', 'C++', '机器学习/回归/18/444.csv', '111', '111', '回归', '11', '');
INSERT INTO `modelbase` VALUES (19, '22', '机器学习', '1级', 'java', '机器学习/回归/19/111.csv', '23', '23', '回归', '11', '');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `Id` int NOT NULL AUTO_INCREMENT,
  `UserName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `PassWord` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'ss', '123456');
INSERT INTO `user` VALUES (2, '123', '1234');
INSERT INTO `user` VALUES (4, '22222', '1234');
INSERT INTO `user` VALUES (5, 'username', 'password');
INSERT INTO `user` VALUES (6, 'sss', '123456');
INSERT INTO `user` VALUES (7, 'sk', '111');
INSERT INTO `user` VALUES (8, 'sssss', '111111');
INSERT INTO `user` VALUES (9, '456156', '123456');

SET FOREIGN_KEY_CHECKS = 1;
