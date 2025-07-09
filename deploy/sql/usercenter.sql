drop database if exists `go-lottery-usercenter`;
create database `go-lottery-usercenter` character set = utf8mb4 collate = utf8mb4_general_ci;

use `go-lottery-usercenter`;
drop table if exists `user`;
create table `user` (
    `id` bigint(0) not null auto_increment,
    `create_time` datetime not null default current_timestamp,
    `update_time` datetime not null default current_timestamp on update current_timestamp(0),
    `mobile` char(11) default '' comment '手机号',
    `password` varchar(255) not null default '' comment '密码',
    `nickname` varchar(255) not null default '' comment '昵称',
    `sex` tinyint(1) not null default 0 comment '性别 0:男 1:女',
    `avatar` varchar(255) not null default '' comment '头像',
    `info` varchar(255) not null default '' comment '简介',
    `is_admin` tinyint(1) default 0 comment '是否管理员 0:否 1:是',
    `signature` varchar(200) not null default '' comment '个性签名',
    `location_name` varchar(100) not null default '' comment '地址名称',
    `longitude` double precision not null default 0 comment '经度',
    `latitude` double precision not null default 0 comment '纬度',
    `total_prize` int(0) not null default 0 comment '累计奖品',
    `fans` int(0) not null default 0 comment '粉丝数量',
    `all_lottery` int(0) not null default 0 comment '全部抽奖包含我发起的、我中奖的',
    `initiation_record` int(0) not null default 0 comment '发起抽奖记录',
    `winning_record` int(0) not null default 0 comment '中奖记录',
    primary key (`id`) using btree,
    unique index `idx_mobile`(`mobile`) using btree
) engine = InnoDB character set = utf8mb4 collate = utf8mb4_general_ci comment = '用户表' row_format = dynamic;

drop table if exists `user_auth`;
create table `user_auth` (
    `id` bigint(0) not null auto_increment,
    `create_time` datetime not null default current_timestamp,
    `update_time` datetime not null default current_timestamp on update current_timestamp,
    `user_id` bigint(0) not null default 0,
    `auth_key` varchar(64) character set utf8mb4 collate utf8mb4_general_ci not null default '' comment '平台唯一id',
    `auth_type` varchar(12) character set utf8mb4 collate utf8mb4_general_ci not null default '' comment '平台类型',
    primary key (`id`) using btree,
    unique index `idx_type_key`(`auth_type`,`auth_key`) using btree,
    unique index `idx_userId_key`(`user_id`,`auth_type`) using btree
) engine = InnoDB character set = utf8mb4 collate = utf8mb4_general_ci comment = '用户授权表' row_format = dynamic;

DROP TABLE IF EXISTS `user_sponsor`;
CREATE TABLE `user_sponsor`  (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `user_id` int(0) NOT NULL DEFAULT 0,
    `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1微信号 2公众号 3小程序 4微信群 5视频号',
    `applet_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT 'type=3时该字段才有意义，1小程序链接 2路径跳转 3二维码跳转',
    `name` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
    `desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '描述',
    `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'https://i-blog.csdnimg.cn/blog_migrate/a4fa5161369727154bc3a7d1c52bb9c0.png' COMMENT '头像',
    `is_show` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1显示 2不显示',
    `qr_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '二维码图片地址, type=1 2 3&applet_type=3 4的时候启用',
    `input_a` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'type=5 applet_type=2 or applet_type=1 输入框A',
    `input_b` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'type=5 applet_type=2 输入框B',
    `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '  抽奖赞助商表' ROW_FORMAT = Dynamic;