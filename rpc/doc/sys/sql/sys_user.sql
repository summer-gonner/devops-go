-- goctl model mysql ddl -src doc/sql/sys/sys_user.sql -dir ./model/sysmodel

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
    `name`        varchar(128) NOT NULL DEFAULT '' COMMENT '账号',
    `nick_name`   varchar(128) NOT NULL DEFAULT '' COMMENT '名称',
    `avatar`      varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
    `password`    varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
    `salt`        varchar(40)  NOT NULL DEFAULT '' COMMENT '加密盐',
    `email`       varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱',
    `mobile`      varchar(32)  NOT NULL DEFAULT '' COMMENT '手机号',
    `status`      tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态  -1：禁用   1：正常',
    `create_by`   varchar(128) NOT NULL DEFAULT '' COMMENT '创建人',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_by`   varchar(128) NOT NULL DEFAULT '' COMMENT '更新人',
    `update_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `del_flag`    tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除  1：已删除  0：正常',
    PRIMARY KEY (`id`),
    KEY           `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户管理';
-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', 'admin', '', '$2a$10$hDlSis2/3IPGNYQhFlFfK.Wmi7iH9/jr6wcN.5c.rh7fc/uUnCo4S', '', 'admin@dsms.com', '13612345678', 1, 'admin', '2018-08-14 11:11:11', '', '2023-01-04 10:17:30', 0);
