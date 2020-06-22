-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- CREATE TABLE "member" ------------------------------------
CREATE TABLE `member` (
     `id` INT( 12 ) AUTO_INCREMENT NOT NULL,
     `phone` VARCHAR( 30 )  NOT NULL,
     `name` VARCHAR( 30 ) NOT NULL,
     PRIMARY KEY ( `id` ),
     UNIQUE KEY `uk_phone` ( `phone` ))
ENGINE = InnoDB
AUTO_INCREMENT = 1;
-- -------------------------------------------------------------

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS `member`;
