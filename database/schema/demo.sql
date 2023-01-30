CREATE TABLE `demo`
(
    `id`         varchar(32)   NOT NULL DEFAULT '',
    `created_at` datetime(3)   NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` datetime(3)   NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` varchar(32)   NOT NULL DEFAULT '',
    `email`      varchar(255)  NOT NULL DEFAULT '',
    `name`       varchar(255)  NOT NULL DEFAULT '',
    `avatar`     varchar(1024) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_email` (`email`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='demo';