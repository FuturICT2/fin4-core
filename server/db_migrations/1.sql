-- +migrate Up

SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `user_balance`;
CREATE TABLE `user_balance` (
  `userId` int(10) unsigned NOT NULL,
  `tokenId` int(10) unsigned NOT NULL,
  `balance` decimal(30,8) NOT NULL DEFAULT '0.00000000',
  `reserved` decimal(30,8) NOT NULL DEFAULT '0.00000000',
  PRIMARY KEY (`userId`,`tokenId`),
  CONSTRAINT `userID_FKK` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
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


DROP TABLE IF EXISTS `asset_block`;
CREATE TABLE `asset_block` (
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

DROP TABLE IF EXISTS `asset_block_favorites`;
CREATE TABLE `asset_block_favorites` (
  `userId` int(10) unsigned NOT NULL,
  `blockId` int(10) unsigned NOT NULL,
  `createdAt` datetime NOT NULL,
  UNIQUE KEY `pk_userid_assetid` (`assetId`,`userId`),
  KEY `userid_idx` (`userId`),
  CONSTRAINT `asset_favorites_fk_assetid` FOREIGN KEY (`assetId`) REFERENCES `asset` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `asset_favorites_fk_userid` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `asset_block_image`;
CREATE TABLE `asset_block_image` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `blockId` int(10) unsigned NOT NULL,
  `filepath` VARCHAR(512) NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS=1;
