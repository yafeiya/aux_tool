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

 Date: 23/09/2024 20:22:24
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
) ENGINE = InnoDB AUTO_INCREMENT = 88 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of datatable
-- ----------------------------
INSERT INTO `datatable` VALUES (84, '数值数据集', '任务1', '33', 'loss.csv', 4, 56, 'float', '数值数据集/任务1/33/loss.csv', '1723109845680');
INSERT INTO `datatable` VALUES (85, '图像数据集', '自然图像', '45', '车辆检测数据', 300000, 1024, 'jpg', '图像数据集/自然图像/45/loss.csv', '1723111228220');
INSERT INTO `datatable` VALUES (86, '数值数据集', '任务1', '33', 'loss2.csv', 4, 16, 'float', '数值数据集/任务1/33/loss2.csv', '1724251690873');
INSERT INTO `datatable` VALUES (87, '数值数据集', '任务1', '42', 'loss.csv', 4, 56, 'float', '数值数据集/任务1/42/loss.csv', '1724251725804');
INSERT INTO `datatable` VALUES (88, '数值数据集', '任务1', '33', 'loss3.csv', 4, 16, 'float', '数值数据集/任务1/33/loss3.csv', '1724251972045');

-- ----------------------------
-- Table structure for example
-- ----------------------------
DROP TABLE IF EXISTS `example`;
CREATE TABLE `example`  (
  `Id` int(0) NOT NULL AUTO_INCREMENT,
  `Example_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Rank` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `State` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Task` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Dataset_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Model_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Train_state` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Metics` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Network_num` int(0) DEFAULT NULL,
  `Learning_rate` float(255, 0) DEFAULT NULL,
  `Act_function` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Radom_seed` int(0) DEFAULT NULL,
  `Optimizer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Batch_size` int(0) DEFAULT NULL,
  `Epoch_num` int(0) DEFAULT NULL,
  `Start_time` bigint(0) DEFAULT NULL,
  `End_time` bigint(0) DEFAULT NULL,
  `Explore_rate` float(255, 0) DEFAULT NULL,
  `Decay_factor` float(255, 0) DEFAULT NULL,
  `Example_id` int(0) DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 52 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of example
-- ----------------------------
INSERT INTO `example` VALUES (50, 'test1', '2级', '运行中', '领域1', 'A领域', '123', '111111111', '动作分布', 'string', 0, 0, '', 0, '', 0, 0, 1726147403682, 0, 0, 0, 2);
INSERT INTO `example` VALUES (51, '1', '1级', '运行中', '领域1', 'A领域', '22', '111111111', '动作分布', 'string', 233111, 0, '444', 111, 'kjkna', 0, 23, 1724398474208, 0, 5, 55, 1);
INSERT INTO `example` VALUES (52, 'testttttt', '1级', '运行中', '领域1', 'A领域', '', '', '', 'string', 0, 0, '', 0, '', 0, 0, 1726153545592, 0, 0, 0, 10);

-- ----------------------------
-- Table structure for modelbase
-- ----------------------------
DROP TABLE IF EXISTS `modelbase`;
CREATE TABLE `modelbase`  (
  `Id` int(0) NOT NULL AUTO_INCREMENT,
  `Dataset_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Rank` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Lan` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Data_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `Code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Task` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Released` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Image_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of modelbase
-- ----------------------------
INSERT INTO `modelbase` VALUES (10, '22222', '机器学习', '1级', 'python', '机器学习/回归/10/2.png', '1', '1', '回归', '11', '机器学习/回归/10/mega耿鬼白底.png');
INSERT INTO `modelbase` VALUES (12, '222', '机器学习', '1级', 'C++', '机器学习/回归/12/20.png', '11', '1', '回归', '11', '机器学习/回归/12/20.png');
INSERT INTO `modelbase` VALUES (13, '12222', '机器学习', '2级', 'C++', '机器学习/回归/13/1.xml', '1', '1', '回归', '11', '机器学习/回归/13/截图 2023-10-08 20-49-13.png');
INSERT INTO `modelbase` VALUES (16, '2222222', '机器学习', '1级', 'python', '机器学习/回归/16/mega耿鬼白底.png', '1', '1', '回归', '11', '');
INSERT INTO `modelbase` VALUES (17, '1', '强化学习', '2级', 'python', '强化学习/值迭代/17/2.png', '无', '无', '值迭代', '11', '');
INSERT INTO `modelbase` VALUES (18, '111111111', '机器学习', '1级', 'C++', '机器学习/回归/18/444.csv', '111', '111', '回归', '11', '');
INSERT INTO `modelbase` VALUES (19, '22', '机器学习', '1级', 'java', '机器学习/回归/19/111.csv', '23', '23', '回归', '11', '');
INSERT INTO `modelbase` VALUES (20, '1233333', '深度学习', '1级', 'python', '', '123', '2323', '卷积神经网络', '11', '');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `Id` int(0) NOT NULL AUTO_INCREMENT,
  `UserName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `PassWord` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

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
