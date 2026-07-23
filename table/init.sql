-- MySQL dump 10.13  Distrib 8.0.45, for Win64 (x86_64)
--
-- Host: localhost    Database: test
-- ------------------------------------------------------
-- Server version	8.0.45

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `blogs`
--

DROP TABLE IF EXISTS `blogs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `blogs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '文章的Id',
  `tags` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签',
  `created_at` int NOT NULL COMMENT '创建时间',
  `published_at` int NOT NULL COMMENT '发布时间',
  `updated_at` int NOT NULL COMMENT '更新时间',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标题',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '作者',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章内容',
  `status` tinyint NOT NULL COMMENT '文章状态',
  `summary` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章概要信息',
  `create_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_title` (`title`) COMMENT 'titile添加唯一键约束'
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blogs`
--

LOCK TABLES `blogs` WRITE;
/*!40000 ALTER TABLE `blogs` DISABLE KEYS */;
INSERT INTO `blogs` VALUES (1,'{}',1777001602,1781510845,1781511650,'Go全站开发','will','Md内容填充\nsadasdsa',1,'文章概要信息','zz'),(2,'{}',1780389636,1781510844,1780389636,'hello','me','内容',1,'概要',''),(3,'{}',1780390437,1781510845,1780390437,'eeeello','me','内容',1,'概要',''),(9,'{}',1781511220,1781511222,1781577244,'www','root','# 1\n```bash\n# save <秒数> <变化的key数量>\n# 意思是：在指定秒数内，有指定数量的key发生变化，就触发一次快照\n# 三个条件任意满足一个就备份\nsave 900 1       # 15分钟内有1个key变化就备份（冷数据）\nsave 300 10      # 5分钟内有10个key变化就备份（温数据）\nsave 60 1000     # 1分钟内有1000个key变化就备份（热数据）\n# 数据变化越频繁，备份越勤，丢失越少\n\n# 备份文件名，恢复时文件名必须与此一致\ndbfilename dump.rdb\n\n# 备份文件存放目录，也是 AOF 文件的存放目录\ndir /var/lib/redis\n\n# 备份失败时停止接受写请求\n# yes：备份失败时拒绝写入，防止数据不一致（推荐）\n# no：备份失败继续写入，数据可能丢失\nstop-writes-on-bgsave-error yes\n\n# 是否压缩 RDB 文件\n# yes：用 LZF 压缩，文件更小，但消耗少量 CPU\n# no：不压缩，文件大但节省 CPU\nrdbcompression yes\n\n# 是否对 RDB 文件做校验和检查\n# yes：保存和加载时都做 CRC64 校验，能检测文件损坏，有约10%性能损耗\n# no：不校验，速度快但无法检测损坏\nrdbchecksum yes\n```\n[百度](www.baidu.com)',1,'redis  测试文章','root'),(10,'{}',1781511684,0,1781511684,'111','root','1111',0,'111','root'),(11,'{}',1781511691,0,1781511691,'222','root','222',0,'222','root'),(12,'{}',1781511695,0,1781511695,'333','root','333',0,'333','root'),(14,'{}',1781511708,1782184462,1782184501,'5555','root','5555\n![ssss](/uploads/1782184486957200800.png)\n',1,'5555','root'),(15,'{}',1782373166,0,1782373166,'xc','root','# 1\nhello',0,'xzc','root');
/*!40000 ALTER TABLE `blogs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tokens`
--

DROP TABLE IF EXISTS `tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tokens` (
  `created_at` int NOT NULL COMMENT '创建时间',
  `updated_at` int NOT NULL COMMENT '更新时间',
  `user_id` int NOT NULL COMMENT '用户的Id',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名, 用户名不允许重复的',
  `access_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户的访问令牌',
  `access_token_expired_at` int NOT NULL COMMENT '令牌过期时间',
  `refresh_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '刷新令牌',
  `refresh_token_expired_at` int NOT NULL COMMENT '刷新令牌过期时间',
  PRIMARY KEY (`access_token`) USING BTREE,
  UNIQUE KEY `idx_token` (`access_token`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tokens`
--

LOCK TABLES `tokens` WRITE;
/*!40000 ALTER TABLE `tokens` DISABLE KEYS */;
INSERT INTO `tokens` VALUES (1776409384,1776409384,5,'www','d7gtma0n69aikl4sdkog',3600,'d7gtma0n69aikl4sdkp0',14400),(1776409389,1776409389,5,'www','d7gtmb8n69aiqp62bpe0',3600,'d7gtmb8n69aiqp62bpeg',14400),(1776413616,1776413616,5,'www','d7gunc0n69ak18362ddg',3600,'d7gunc0n69ak18362de0',14400),(1779843996,1779843996,14,'ne','d8b4770n69ajpm3vqfbg',3600,'d8b4770n69ajpm3vqfc0',14400),(1780038334,1780038334,14,'ne','d8cjlfgn69ajuq66mef0',3600,'d8cjlfgn69ajuq66mefg',14400),(1780045413,1780045413,14,'ne','d8clcp8n69ak601kjhlg',3600,'d8clcp8n69ak601kjhm0',14400),(1780105531,1780105531,14,'ne','d8d42eon69aj1r504t70',3600,'d8d42eon69aj1r504t7g',14400),(1780984993,1780984993,14,'ne','d8jqp88n69aj3756od6g',3600,'d8jqp88n69aj3756od70',14400),(1780987014,1780987014,14,'ne','d8jr91gn69ajrl60uu30',3600,'d8jr91gn69ajrl60uu3g',14400),(1780987777,1780987777,14,'ne','d8jrf08n69ajrl60uu40',3600,'d8jrf08n69ajrl60uu4g',14400),(1780989654,1780989654,14,'ne','d8jrtlgn69ajrl60uu50',3600,'d8jrtlgn69ajrl60uu5g',14400),(1780990665,1780990665,14,'ne','d8js5i8n69ajrl60uu60',3600,'d8js5i8n69ajrl60uu6g',14400),(1781508856,1781508856,18,'root','d8nqlu0n69akk72082ag',3600,'d8nqlu0n69akk72082b0',14400),(1781510792,1781510792,18,'root','d8nr520n69akk72082bg',3600,'d8nr520n69akk72082c0',14400),(1781577223,1781577223,18,'root','d8obc1on69akja3b2jgg',3600,'d8obc1on69akja3b2jh0',14400),(1781763696,1781763696,18,'root','d8poss0n69ailb63m88g',3600,'d8poss0n69ailb63m890',14400),(1781763788,1781763788,18,'root','d8potj0n69ailb63m89g',3600,'d8potj0n69ailb63m8a0',14400),(1781763820,1781763820,18,'root','d8potr0n69ailb63m8ag',3600,'d8potr0n69ailb63m8b0',14400),(1781763888,1781763888,18,'root','d8pouc0n69ailb63m8bg',3600,'d8pouc0n69ailb63m8c0',14400),(1782184377,1782184377,18,'root','d8svje8n69alk66v649g',3600,'d8svje8n69alk66v64a0',14400),(1782195911,1782195911,18,'root','d8t2dhon69alhe6blu70',3600,'d8t2dhon69alhe6blu7g',14400),(1782441296,1782441296,19,'mmm','d8uuak0n69aji24975pg',3600,'d8uuak0n69aji24975q0',14400);
/*!40000 ALTER TABLE `tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` int NOT NULL COMMENT '创建时间',
  `updated_at` int NOT NULL COMMENT '更新时间',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名, 用户名不允许重复的',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '不能保持用户的明文密码',
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户标签',
  `role` tinyint NOT NULL COMMENT '用户的角色',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_user` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (5,1776222700,1776222700,'www','$2a$10$ZvtU/Z6/U9paJiyxAg/2w.XALYKM54pnmKSUZ3rP8V8DuDcE9Spqe','{}',0),(6,1776404225,1776404225,'wo','$2a$10$lGl7ykp3ClwtfCkAZNdOau3hyeATxKPj8fjztG.Tg7eR6UXIfyKIu','{}',0),(7,1779516798,1779516798,'huang','$2a$10$1kAPGZFm1Nn3UtNDVq4ec.BcJ9d4WtXXnu2o/wOOGoOG.jvusX7UC','{}',0),(8,1779516999,1779516999,'h','$2a$10$Sfyw4Z41YKD3DX/6/y46fOnfkBb3rFrSFMTo1apA/8q.9ApEzv8km','{}',0),(9,1779518467,1779518467,'hx','$2a$10$4FSyB9HV7Z0wT1vBeJO4yu33hR.exQpayZu7RRGpmf6ud4tna8QbC','{}',0),(11,1779674233,1779674233,'visitor','$2a$10$dg51bZPtB./m1Rpm5DSSp.kAo1Na5co3fgBBRe/TFhB1/3P508Rpa','{}',0),(14,1779689133,1779689133,'ne','$2a$10$4FKgXdVbFq6iG493EXn.SeeMlZw7pqA6LsC1V5qwJwKlxbWzOfQwK','{}',0),(17,1781315175,1781315175,'fa','$2a$10$vOILclVIxqQM2Lv1PEQhQu14a4nOp9rOjc7elv4OmTrQuhKZBlhe6','{}',0),(18,1781505135,1781505135,'root','$2a$10$Kon0.btBwyHOkrQ9I/RCVe9EtWCK0NwXYQocgha.f4V3m4kDvNwRi','{}',0),(19,1782437382,1782437382,'mmm','$2a$10$WEOKXtsKJXdiEyjqP6PHpui1bDPt0uofO/.yVUvSAnph6KeVyRhpG','{}',0);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-07-20 16:34:03
