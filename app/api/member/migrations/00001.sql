package migrations

-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- CREATE TABLE "live_live" ------------------------------------
CREATE TABLE `live_live` (
`id` Int( 12 ) AUTO_INCREMENT NOT NULL,
`publish` VarChar( 255 ) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
`live` TinyInt( 4 ) NOT NULL,
`channelid` VarChar( 12 ) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
`actid` Int( 12 ) NOT NULL,
PRIMARY KEY ( `id` ),
CONSTRAINT `unique_actid` UNIQUE( `actid` ),
CONSTRAINT `unique_ChannelID` UNIQUE( `channelid` ) )
CHARACTER SET = utf8mb4
COLLATE = utf8mb4_general_ci
ENGINE = InnoDB
AUTO_INCREMENT = 1;
-- -------------------------------------------------------------

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS `live_live`;

