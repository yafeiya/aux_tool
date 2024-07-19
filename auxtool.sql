/*
 Navicat Premium Data Transfer


 Source Server         : aux
 Source Server Type    : MySQL
 Source Server Version : 50743 (5.7.43-log)
 Source Host           : localhost:3306
 Source Schema         : auxtool

 Target Server Type    : MySQL
 Target Server Version : 50743 (5.7.43-log)
 File Encoding         : 65001

 Date: 06/07/2024 13:53:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for database
-- ----------------------------
DROP TABLE IF EXISTS `database`;
CREATE TABLE `database`  (

  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `Released` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Dataset_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Rank` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Character_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Header` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Data_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Description` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `Task` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;


-- ----------------------------
-- Records of database
-- ----------------------------
INSERT INTO `database` VALUES (1, '11', '波士顿房价数据集0', '数值数据集', '1级', 'Utf-8', '无', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值', '任务1');
INSERT INTO `database` VALUES (2, '11', '波士顿房价数据集1', '数值数据集', '1级', 'Utf-8', '无', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集', '任务2');
INSERT INTO `database` VALUES (3, '11', '波士顿房价数据集2', '数值数据集', '1级', 'Utf-8', '无', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集', '任务2');
INSERT INTO `database` VALUES (4, '11', '波士顿房价数据集3', '数值数据集', '1级', 'Utf-8', '无', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集', '任务1');
INSERT INTO `database` VALUES (5, '11', '波士顿房价数据集4', '数值数据集', '1级', 'Utf-8', '无', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集', '任务1');
INSERT INTO `database` VALUES (11, '11', '测试1', '数值数据集', '2级', 'uf8', '有', './任务1/数值数据集/测试1/', '测试', '任务1');


-- ----------------------------
-- Table structure for datatable
-- ----------------------------
DROP TABLE IF EXISTS `datatable`;
CREATE TABLE `datatable`  (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `Type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Task` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Dataset_Id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Table_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Header_num` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Data_len` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Data_type` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `Csv_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 42 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;


-- ----------------------------
-- Records of datatable
-- ----------------------------

INSERT INTO `datatable` VALUES (34, '数值数据集', '任务2', '2', 'loss00.csv', '4', '56', 'float', '数值数据集/任务2/2/loss00.csv', '1720180348382');
INSERT INTO `datatable` VALUES (35, '数值数据集', '任务1', '1', '测试1.csv', '4', '56', 'float', '数值数据集/任务1/1/测试1.csv', '1720180551766');
INSERT INTO `datatable` VALUES (36, '数值数据集', '任务2', '2', '测试1.csv', '4', '56', 'float', '数值数据集/任务2/2/测试1.csv', '1720182186844');
INSERT INTO `datatable` VALUES (38, '数值数据集', '任务1', '5', '测试1.csv', '4', '56', 'float', '数值数据集/任务1/5/测试1.csv', '1720185024170');
INSERT INTO `datatable` VALUES (39, '数值数据集', '任务1', '4', '测试2.csv', '4', '56', 'float', '数值数据集/任务1/4/测试2.csv', '1720186551867');
INSERT INTO `datatable` VALUES (40, '数值数据集', '任务1', '11', '测试1.csv', '4', '56', 'float', '数值数据集/任务1/11/测试1.csv', '1720186776891');
INSERT INTO `datatable` VALUES (41, '数值数据集', '任务1', '11', '测试2.csv', '4', '56', 'float', '数值数据集/任务1/11/测试2.csv', '1720186784790');

-- ----------------------------
-- Table structure for example
-- ----------------------------
DROP TABLE IF EXISTS `example`;
CREATE TABLE `example`  (
  `Id` int(11) NOT NULL,
  `Example_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Rank` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `State` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Cpu_num` int(11) NULL DEFAULT NULL,
  `Gpu_num` int(11) NULL DEFAULT NULL,
  `Post_date` datetime NULL DEFAULT NULL,
  `Dataset_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Model_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Epoch_num` int(11) NULL DEFAULT NULL,
  `Model_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Loss` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Optimizer` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Decay` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Evalution` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Model_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Memory` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Start_time` datetime NULL DEFAULT NULL,
  `End_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of example
-- ----------------------------
INSERT INTO `example` VALUES (1, '方案2', '1级', '已终止', 1, 10, '2000-01-01 00:00:00', './data', 'YOLOv3', 10000, '目标检测', '奖励分布', 'SGD', '0.1', 'mAP', './model', '2000MB', '2000-01-01 00:00:00', '2000-01-05 00:00:00');


-- ----------------------------
-- Table structure for modelbase
-- ----------------------------
DROP TABLE IF EXISTS `modelbase`;
CREATE TABLE `modelbase`  (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Dataset_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Rank` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Lan` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Data_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Description` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `Code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Task` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `Released` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

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

  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `UserName` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `PassWord` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;


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
