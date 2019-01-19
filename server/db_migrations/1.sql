-- +migrate Up

SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `gorp_migrations`;
CREATE TABLE `gorp_migrations` (
  `id` varchar(255) NOT NULL,
  `applied_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `salt` varchar(32) NOT NULL,
  `ethereumAddress` VARCHAR(512) NOT NULL,
  `createdAt` datetime NOT NULL,
  `updatedAt` datetime NOT NULL,
  `agreeToTerms` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `isDeleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8;



DROP TABLE IF EXISTS `user_balance`;
CREATE TABLE `user_balance` (
  `userId` int(10) unsigned NOT NULL,
  `tokenId` int(10) unsigned NOT NULL,
  `balance` decimal(30,8) NOT NULL DEFAULT '0.00000000',
  `reserved` decimal(30,8) NOT NULL DEFAULT '0.00000000',
  PRIMARY KEY (`userId`,`tokenId`),
  CONSTRAINT `userID_FKK` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `user_change_email_confirm`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_change_email_confirm` (
  `userId` int(10) unsigned NOT NULL,
  `email` varchar(255) NOT NULL,
  `token` varchar(255) NOT NULL,
  `createdAt` datetime NOT NULL,
  PRIMARY KEY (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `user_password_reset`;

CREATE TABLE `user_password_reset` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `userId` int(10) unsigned NOT NULL,
  `token` varchar(255) DEFAULT NULL,
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `userId` (`userId`),
  CONSTRAINT `userId` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `token`;
CREATE TABLE `token` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `creatorId` int(10) NOT NULL,
  `name` varchar(255) NOT NULL,
  `symbol` varchar(45) NOT NULL,
  `hardCap` decimal(30,8) NOT NULL DEFAULT '0.00000000',
  `totalSupply` decimal(30,8) NOT NULL DEFAULT '0.00000000',
  `purpose` varchar(255) DEFAULT NULL,
  `blockchainAddress` varchar(512) NOT NULL,
  `txAddress` varchar(512) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  UNIQUE KEY `symbol_UNIQUE` (`symbol`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `token_like`;
CREATE TABLE `token_like` (
  `userId` int(10) NOT NULL,
  `tokenId` int(10) NOT NULL,
  PRIMARY KEY (`userId`,`tokenId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `claim`;
CREATE TABLE `claim` (
  `id` INT(10) NOT NULL AUTO_INCREMENT,
  `tokenId` INT NOT NULL,
  `userId` INT NOT NULL,
  `isApproved` TINYINT(1) NOT NULL,
  `text` VARCHAR(10000) NULL,
  `logoFile` VARCHAR(512) NULL,
  PRIMARY KEY (`id`));

SET FOREIGN_KEY_CHECKS=1;
