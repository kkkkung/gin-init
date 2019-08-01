-- 用户表
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`           int(11)        NOT NULL AUTO_INCREMENT COMMENT '主键',

    `unionid`      char(80) COMMENT 'UNIONID',
    `openid`       char(80) COMMENT 'OPENID',
    `session_key`  char(80) COMMENT 'OPENID',
    `nickname`     char(80)       NOT NULL DEFAULT '' COMMENT '用户名',
    `avatar`       varchar(255)   NOT NULL DEFAULT '' COMMENT '头像',
    `sex`          tinyint(1)     NOT NULL DEFAULT 0 COMMENT '性别,0-未知 1-男 2-女',
    `mobile`       char(18)       NOT NULL DEFAULT '' COMMENT '手机',
    `introduction` char(80)       NOT NULL DEFAULT '' COMMENT '简介',
    `balance`      decimal(12, 2) NOT NULL DEFAULT 0.00 COMMENT '余额',
    `coin`         decimal(12, 2) NOT NULL DEFAULT 0.00 COMMENT '平台币',
    `password`     varchar(255)   NOT NULL DEFAULT '' COMMENT '密码',
    `ssk`          varchar(255)   NOT NULL DEFAULT '' COMMENT 'sessionKey',

    `created_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '写入时间',
    `updated_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '更新时间',
    `deleted_at`   timestamp      NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE (`mobile`),
    KEY (`openid`),
    KEY (`unionid`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='用户表';

-- 管理员表
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins`
(
    `id`         int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',

    `nickname`   char(80)     NOT NULL DEFAULT '' COMMENT '用户名',
    `avatar`     varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
    `username`   char(18)     NOT NULL DEFAULT '' COMMENT 'Username',
    `password`   varchar(255) NOT NULL DEFAULT '' COMMENT '密码',

    `created_at` timestamp    NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '写入时间',
    `updated_at` timestamp    NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '更新时间',
    `deleted_at` timestamp    NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='管理员表';
-- 插入管理员
INSERT INTO `admins`(nickname, avatar, username, password)
VALUES ('Admin', 'xxxxxx', 'admin', '123456')