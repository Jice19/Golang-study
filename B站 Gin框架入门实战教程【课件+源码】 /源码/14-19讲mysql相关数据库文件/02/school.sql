

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for class
-- ----------------------------
DROP TABLE IF EXISTS `class`;
CREATE TABLE `class`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `score` tinyint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of class
-- ----------------------------
INSERT INTO `class` VALUES (1, '张三', 'itying@qq.com', 89);
INSERT INTO `class` VALUES (2, '李四', '441@qq.com', 87);
INSERT INTO `class` VALUES (3, '王五', 'nodejs@qq.com', 98);
INSERT INTO `class` VALUES (4, '赵四', 'java@qq.com', 59);
INSERT INTO `class` VALUES (5, '小马', '888@qq.com', 66);
INSERT INTO `class` VALUES (6, '小李', '999@hotmail.com', 88);
INSERT INTO `class` VALUES (7, '小张', 'golang@163.com', 55);

SET FOREIGN_KEY_CHECKS = 1;
