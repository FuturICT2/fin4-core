-- +migrate Up

CREATE TABLE `action` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `description` varchar(10000) NOT NULL,
  `status` tinyint(3) NOT NULL,
  `creatorId` int(11) NOT NULL,
  `doerId` int(11) DEFAULT NULL,
  `startsAt` datetime NOT NULL,
  `endsAt` datetime NOT NULL,
  `logoFile` varchar(512) DEFAULT NULL,
  `createdAt` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

CREATE TABLE `action_reward` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `actionId` int(11) NOT NULL,
  `tokenId` int(11) NOT NULL,
  `userId` int(11) NOT NULL,
  `amount` decimal(30,8) DEFAULT '0.00000000',
  `status` smallint(2) DEFAULT '0',
  PRIMARY KEY (`id`,`actionId`,`tokenId`,`userId`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

ALTER TABLE `user_holding`
ADD COLUMN `reserved` DECIMAL(30,8) NULL DEFAULT '0.00000000' AFTER `balance`;
