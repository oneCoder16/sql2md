/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50553
 Source Host           : localhost:3306
 Source Schema         : dbitop_admin

 Target Server Type    : MySQL
 Target Server Version : 50553
 File Encoding         : 65001

 Date: 15/08/2019 13:11:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_channel
-- ----------------------------
DROP TABLE IF EXISTS `tb_channel`;
CREATE TABLE `tb_channel`  (
  `id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '渠道名',
  `desc` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '渠道描述',
  `name_en` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '渠道英文简称',
  `param_conf` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '渠道配置参数，json格式',
  `cmcc_server` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '配置中心的server信息',
  `cmcc_server_id` int(11) NOT NULL COMMENT '配置中心的tbServerInfo表中对应的id --已弃用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of tb_channel
-- ----------------------------
INSERT INTO `tb_channel` VALUES (1, '微信', '', '', '{\"app_id\": {\"name\": \"应用id\",\"desc\": \"wechart应用ID\",\"required\": 1},\"app_key\":{\"name\": \"应用密钥\",\"desc\": \"应用的密钥\",\"required\": 1},\"is_need_game_health_query\":{\"name\": \"是否进行健康检查\",\"desc\": \"是否进行健康检查\",\"required\": 0,\"type\":\"intval\",\"tip\":\"0:不检查 1:检查\"},\"need_name_auth\":{\"name\": \"是否开启实名认证\",\"desc\": \"是否开启实名认证\",\"required\": 0,\"type\":\"intval\",\"tip\":\"0:不开启 1:开启\"},\"mock_flag\":{\"name\": \"是否允许模拟登录\",\"desc\": \"是否允许模拟登录\",\"required\": 0,\"type\":\"intval\",\"tip\":\"0:不允许 1:允许\"}}', '', 0);
INSERT INTO `tb_channel` VALUES (2, '手Q', '', '', '{\"app_id\":{\"name\": \"应用id\",\"desc\": \"手Q应用ID\",\"required\": 1},\"app_key\": {\"name\": \"应用密钥\",\"desc\": \"应用的密钥\",\"required\": 1},\"is_need_game_health_query\":{\"name\": \"是否进行健康检查\",\"desc\": \"是否进行健康检查\",\"required\": 0,\"type\":\"intval\",\"tip\":\"0:不检查 1:检查\"},\"need_name_auth\":{\"name\": \"是否开启实名认证\",\"desc\": \"是否开启实名认证\",\"required\": 0,\"type\":\"intval\",\"tip\":\"0:不开启 1:开启\"},\"mock_flag\":{\"name\": \"是否允许模拟登录\",\"desc\": \"是否允许模拟登录\",\"required\": 0,\"type\":\"intval\",\"tip\":\"0:不允许 1:允许\"}}', '', 0);
INSERT INTO `tb_channel` VALUES (3, '游客', '', '', '', '', 0);
INSERT INTO `tb_channel` VALUES (4, 'Facebook', 'Facebook', 'Facebook', '{\"app_id\":{\"name\": \"应用id\",\"desc\": \"游戏在facebook注册的应用ID\",\"required\": 1},\"app_secret\": {\"name\": \"应用密钥\",\"desc\": \"游戏在facebook注册应用的密钥\",\"required\": 1},\"join_business\": {\"name\": \"是否加入facebook商务应用平台\",\"desc\": \"是否加入facebook商务应用平台\",\"required\": 0,\"type\":\"intval\",\"tip\":\"0:不加入 1:加入\"},\"display_name\": {\"name\": \"displayName\",\"desc\": \"应用分享是需要用到\",\"required\": 0,\"type\":\"intval\",\"tip\":\"0:不分享 1:分享\"},\"mock_flag\": {\"name\": \"是否允许模拟登录\",\"desc\": \"是否允许模拟登录\",\"required\": 0,\"type\":\"intval\",\"tip\":\"0:不允许 1:允许\"}}', '', 0);
INSERT INTO `tb_channel` VALUES (5, 'GameCenter', '', '', '{\"mock_flag\": {\"name\": \"是否允许模拟登录\",\"desc\": \"是否允许模拟登录\",\"required\": 0,\"type\":\"intval\",\"tip\":\"0:不允许 1:允许\"}}', '', 0);
INSERT INTO `tb_channel` VALUES (6, 'GooglePlay', '', '', '{\"app_id\":{\"name\": \"应用id\",\"desc\": \"游戏在google+平台注册的应用ID\",\"required\": 1},\"app_secret\":{\"name\":\"应用密钥\",\"desc\": \"游戏在google+平台应用的密钥\",\"required\": 1},\"redirect_url\": {\"name\": \"应用回调地址\",\"desc\": \"游戏在google+平台应用配置的回调地址\",\"required\": 0},\"mock_flag\": {\"name\": \"是否允许模拟登录\",\"desc\": \"是否允许模拟登录\",\"required\": 0,\"type\":\"intval\",\"tip\":\"0:不允许 1:允许\"}}', '', 0);
INSERT INTO `tb_channel` VALUES (9, 'Twitter', '', '', '{\"app_id\":{\"name\": \"应用id\",\"desc\": \"游戏在Twitter平台注册的应用ID\",\"required\": 1},\"app_secret\":{\"name\":\"应用密钥\",\"desc\": \"游戏在Twitter平台应用的密钥\",\"required\": 1},\"mock_flag\": {\"name\": \"是否允许模拟登录\",\"desc\": \"是否允许模拟登录\",\"required\": 0,\"type\":\"intval\",\"tip\":\"0:不允许 1:允许\"}}', '', 0);

SET FOREIGN_KEY_CHECKS = 1;
