-- +migrate Up

SET FOREIGN_KEY_CHECKS=0;

ALTER table asset ADD column oracleType tinyint(2) unsigned NOT NULL DEFAULT 0 AFTER creatorId;
ALTER table asset ADD column lastOraclePing datetime NOT NULL DEFAULT Now() AFTER oracleType;
ALTER table asset ADD column accessToken varchar(255) NOT NULL AFTER oracleType;

SET FOREIGN_KEY_CHECKS=1;
