
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for nav
-- ----------------------------
DROP TABLE IF EXISTS `nav`;
CREATE TABLE `nav`  (
  `id` int(0) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of nav
-- ----------------------------
INSERT INTO `nav` VALUES (1, '首页');
INSERT INTO `nav` VALUES (4, '首页');
INSERT INTO `nav` VALUES (3, 'golang');
INSERT INTO `nav` VALUES (5, 'php');
INSERT INTO `nav` VALUES (2, 'java');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(0) NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `age` int(0) NULL DEFAULT NULL,
  `sex` int(0) NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '张三', 32, 1);
INSERT INTO `users` VALUES (3, 'zhaosi', 33, 2);
INSERT INTO `users` VALUES (4, '哈哈哈', 34, 2);
INSERT INTO `users` VALUES (5, 'lisi', 22, 1);
INSERT INTO `users` VALUES (6, 'wangwu', 33, 1);
INSERT INTO `users` VALUES (7, 'zhaosi1111', 33, 3);

SET FOREIGN_KEY_CHECKS = 1;
