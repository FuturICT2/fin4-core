-- +migrate Up

CREATE TABLE `action_proposal` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `actionId` INT(11) NULL,
  `userId` INT(11) NULL,
  `proposalTxt` VARCHAR(10000) NULL,
  `isApproved` TINYINT(1) NULL,
  `createdAt` DATETIME NULL,
  PRIMARY KEY (`id`));
