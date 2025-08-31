-- MySQL dump 10.13  Distrib 8.0.37, for Linux (x86_64)
--
-- Host: localhost    Database: auxtool
-- ------------------------------------------------------
-- Server version	8.0.37-0ubuntu0.22.04.3

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */
;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */
;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */
;
/*!50503 SET NAMES utf8mb4 */
;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */
;
/*!40103 SET TIME_ZONE='+00:00' */
;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */
;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */
;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */
;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */
;

--
-- Table structure for table `database`
--

DROP TABLE IF EXISTS `database`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `database` (
    `Id` bigint NOT NULL AUTO_INCREMENT,
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
) ENGINE = InnoDB AUTO_INCREMENT = 47 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */
;

--
-- Dumping data for table `database`
--

LOCK TABLES `database` WRITE;
/*!40000 ALTER TABLE `database` DISABLE KEYS */
;
INSERT INTO
    `database`
VALUES (
        33,
        '11',
        '百万数据集',
        '数值数据集',
        '2级',
        '2',
        '有',
        './火力分配/数值数据集/百万数据集/',
        '2',
        '火力分配'
    ),
    (
        42,
        '11',
        '小样本数据',
        '数值数据集',
        '3级',
        '22',
        '有',
        './火力分配/数值数据集/小样本数据/',
        '2',
        '火力分配'
    ),
    (
        44,
        '11',
        '1012',
        '数值数据集',
        '3级',
        '11',
        '有',
        './任务2/数值数据集/1012/',
        '11',
        '任务2'
    ),
    (
        45,
        '11',
        '车辆数据',
        '图像数据集',
        '2级',
        'int,png,jpg',
        '有',
        './自然图像/图像数据集/车辆数据/',
        '车辆检测数据集',
        '自然图像'
    ),
    (
        46,
        '11',
        '测试数据集',
        '数值数据集',
        '1级',
        '23',
        '有',
        './火力分配/数值数据集/测试数据集/',
        '32',
        '火力分配'
    );
/*!40000 ALTER TABLE `database` ENABLE KEYS */
;
UNLOCK TABLES;

--
-- Table structure for table `datatable`
--

DROP TABLE IF EXISTS `datatable`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `datatable` (
    `Id` int NOT NULL AUTO_INCREMENT,
    `Type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Task` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Dataset_Id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Table_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Header_num` int DEFAULT NULL,
    `Data_len` int DEFAULT NULL,
    `Data_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Csv_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 89 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */
;

--
-- Dumping data for table `datatable`
--

LOCK TABLES `datatable` WRITE;
/*!40000 ALTER TABLE `datatable` DISABLE KEYS */
;
INSERT INTO
    `datatable`
VALUES (
        84,
        '数值数据集',
        '火力分配',
        '33',
        'loss.csv',
        4,
        56,
        'float',
        '数值数据集/火力分配/33/loss.csv',
        '1723109845680'
    ),
    (
        85,
        '图像数据集',
        '自然图像',
        '45',
        '车辆检测数据',
        300000,
        1024,
        'jpg',
        '图像数据集/自然图像/45/loss.csv',
        '1723111228220'
    ),
    (
        86,
        '数值数据集',
        '火力分配',
        '33',
        'loss2.csv',
        4,
        16,
        'float',
        '数值数据集/火力分配/33/loss2.csv',
        '1724251690873'
    ),
    (
        87,
        '数值数据集',
        '火力分配',
        '42',
        'loss.csv',
        4,
        56,
        'float',
        '数值数据集/火力分配/42/loss.csv',
        '1724251725804'
    ),
    (
        88,
        '数值数据集',
        '火力分配',
        '33',
        'loss3.csv',
        4,
        16,
        'float',
        '数值数据集/火力分配/33/loss3.csv',
        '1724251972045'
    );
/*!40000 ALTER TABLE `datatable` ENABLE KEYS */
;
UNLOCK TABLES;

--
-- Table structure for table `example`
--

DROP TABLE IF EXISTS `example`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `example` (
    `Id` int NOT NULL AUTO_INCREMENT,
    `Example_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Rank` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `State` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Task` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Dataset_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Model_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Train_state` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Metics` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Network_num` int DEFAULT NULL,
    `Learning_rate` float(255, 0) DEFAULT NULL,
    `Act_function` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Radom_seed` int DEFAULT NULL,
    `Optimizer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `Batch_size` int DEFAULT NULL,
    `Epoch_num` int DEFAULT NULL,
    `Start_time` bigint DEFAULT NULL,
    `End_time` bigint DEFAULT NULL,
    `Explore_rate` float(255, 0) DEFAULT NULL,
    `Decay_factor` float(255, 0) DEFAULT NULL,
    `Example_id` int DEFAULT NULL,
    PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 64 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */
;

--
-- Dumping data for table `example`
--

LOCK TABLES `example` WRITE;
/*!40000 ALTER TABLE `example` DISABLE KEYS */
;
INSERT INTO
    `example`
VALUES (
        56,
        '小样本',
        '2级',
        '运行中',
        '领域1',
        'A领域',
        '123',
        '23213',
        '动作分布',
        'string',
        0,
        0,
        '',
        0,
        '',
        0,
        0,
        1727096927328,
        0,
        0,
        0,
        10
    ),
    (
        57,
        '大样本',
        '1级',
        '运行中',
        '领域1',
        'A领域',
        '22',
        '回归',
        '',
        'string',
        0,
        0,
        '',
        0,
        '',
        0,
        0,
        1727096986638,
        0,
        0,
        0,
        11
    ),
    (
        63,
        '测试',
        '1级',
        '运行中',
        '领域1',
        'A领域',
        '',
        '',
        '',
        'string',
        0,
        0,
        '',
        0,
        '',
        0,
        0,
        1727101773068,
        0,
        0,
        0,
        12
    );
/*!40000 ALTER TABLE `example` ENABLE KEYS */
;
UNLOCK TABLES;

--
-- Table structure for table `modelbase`
--

DROP TABLE IF EXISTS `modelbase`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `modelbase` (
    `Id` int NOT NULL AUTO_INCREMENT,
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
) ENGINE = InnoDB AUTO_INCREMENT = 24 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */
;

--
-- Dumping data for table `modelbase`
--

LOCK TABLES `modelbase` WRITE;
/*!40000 ALTER TABLE `modelbase` DISABLE KEYS */
;
INSERT INTO
    `modelbase`
VALUES (
        17,
        '1',
        '强化学习',
        '2级',
        'python',
        '强化学习/值迭代/17/2.png',
        '无',
        '无',
        '值迭代',
        '11',
        ''
    ),
    (
        20,
        '1233333',
        '深度学习',
        '1级',
        'python',
        '',
        '123',
        '2323',
        '卷积神经网络',
        '11',
        ''
    ),
    (
        21,
        '23213',
        '机器学习',
        '2级',
        'python',
        '',
        '312312',
        '3123123',
        '分类',
        '11',
        ''
    ),
    (
        22,
        '逻辑回归',
        '机器学习',
        '1级',
        'python',
        '',
        '1',
        '1',
        '回归',
        '11',
        ''
    ),
    (
        23,
        '岭回归',
        '机器学习',
        '1级',
        'python',
        '',
        '1',
        '1',
        '回归',
        '11',
        ''
    );
/*!40000 ALTER TABLE `modelbase` ENABLE KEYS */
;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `user` (
    `Id` int NOT NULL AUTO_INCREMENT,
    `UserName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `PassWord` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    PRIMARY KEY (`Id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */
;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */
;
INSERT INTO
    `user`
VALUES (1, 'ss', '123456'),
    (2, '123', '1234'),
    (4, '22222', '1234'),
    (5, 'username', 'password'),
    (6, 'sss', '123456'),
    (7, 'sk', '111'),
    (8, 'sssss', '111111'),
    (9, '456156', '123456');
/*!40000 ALTER TABLE `user` ENABLE KEYS */
;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */
;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */
;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */
;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */
;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */
;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */
;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */
;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */
;

-- Dump completed on 2024-09-23 22:35:06