/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80034 (8.0.34)
 Source Host           : localhost:3306
 Source Schema         : auxtool

 Target Server Type    : MySQL
 Target Server Version : 80034 (8.0.34)
 File Encoding         : 65001

 Date: 18/09/2023 14:56:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for database
-- ----------------------------
DROP TABLE IF EXISTS `database`;
CREATE TABLE `database`  (
  `Id` bigint NOT NULL AUTO_INCREMENT,
  `Released` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Dataset_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Rank` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Character_type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Header` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Data_path` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Description` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `Task` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of database
-- ----------------------------
INSERT INTO `database` VALUES (1, '11', '波士顿房价数据集0', '数值数据集', '1级', 'Utf-8', '无', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值', '任务1');
INSERT INTO `database` VALUES (2, '11', '波士顿房价数据集1', '数值数据集', '1级', 'Utf-8', '无', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集', '任务2');
INSERT INTO `database` VALUES (3, '11', '波士顿房价数据集2', '数值数据集', '1级', 'Utf-8', '无', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集', '任务2');
INSERT INTO `database` VALUES (4, '11', '波士顿房价数据集3', '数值数据集', '1级', 'Utf-8', '无', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集', '任务1');
INSERT INTO `database` VALUES (5, '11', '波士顿房价数据集4', '数值数据集', '1级', 'Utf-8', '无', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集', '任务1');
INSERT INTO `database` VALUES (10, '11', '11', '11', '11', '11', '11', '11', '11', '11');

-- ----------------------------
-- Table structure for datatable
-- ----------------------------
DROP TABLE IF EXISTS `datatable`;
CREATE TABLE `datatable`  (
  `Id` bigint NOT NULL AUTO_INCREMENT,
  `Type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Task` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Dataset_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Table_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Header_num` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Data_len` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Data_type` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `Csv_path` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Time` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of datatable
-- ----------------------------
INSERT INTO `datatable` VALUES (1, '任务1', '数值数据集', '波士顿房价数据集4', 'loss.csv', '56', '4', 'float', './任务1/数值数据集/波士顿房价数据集4/loss.csv', '1');
INSERT INTO `datatable` VALUES (30, '数值数据集', '任务1', '波士顿房价数据集0', 'loss4.csv', '56', '4', 'float', './数值数据集/任务1/波士顿房价数据集0/loss4.csv', '1695019977463');
INSERT INTO `datatable` VALUES (31, '数值数据集', '任务1', '波士顿房价数据集0', 'loss.csv', '56', '4', 'float', './数值数据集/任务1/波士顿房价数据集0/loss.csv', '1695020077151');
INSERT INTO `datatable` VALUES (32, '数值数据集', '任务1', '波士顿房价数据集0', 'loss2.csv', '56', '4', 'float', './数值数据集/任务1/波士顿房价数据集0/loss2.csv', '1695020099277');
INSERT INTO `datatable` VALUES (33, '数值数据集', '任务1', '波士顿房价数据集3', 'loss2.csv', '56', '4', 'float', './数值数据集/任务1/波士顿房价数据集3/loss2.csv', '1695020106342');

-- ----------------------------
-- Table structure for example
-- ----------------------------
DROP TABLE IF EXISTS `example`;
CREATE TABLE `example`  (
  `Id` int NOT NULL,
  `Example_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Rank` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `State` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Cpu_num` int NULL DEFAULT NULL,
  `Gpu_num` int NULL DEFAULT NULL,
  `Post_date` datetime NULL DEFAULT NULL,
  `Dataset_url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Model_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Epoch_num` int NULL DEFAULT NULL,
  `Model_type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Loss` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Optimizer` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Decay` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Evalution` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Model_url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Memory` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Start_time` datetime NULL DEFAULT NULL,
  `End_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of example
-- ----------------------------

-- ----------------------------
-- Table structure for modelbase
-- ----------------------------
DROP TABLE IF EXISTS `modelbase`;
CREATE TABLE `modelbase`  (
  `Id` int NOT NULL AUTO_INCREMENT,
  `Dataset_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Rank` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Lan` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Data_path` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Description` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `Code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Task` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `Released` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of modelbase
-- ----------------------------
INSERT INTO `modelbase` VALUES (1, '多项式回归', '机器学习', '1级', 'C++', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。', 'import...', '回归', '11');
INSERT INTO `modelbase` VALUES (2, 'DQN', '强化学习', '1级', 'C++', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。', 'import...', '值迭代', '11');
INSERT INTO `modelbase` VALUES (3, 'A3C', '强化学习', '1级', 'C++', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。', 'import...', '策略学习', '11');
INSERT INTO `modelbase` VALUES (4, 'FCN', '深度学习', '1级', 'C++', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。', 'import...', '卷积神经网络', '11');
INSERT INTO `modelbase` VALUES (5, 'LSTM', '深度学习', '1级', 'C++', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。', 'import...', '循环神经网络', '11');
INSERT INTO `modelbase` VALUES (6, 'SVM', '机器学习', '1级', 'C++', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。', 'import...', '分类', '11');
INSERT INTO `modelbase` VALUES (7, '蚁群算法', '机器学习', '1级', 'C++', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。', 'import...', '优化', '11');
INSERT INTO `modelbase` VALUES (8, '贝叶斯算法', '机器学习', '1级', 'C++', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。', 'import...', '优化', '11');
INSERT INTO `modelbase` VALUES (9, '粒子种群算法', '机器学习', '1级', 'C++', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。', 'import...', '优化', '11');
INSERT INTO `modelbase` VALUES (10, '决策树', '机器学习', '1级', 'C++', 'D:\\\\01_dataset\\\\波士顿房价数据集.csv', '波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。', 'import...', '分类', '11');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `Id` int NOT NULL AUTO_INCREMENT,
  `UserName` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `PassWord` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

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
INSERT INTO `user` VALUES (9, '1234', '123456');
INSERT INTO `user` VALUES (10, 'admin', '123456');

SET FOREIGN_KEY_CHECKS = 1;
