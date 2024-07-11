CREATE DATABASE  IF NOT EXISTS `kmkglass` /*!40100 DEFAULT CHARACTER SET utf8mb3 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `kmkglass`;
-- MySQL dump 10.13  Distrib 8.0.36, for Win64 (x86_64)
--
-- Host: localhost    Database: kmkglass
-- ------------------------------------------------------
-- Server version	8.0.37

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `brands`
--

DROP TABLE IF EXISTS `brands`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `brands` (
  `idbrands` int NOT NULL AUTO_INCREMENT,
  `name` varchar(250) NOT NULL,
  PRIMARY KEY (`idbrands`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `glass_options`
--

DROP TABLE IF EXISTS `glass_options`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `glass_options` (
  `idglass_options` int NOT NULL AUTO_INCREMENT,
  `name` varchar(250) NOT NULL,
  `glass_type_name` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`idglass_options`),
  KEY `name` (`name`),
  KEY `glass_types_name_idx` (`glass_type_name`),
  CONSTRAINT `glass_types_name` FOREIGN KEY (`glass_type_name`) REFERENCES `glass_types` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `glass_types`
--

DROP TABLE IF EXISTS `glass_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `glass_types` (
  `idglass_types` int NOT NULL AUTO_INCREMENT,
  `name` varchar(250) NOT NULL,
  PRIMARY KEY (`idglass_types`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `models`
--

DROP TABLE IF EXISTS `models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `models` (
  `idmodels` int NOT NULL AUTO_INCREMENT,
  `name` varchar(250) NOT NULL,
  `brand_name` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`idmodels`),
  KEY `name` (`name`),
  KEY `brands_name_idx` (`brand_name`),
  CONSTRAINT `brands_name` FOREIGN KEY (`brand_name`) REFERENCES `brands` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `idproducts` int NOT NULL AUTO_INCREMENT,
  `price` int NOT NULL,
  `name` varchar(200) NOT NULL,
  `article` varchar(30) NOT NULL,
  `length` int NOT NULL,
  `photo` varchar(500) DEFAULT NULL,
  `width` int NOT NULL,
  `amount` int DEFAULT NULL,
  `brands_name` varchar(250) NOT NULL,
  `models_name` varchar(250) NOT NULL,
  `year_model_name` varchar(250) NOT NULL,
  `glass_types_name` varchar(250) NOT NULL,
  `glass_options_name` varchar(250) NOT NULL,
  PRIMARY KEY (`idproducts`),
  KEY `index2` (`name`),
  KEY `brands_id_idx` (`brands_name`),
  KEY `models_id_idx` (`models_name`),
  KEY `year_model_id_idx` (`year_model_name`),
  KEY `glass_types_id_idx` (`glass_types_name`),
  KEY `glass_options_id_idx` (`glass_options_name`),
  KEY `name` (`name`),
  KEY `article` (`article`),
  CONSTRAINT `product_brands_name` FOREIGN KEY (`brands_name`) REFERENCES `brands` (`name`),
  CONSTRAINT `product_glass_options_name` FOREIGN KEY (`glass_options_name`) REFERENCES `glass_options` (`name`),
  CONSTRAINT `product_glass_types_name` FOREIGN KEY (`glass_types_name`) REFERENCES `glass_types` (`name`),
  CONSTRAINT `product_models_name` FOREIGN KEY (`models_name`) REFERENCES `models` (`name`),
  CONSTRAINT `product_year_model_name` FOREIGN KEY (`year_model_name`) REFERENCES `year_model` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 TRIGGER `products_BEFORE_INSERT` BEFORE INSERT ON `products` FOR EACH ROW BEGIN
	IF NOT EXISTS (
		SELECT 1
		FROM brands
		WHERE name = NEW.brands_name)
	THEN
		INSERT INTO brands (name)
		VALUES (NEW.brands_name);
    END IF;
    
        IF NOT EXISTS (
		SELECT 1
		FROM models
		WHERE name = NEW.models_name AND brand_name = NEW.brands_name)
	THEN
		INSERT INTO models (name, brand_name)
		VALUES (NEW.models_name, NEW.brands_name);
	END IF;
    
	IF NOT EXISTS (
		SELECT 1
		FROM year_model
		WHERE name = NEW.year_model_name AND model_name = NEW.models_name)
	THEN
		INSERT INTO year_model (name, model_name)
		VALUES (NEW.year_model_name, NEW.models_name);
    END IF;
    
	IF NOT EXISTS (
		SELECT 1
		FROM glass_types
		WHERE name = NEW.glass_types_name)
	THEN
		INSERT INTO glass_types (name)
		VALUES (NEW.glass_types_name);
    END IF;
    
    IF NOT EXISTS (
		SELECT 1
		FROM glass_options
		WHERE name = NEW.glass_options_name AND glass_type_name = NEW.glass_types_name)
	THEN
		INSERT INTO glass_options (name, glass_type_name)
		VALUES (NEW.glass_options_name, NEW.glass_types_name);
	END IF;
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 TRIGGER `products_BEFORE_UPDATE` BEFORE UPDATE ON `products` FOR EACH ROW BEGIN
	IF NOT EXISTS (
		SELECT 1
		FROM brands
		WHERE name = NEW.brands_name)
	THEN
		INSERT INTO brands (name)
		VALUES (NEW.brands_name);
    END IF;
    
        IF NOT EXISTS (
		SELECT 1
		FROM models
		WHERE name = NEW.models_name AND brand_name = NEW.brands_name)
	THEN
		INSERT INTO models (name, brand_name)
		VALUES (NEW.models_name, NEW.brands_name);
	END IF;
    
	IF NOT EXISTS (
		SELECT 1
		FROM year_model
		WHERE name = NEW.year_model_name AND model_name = NEW.models_name)
	THEN
		INSERT INTO year_model (name, model_name)
		VALUES (NEW.year_model_name, NEW.models_name);
    END IF;
    
	IF NOT EXISTS (
		SELECT 1
		FROM glass_types
		WHERE name = NEW.glass_types_name)
	THEN
		INSERT INTO glass_types (name)
		VALUES (NEW.glass_types_name);
    END IF;
    
    IF NOT EXISTS (
		SELECT 1
		FROM glass_options
		WHERE name = NEW.glass_options_name AND glass_type_name = NEW.glass_types_name)
	THEN
		INSERT INTO glass_options (name, glass_type_name)
		VALUES (NEW.glass_options_name, NEW.glass_types_name);
	END IF;
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `year_model`
--

DROP TABLE IF EXISTS `year_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `year_model` (
  `idyear_model` int NOT NULL AUTO_INCREMENT,
  `name` varchar(250) NOT NULL,
  `model_name` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`idyear_model`),
  KEY `name` (`name`),
  KEY `models_name_idx` (`model_name`),
  CONSTRAINT `models_name` FOREIGN KEY (`model_name`) REFERENCES `models` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping events for database 'kmkglass'
--

--
-- Dumping routines for database 'kmkglass'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-07-10  1:29:26
