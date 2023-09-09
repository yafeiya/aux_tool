-- MySQL dump 10.13  Distrib 5.7.43, for Win64 (x86_64)
--
-- Host: localhost    Database: auxtool
-- ------------------------------------------------------
-- Server version	5.7.43-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `database`
--

DROP TABLE IF EXISTS `database`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `database` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `Released` varchar(255) DEFAULT NULL,
  `Dataset_name` varchar(255) DEFAULT NULL,
  `Type` varchar(255) DEFAULT NULL,
  `Rank` varchar(255) DEFAULT NULL,
  `Character_type` varchar(255) DEFAULT NULL,
  `Header` varchar(255) DEFAULT NULL,
  `Data_path` varchar(255) DEFAULT NULL,
  `Description` longtext,
  `Task` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `database`
--

LOCK TABLES `database` WRITE;
/*!40000 ALTER TABLE `database` DISABLE KEYS */;
INSERT INTO `database` VALUES (1,'01','波士顿房价数据集1','数值数据集','1级','Utf-8','无','D:\\\\01_dataset\\\\波士顿房价数据集.csv','波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。','任务1'),(2,'01','波士顿房价数据集0','数值数据集','1级','Utf-8','无','D:\\\\01_dataset\\\\波士顿房价数据集.csv','波士顿房价数据集是一个经典的回归问题数据集','任务2'),
,(1,'11','波士顿房价数据集1','数值数据集','1级','Utf-8','无','D:\\\\01_dataset\\\\波士顿房价数据集.csv','波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。','任务1'),(2,'01','波士顿房价数据集0','数值数据集','1级','Utf-8','无','D:\\\\01_dataset\\\\波士顿房价数据集.csv','波士顿房价数据集是一个经典的回归问题数据集','任务2'),;


/*!40000 ALTER TABLE `database` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `example`
--

DROP TABLE IF EXISTS `example`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `example` (
  `Id` int(11) NOT NULL,
  `Example_name` varchar(255) DEFAULT NULL,
  `Rank` varchar(255) DEFAULT NULL,
  `State` varchar(255) DEFAULT NULL,
  `Cpu_num` int(11) DEFAULT NULL,
  `Gpu_num` int(11) DEFAULT NULL,
  `Post_date` datetime DEFAULT NULL,
  `Dataset_url` varchar(255) DEFAULT NULL,
  `Model_name` varchar(255) DEFAULT NULL,
  `Epoch_num` int(11) DEFAULT NULL,
  `Model_type` varchar(255) DEFAULT NULL,
  `Loss` varchar(255) DEFAULT NULL,
  `Optimizer` varchar(255) DEFAULT NULL,
  `Decay` varchar(255) DEFAULT NULL,
  `Evalution` varchar(255) DEFAULT NULL,
  `Model_url` varchar(255) DEFAULT NULL,
  `Memory` varchar(255) DEFAULT NULL,
  `Start_time` datetime DEFAULT NULL,
  `End_time` datetime DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `example`
--

LOCK TABLES `example` WRITE;
/*!40000 ALTER TABLE `example` DISABLE KEYS */;
/*!40000 ALTER TABLE `example` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `modelbase`
--

DROP TABLE IF EXISTS `modelbase`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `modelbase` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Dataset_name` varchar(255) DEFAULT NULL,
  `Type` varchar(255) DEFAULT NULL,
  `Rank` varchar(255) DEFAULT NULL,
  `Lan` varchar(255) DEFAULT NULL,
  `Data_path` varchar(255) DEFAULT NULL,
  `Description` longtext,
  `Code` varchar(255) DEFAULT NULL,
  `Task` varchar(255) DEFAULT NULL,
  `Released` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `modelbase`
--

LOCK TABLES `modelbase` WRITE;
/*!40000 ALTER TABLE `modelbase` DISABLE KEYS */;
INSERT INTO `modelbase` VALUES (1,'KNN模型','机器学习','1级','C++','D:\\\\01_dataset\\\\波士顿房价数据集.csv','波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。','import...','回归','11'),(2,'CNN模型','机器学习','1级','C++','D:\\\\01_dataset\\\\波士顿房价数据集.csv','波士顿房价数据集是一个经典的回归问题数据集，包含了1970年代末期波士顿的南部郊区共506个房屋样本。这些数据都是在20世纪70年代末期以旧金山海湾区房价的情况为基准，采集了波士顿房价相关的特征。这些特征包括了房屋所在城市的犯罪率、每个城镇平均房间数以及自有住房比例等。每个特征的值都已经经过了预处理，例如处理了缺失值和异常值。','import...','回归','11');
/*!40000 ALTER TABLE `modelbase` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `UserName` varchar(255) DEFAULT NULL,
  `PassWord` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'ss','123456'),(2,'123','1234'),(4,'22222','1234'),(5,'username','password'),(6,'sss','123456'),(7,'sk','111'),(8,'sssss','111111'),(9,'1234','123456'),(10,'admin','123456');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-09-07  4:40:50
