CREATE TABLE `object_hash`  (
  `id` bigint unsigned NOT NULL COMMENT 'id',
  `hashcode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '对象hashcode',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `hashcoed_unique`(`hashcode`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;