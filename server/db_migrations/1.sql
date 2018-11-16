-- +migrate Up
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `ethereumAddress` varchar(512) NOT NULL,
  `createdAt` datetime NOT NULL,
  `updatedAt` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8;

INSERT INTO `user` SET id=1, name='Demo Community', ethereumAddress='0xeed1a5d1d51893bb6fa057a1a27262c1ebabb57b', createdAt=NOW(), updatedAt=NOW();

DROP TABLE IF EXISTS `token`;
CREATE TABLE `token` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `creatorId` int(10) NOT NULL,
  `name` varchar(255) NOT NULL,
  `symbol` varchar(45) NOT NULL,
  `totalSupply` int(10) unsigned NOT NULL,
  `purpose` varchar(255) DEFAULT NULL,
  `blockchainAddress` varchar(512) NOT NULL,
  `txAddress` varchar(512) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0',
  `logo` varchar(512) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  UNIQUE KEY `symbol_UNIQUE` (`symbol`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

INSERT INTO `token` SET
  id=1,
  creatorId=1,
  name='Filsat',
  symbol='Fils',
  totalSupply=0,
  purpose='demo token',
  blockchainAddress='0x6Cd99CD0Ca8fDBBBa2c5d6D1821eF932f9fB3170',
  txAddress='0x85295762a185b80343d5c03b08dcaf408b54ab8412768aaba92cac4c1a0c2e20',
  logo='https://trello-attachments.s3.amazonaws.com/5b39f1d06f761ae7c1c7d22c/5b9f76d5afa89e794357c6d9/df4703eeb10e3fbcf1bd63288281cf9d/coin2.png';

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

INSERT INTO `user_holding` SET userId=1, tokenId=1, balance=0;

DROP TABLE IF EXISTS `token_like`;
CREATE TABLE `token_like` (
  `userId` int(11) NOT NULL,
  `tokenId` int(11) NOT NULL,
  PRIMARY KEY (`userId`,`tokenId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
