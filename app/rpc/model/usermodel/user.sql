CREATE TABLE `user` (
    `id`          bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `username`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名称',
    `password`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户密码',
    `create_time` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP (0) COMMENT '更新时间',
    `deleted`     tinyint unsigned NOT NULL COMMENT '0（未删除） | 1（删除） | 逻辑删除',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `username_unique`(`username`) USING BTREE COMMENT '用户名唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;