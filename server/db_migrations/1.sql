-- +migrate Up

SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `gorp_migrations`;
CREATE TABLE `gorp_migrations` (
  `id` varchar(255) NOT NULL,
  `applied_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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
) ENGINE=InnoDB AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `user_balance`;
CREATE TABLE `user_balance` (
  `userId` int(10) unsigned NOT NULL,
  `tokenId` int(10) unsigned NOT NULL,
  `balance` decimal(30,8) NOT NULL DEFAULT '0.00000000',
  `reserved` decimal(30,8) NOT NULL DEFAULT '0.00000000',
  PRIMARY KEY (`userId`,`tokenId`),
  CONSTRAINT `userID_FKK` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `user_change_email_confirm`;
CREATE TABLE `user_change_email_confirm` (
  `userId` int(10) unsigned NOT NULL,
  `email` varchar(255) NOT NULL,
  `token` varchar(255) NOT NULL,
  `createdAt` datetime NOT NULL,
  PRIMARY KEY (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `user_password_reset`;
CREATE TABLE `user_password_reset` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `userId` int(10) unsigned NOT NULL,
  `token` varchar(255) DEFAULT NULL,
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `userId` (`userId`),
  CONSTRAINT `userId` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `asset`;
CREATE TABLE `asset` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `symbol` varchar(45) NOT NULL,
  `description` text NOT NULL,
  `supply` int(10) unsigned NOT NULL DEFAULT '1000',
  `creatorId` int(10) unsigned NOT NULL,
  `minersCounter` int(10) unsigned NOT NULL DEFAULT '0',
  `favoritesCounter` int(10) unsigned NOT NULL DEFAULT '0',
  `ethereumAddress` varchar(512) NOT NULL,
  `ethereumTransactionAddress` varchar(512) NOT NULL,
  `createdAt` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  UNIQUE KEY `symbol_UNIQUE` (`symbol`)
) ENGINE=InnoDB AUTO_INCREMENT=250 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `asset_favorites`;
CREATE TABLE `asset_favorites` (
  `userId` int(10) unsigned NOT NULL,
  `assetId` int(10) unsigned NOT NULL,
  `addedAt` datetime NOT NULL,
  UNIQUE KEY `pk_userid_assetid` (`assetId`,`userId`),
  KEY `userid_idx` (`userId`),
  CONSTRAINT `asset_favorites_fk_assetid` FOREIGN KEY (`assetId`) REFERENCES `asset` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `asset_favorites_fk_userid` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `asset_claim`;
CREATE TABLE `asset_claim` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `assetId` int(10) unsigned NOT NULL,
  `userId` int(10) unsigned NOT NULL,
  `text` varchar(10000),
  `videoID` varchar(256) NOT NULL,
  `status` tinyint(2) unsigned NOT NULL,
  `favoritesCounter` int(10) unsigned NOT NULL DEFAULT '0',
  `ethereumTransactionAddress` varchar(512) NOT NULL DEFAULT '',
  `createdAt` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `asset_fkk` (`assetId`),
  KEY `user_fkk` (`userId`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `asset_claim_favorites`;
CREATE TABLE `asset_claim_favorites` (
  `userId` int(10) unsigned NOT NULL,
  `claimId` int(10) unsigned NOT NULL,
  `createdAt` datetime NOT NULL,
  UNIQUE KEY `pk_userid_assetid` (`assetId`,`userId`),
  KEY `userid_idx` (`userId`),
  CONSTRAINT `asset_favorites_fk_assetid` FOREIGN KEY (`assetId`) REFERENCES `asset` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `asset_favorites_fk_userid` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `asset_claim_image`;
CREATE TABLE `asset_claim_image` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `claimId` int(10) unsigned NOT NULL,
  `filepath` VARCHAR(512) NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS=1;
