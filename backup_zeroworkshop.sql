-- MySQL dump 10.13  Distrib 5.7.16, for Linux (x86_64)
--
-- Host: 172.16.150.199    Database: zeroworkshop
-- ------------------------------------------------------
-- Server version	5.7.17

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
-- Table structure for table `ACN`
--

DROP TABLE IF EXISTS `ACN`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ACN` (
  `CID` int(10) NOT NULL DEFAULT '0',
  `ACCNAME` varchar(180) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `BAL` decimal(12,2) DEFAULT NULL,
  `BALAVL` decimal(12,2) DEFAULT NULL,
  `PROVINCE` int(2) DEFAULT NULL,
  `TRNO` int(11) DEFAULT NULL,
  `MPH` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`CID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ACN`
--

LOCK TABLES `ACN` WRITE;
/*!40000 ALTER TABLE `ACN` DISABLE KEYS */;
INSERT INTO `ACN` VALUES (12345678,'นาย ทดสอบ ระบบ',10000.00,10000.00,10,0,'0891234567'),(1000000001,'นาง สมศรี มีชัย',200350.00,200350.00,50,0,'891234567'),(1000000002,'นาง สมศรี มีชัย',200850.00,200850.00,50,0,'891234567'),(1000000003,'นาง สมศรี มีชัย',200000.00,200000.00,50,0,'891234567'),(1122334455,'นาย ปีเตอร์ แพน',8735.00,8735.00,10,0,'877972751'),(1122334456,'นาย ปีเตอร์ แพน1',11265.00,11265.00,50,0,'877972751'),(1234567890,'นาย ทดสอบ ระบบ',10000.00,10000.00,10,0,'0891234567'),(1234567891,'นาย ทดสอบ ระบบ',10000.00,8000.00,10,0,'0891234567'),(1234567892,'นาย ทดสอบ ระบบ',10000.80,8000.13,10,0,'0891234567'),(1234567893,'นาย ทดสอบ ระบบ',1500.00,-1500.00,10,0,'0891234567'),(1234567894,'นาย ทดสอบ ระบบ',199600.00,199600.00,10,0,'0891234567'),(1234567895,'นาย ทดสอบ ระบบ',200000.00,200000.00,10,0,'0891234567'),(1234567896,'นาย ทดสอบ ระบบ',-500000.00,-500000.00,10,0,'0891234567'),(1234567897,'นาย ทดสอบ ระบบ',200000.00,200000.00,10,0,'0891234567'),(1234567898,'นาย ทดสอบ ระบบ',199650.00,199650.00,10,0,'891234567'),(1234567899,'นาย ทดสอบ ระบบ',199350.00,199350.00,10,0,'891234567'),(2000000001,'นาย สมชาย เข็มกลัด',400.00,400.00,10,0,'0898765432'),(2000000002,'นาย สมชาย เข็มกลัด',0.00,0.00,10,0,'0898765432'),(2000000003,'นาย สมชาย เข็มกลัด',700000.00,700000.00,10,0,'0898765432'),(2000000004,'นาย สมชาย เข็มกลัด',0.00,0.00,10,0,'0898765432'),(2000000005,'YJ Test acc1',156899.00,156899.00,10,0,'891234567'),(2000000006,'YJ Test acc2',53101.00,53101.00,10,0,'892468901');
/*!40000 ALTER TABLE `ACN` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ZUTBLBK`
--

DROP TABLE IF EXISTS `ZUTBLBK`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ZUTBLBK` (
  `BKCD` int(11) NOT NULL DEFAULT '0',
  `NAME` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `ENAME` varchar(60) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`BKCD`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ZUTBLBK`
--

LOCK TABLES `ZUTBLBK` WRITE;
/*!40000 ALTER TABLE `ZUTBLBK` DISABLE KEYS */;
INSERT INTO `ZUTBLBK` VALUES (0,'ธนาคารสมเกียรติ จำกัด','Somkiat Bank'),(2,'ธนาคารกรุงเทพ จำกัด','Bangkok Bank'),(4,'ธนาคารกสิกรไทย จำกัด','Kasikorn Bank'),(6,'ธนาคารกรุงไทย จำกัด','Krungthai Bank'),(11,'ธนาคารทหารไทย จำกัด','TMB Bank');
/*!40000 ALTER TABLE `ZUTBLBK` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-12-22 18:25:56
