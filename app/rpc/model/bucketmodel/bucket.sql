CREATE TABLE `bucket`
(
    `id`           bigint(0)                                                     NOT NULL AUTO_INCREMENT COMMENT 'id',
    `username`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名称',
    `bucket_name`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '桶名称',
    `capacity_cur` float(30, 5) UNSIGNED                                         NOT NULL DEFAULT 0.00000 COMMENT '当前桶存储数据量大小（以kb为单位）',
    `object_num`   int unsigned                                                  NOT NULL COMMENT '桶中对象数量',
    `permission`   tinyint unsigned                                              NOT NULL COMMENT '桶权限（0 公共 | 1 私有 ）',
    `create_time`  timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `bucketname_unique` (`bucket_name`) USING BTREE COMMENT '桶名称唯一'
) ENGINE = InnoDB
  AUTO_INCREMENT = 3
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci
  ROW_FORMAT = Dynamic;

