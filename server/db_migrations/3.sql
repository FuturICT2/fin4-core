-- +migrate Up

SET FOREIGN_KEY_CHECKS=0;

ALTER table asset ADD column oracleType tinyint(2) unsigned NOT NULL DEFAULT 0 AFTER creatorId;
ALTER table asset ADD column lastOraclePing datetime NOT NULL DEFAULT Now() AFTER oracleType;

SET FOREIGN_KEY_CHECKS=1;
