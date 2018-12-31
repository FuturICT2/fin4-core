-- +migrate Up

CREATE TABLE `action_proposal` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `actionId` INT(11) NULL,
  `userId` INT(11) NULL,
  `proposalTxt` VARCHAR(10000) NULL,
  `isApproved` TINYINT(1) NULL,
  `createdAt` DATETIME NULL,
  PRIMARY KEY (`id`));


  CREATE TABLE `claim` (
    `id` INT NOT NULL,
    `tokenId` INT NOT NULL,
    `userId` INT NOT NULL,
    `isApproved` TINYINT(1) NOT NULL,
    `text` VARCHAR(10000) NULL,
    `logoFile` VARCHAR(512) NULL,
    PRIMARY KEY (`id`));

    ALTER TABLE `claim` 
    CHANGE COLUMN `id` `id` INT(11) NOT NULL AUTO_INCREMENT ;
