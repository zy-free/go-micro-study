-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- CREATE TABLE "member" ------------------------------------
CREATE TABLE `member`
(
    `id`         INT AUTO_INCREMENT NOT NULL,
    `phone`      VARCHAR(30)        NOT NULL,
    `name`       VARCHAR(30)        NOT NULL,
    `age`        TINYINT            NOT NULL DEFAULT 0,
    `address`    VARCHAR(255)       NOT NULL DEFAULT '',
    `order_num`  INT                NOT NULL DEFAULT 0,
    `created_at` DATETIME           NOT NULL,
    `updated_at` DATETIME           NOT NULL,
    `deleted`    TINYINT            NOT NULL DEFAULT 0,
    `deleted_at` DATETIME,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_phone` (`phone`),
    INDEX `idx_deleted` (deleted)
)
    ENGINE = InnoDB
    AUTO_INCREMENT = 1;
-- -------------------------------------------------------------

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS `member`;
