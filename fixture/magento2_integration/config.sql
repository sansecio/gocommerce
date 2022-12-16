-- MariaDB dump 10.19  Distrib 10.4.25-MariaDB, for debian-linux-gnu (aarch64)
--
-- Host: localhost    Database: magento
-- ------------------------------------------------------
-- Server version	10.4.25-MariaDB-1:10.4.25+maria~focal

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
  `config_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Config ID',
  `scope` varchar(8) NOT NULL DEFAULT 'default' COMMENT 'Config Scope',
  `scope_id` int(11) NOT NULL DEFAULT 0 COMMENT 'Config Scope ID',
  `path` varchar(255) NOT NULL DEFAULT 'general' COMMENT 'Config Path',
  `value` text DEFAULT NULL COMMENT 'Config Value',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT 'Updated At',
  PRIMARY KEY (`config_id`),
  UNIQUE KEY `CORE_CONFIG_DATA_SCOPE_SCOPE_ID_PATH` (`scope`,`scope_id`,`path`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8 COMMENT='Config Data';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `core_config_data`
--

LOCK TABLES `core_config_data` WRITE;
/*!40000 ALTER TABLE `core_config_data` DISABLE KEYS */;
INSERT INTO `core_config_data` VALUES (1,'default',0,'catalog/search/engine','elasticsearch7','2022-11-16 13:22:39'),(2,'default',0,'catalog/search/elasticsearch7_server_hostname','elasticsearch','2022-11-16 13:22:39'),(3,'default',0,'catalog/search/elasticsearch7_server_port','9200','2022-11-16 13:22:39'),(4,'default',0,'catalog/search/elasticsearch7_enable_auth','0','2022-11-16 13:22:39'),(5,'default',0,'catalog/search/elasticsearch7_index_prefix','magento2','2022-11-16 13:22:39'),(6,'default',0,'catalog/search/elasticsearch7_server_timeout','15','2022-11-16 13:22:39'),(7,'default',0,'general/region/display_all','1','2022-11-16 13:22:39'),(8,'default',0,'general/region/state_required','AL,AR,AU,BG,BO,BR,BY,CA,CH,CL,CN,CO,DK,EC,EE,ES,GR,GY,HR,IN,IS,IT,LT,LV,MX,PE,PL,PT,PY,RO,SE,SR,US,UY,VE','2022-11-16 13:22:41'),(9,'default',0,'catalog/category/root_id','2','2022-11-16 13:22:42'),(10,'default',0,'analytics/subscription/enabled','1','2022-11-16 13:22:43'),(11,'default',0,'crontab/default/jobs/analytics_subscribe/schedule/cron_expr','0 * * * *','2022-11-16 13:22:43'),(12,'default',0,'crontab/default/jobs/analytics_collect_data/schedule/cron_expr','00 02 * * *','2022-11-16 13:22:43'),(13,'default',0,'msp_securitysuite_recaptcha/frontend/enabled','0','2022-11-16 13:22:43'),(14,'default',0,'msp_securitysuite_recaptcha/backend/enabled','0','2022-11-16 13:22:43'),(15,'default',0,'twofactorauth/duo/application_key','sqxTJMd2R2NKb711vmiIB3Dq6RRfW4jgfREsxYnJ85CL6Eyiw1T85GuTLVBffcAn','2022-11-16 13:22:43'),(16,'default',0,'design/theme/theme_id','2','2022-11-16 13:24:04'),(17,'default',0,'design/head/includes','<link  rel=\"stylesheet\" type=\"text/css\"  media=\"all\" href=\"{{MEDIA_URL}}styles.css\" />','2022-11-16 13:24:04'),(18,'default',0,'carriers/tablerate/active','1','2022-11-16 13:25:38'),(19,'default',0,'carriers/tablerate/condition_name','package_value_with_discount','2022-11-16 13:25:38'),(20,'default',0,'sales/msrp/enabled','1','2022-11-16 13:25:40'),(21,'default',0,'admin/usage/enabled','1','2022-12-13 09:54:36'),(22,'default',0,'web/unsecure/base_url','https://sansec.io/','2022-12-13 09:57:48'),(23,'stores',2,'web/unsecure/base_url','https://api.sansec.io/','2022-12-13 09:58:26');
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

-- Dump completed on 2022-12-13 10:08:47
