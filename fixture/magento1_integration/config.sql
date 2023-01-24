-- MySQL dump 10.19  Distrib 10.3.35-MariaDB, for debian-linux-gnu (aarch64)
--
-- Host: localhost    Database: magento
-- ------------------------------------------------------
-- Server version	10.3.35-MariaDB-1:10.3.35+maria~focal

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `core_config_data`
--

DROP TABLE IF EXISTS `core_config_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `core_config_data` (
  `config_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Config Id',
  `scope` varchar(8) NOT NULL DEFAULT 'default' COMMENT 'Config Scope',
  `scope_id` int(11) NOT NULL DEFAULT 0 COMMENT 'Config Scope Id',
  `path` varchar(255) NOT NULL DEFAULT 'general' COMMENT 'Config Path',
  `value` text DEFAULT NULL COMMENT 'Config Value',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`config_id`),
  UNIQUE KEY `UNQ_CORE_CONFIG_DATA_SCOPE_SCOPE_ID_PATH` (`scope`,`scope_id`,`path`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8 COMMENT='Config Data';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `core_config_data`
--

LOCK TABLES `core_config_data` WRITE;
/*!40000 ALTER TABLE `core_config_data` DISABLE KEYS */;
INSERT INTO `core_config_data` VALUES (1,'default',0,'advanced/modules_disable_output/Mage_Backup','1','2023-01-24 10:13:32'),(2,'default',0,'admin/security/validate_formkey_checkout','1','2023-01-24 10:13:34'),(3,'default',0,'general/region/display_all','1','2023-01-24 10:13:35'),(4,'default',0,'general/region/state_required','AT,CA,CH,DE,EE,ES,FI,FR,LT,LV,RO,US','2023-01-24 10:13:35'),(5,'default',0,'catalog/category/root_id','2','2023-01-24 10:13:36'),(6,'default',0,'payment/paypal_express/skip_order_review_step','1','2023-01-24 10:13:36'),(7,'default',0,'payment/payflow_link/mobile_optimized','1','2023-01-24 10:13:36'),(8,'default',0,'payment/payflow_advanced/mobile_optimized','1','2023-01-24 10:13:36'),(9,'default',0,'payment/hosted_pro/mobile_optimized','1','2023-01-24 10:13:36'),(10,'default',0,'admin/dashboard/enable_charts','1','2023-01-24 10:13:36'),(11,'default',0,'web/unsecure/base_url','https://app.magento1.test/','2023-01-24 10:13:36'),(12,'default',0,'web/secure/base_url','https://app.magento1.test/','2023-01-24 10:13:36'),(13,'default',0,'general/locale/code','en_US','2023-01-24 10:13:36'),(14,'default',0,'general/locale/timezone','America/Los_Angeles','2023-01-24 10:13:36'),(15,'default',0,'currency/options/base','USD','2023-01-24 10:13:36'),(16,'default',0,'currency/options/default','USD','2023-01-24 10:13:36'),(17,'default',0,'currency/options/allow','USD','2023-01-24 10:13:36'),(18,'stores',2,'web/secure/base_url','https://second.magento1.test/','2023-01-24 10:14:59');
/*!40000 ALTER TABLE `core_config_data` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-01-24 10:16:26
