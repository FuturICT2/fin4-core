-- +migrate Up

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `salt` varchar(32) NOT NULL,
  `createdAt` datetime NOT NULL,
  `updatedAt` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `token`;
CREATE TABLE `token` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `creatorId` INT(10) NOT NULL,
  `name` varchar(255) NOT NULL,
  `symbol` varchar(45) NOT NULL,
  `totalSupply` int(10) unsigned NOT NULL,
  `purpose` varchar(255),
  `blockchainAddress` varchar(512) NOT NULL,
  `txAddress` VARCHAR(512) NOT NULL,
  `status` INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  UNIQUE KEY `symbol_UNIQUE` (`symbol`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `user_holding`;
CREATE TABLE `user_holding` (
  `userId` int(10) unsigned NOT NULL,
  `tokenId` int(10) unsigned NOT NULL,
  `balance` decimal(30,8) NOT NULL DEFAULT '0.00000000',
  PRIMARY KEY (`userId`,`tokenId`),
  KEY `tokenId` (`tokenId`),
  CONSTRAINT `tokenID_FKK` FOREIGN KEY (`tokenId`) REFERENCES `token` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `userID_FKK` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
