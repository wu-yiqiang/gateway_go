drop database gin_gateway;
create database if not exists gin_gateway;
use gin_gateway;
create table IF NOT EXISTS `users` (
  `id` int not null auto_increment,
  `uuid` varchar(40) not null,
  `username` varchar(40) not null,
  `password` varchar(400) not null,
  `nickname` varchar(40) not null,
  `email` varchar(40) not null,
  `phone` varchar(11) not null,
  `role` varchar(400),
  `avatar` varchar(400) not null,
  `created_time` bigint not null,
  `updated_time` bigint not null,
  `is_delete` bool not null default false,
  PRIMARY KEY ( `id` ),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `idx_uuid` (`uuid`),
  UNIQUE KEY `username_2` (`username`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO users ( username, password, uuid, avatar,nickname, email, phone, role, created_time, updated_time ) VALUES ( 'sutter', 'password123', '1', 'https;//','sutter', 'wu_yiqiang@outlook.com', '15770870823', 'admin, sys_admin', 123456799, 123456799);
INSERT INTO users ( username, password, uuid,avatar,nickname, email, phone, role, created_time, updated_time ) VALUES ( 'altas', 'passwd123', '2','http://','atals', 'wu_yiqiang@aliyun.com', '15770870825', 'admin', 123456789, 123456789);

create table IF NOT EXISTS `routers` (
    `id` int not null auto_increment,
    `name` varchar(40) not null,
    `parent_id` int not null,
    `icon` varchar(40) not null,
    `cn_name` varchar(40) not null,
    `create_time` bigint not null,
    `update_time` bigint not null,
    `is_delete` bool not null default false,
    PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO routers ( name, parent_id, icon, cn_name, create_time, update_time ) VALUES ( '/user/roles/setting', 0, 'role', '角色管理', 11234567, 11234567);
INSERT INTO routers ( name, parent_id, icon, cn_name, create_time, update_time ) VALUES ( '/user/user/setting', 0, 'user', '用户管理',11234568, 11234568);

create table IF NOT EXISTS `roles` (
  `id` int not null auto_increment,
  `name` varchar(40) not null,
  `cn_name` varchar(40) not null,
  `router` varchar(400) not null,
  `menu` varchar(400) not null,
  `create_time` bigint not null,
  `update_time` bigint not null,
  `is_delete` bool not null default false,
  PRIMARY KEY ( `id` )
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO roles ( name, cn_name, router, menu, create_time, update_time ) VALUES ( 'admin', '超级管理员', '0,1', '', 1127689, 1127689);
INSERT INTO roles ( name, cn_name, router, menu, create_time, update_time ) VALUES ( 'sys_admin', '系统管理员', '0,1', '', 1127688, 1127688);

create table IF NOT EXISTS `menus` (
  `id` int not null auto_increment,
  `name` varchar(40) not null,
  `permission` varchar(400) not null,
  `description` varchar(200) not null,
  `create_time` bigint not null,
  `update_time` bigint not null,
  `is_delete` bool not null default false,
  PRIMARY KEY ( `id` )
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO menus ( name, permission, description, create_time, update_time ) VALUES ( 'user:view', '查看', '查看权限', 112343,112343);
INSERT INTO menus ( name, permission, description, create_time, update_time ) VALUES ( 'user:update', '更新', '更新权限', 112353,112353);
INSERT INTO menus ( name, permission, description, create_time, update_time ) VALUES ( 'user:other', '特殊', '特殊权限', 112363,112363);


create table IF NOT EXISTS `files` (
    `id` int not null auto_increment,
    `file_name` varchar(400) not null,
    `file_hash` varchar(400) not null,
    `update_time` bigint not null,
    `is_delete` bool not null default false,
    PRIMARY KEY ( `id` )
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `user_friends`;
CREATE TABLE IF NOT EXISTS `user_friends` (
                                              `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
                                              `created_time` bigint not null,
                                              `updated_time` bigint not null,
                                              `is_delete` bool not null default false,
                                              `user_id` varchar(200) not null COMMENT '用户ID',
                                              `friend_id` varchar(200)  not null COMMENT '好友ID',
                                              PRIMARY KEY (`id`),
                                              KEY `idx_user_friends_user_id` (`user_id`),
                                              KEY `idx_user_friends_friend_id` (`friend_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '好友信息表';


DROP TABLE IF EXISTS `messages`;
CREATE TABLE IF NOT EXISTS `messages` (
                                          `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                                          `created_time` bigint not null,
                                          `updated_time` bigint not null,
                                          `is_delete` bool not null default false,
                                          `from_user_id` int DEFAULT NULL COMMENT '发送人ID',
                                          `to_user_id` int DEFAULT NULL COMMENT '发送对象ID',
                                          `content` varchar(2500) DEFAULT NULL COMMENT '消息内容',
                                          `url` varchar(350) DEFAULT NULL COMMENT '''文件或者图片地址''',
                                          `pic` text COMMENT '缩略图',
                                          `message_type` smallint DEFAULT NULL COMMENT '''消息类型：1单聊，2群聊''',
                                          `content_type` smallint DEFAULT NULL COMMENT '''消息内容类型：1文字，2语音，3视频''',
                                          PRIMARY KEY (`id`),
                                          KEY `idx_messages_deleted_at` (`is_delete`),
                                          KEY `idx_messages_from_user_id` (`from_user_id`),
                                          KEY `idx_messages_to_user_id` (`to_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT '消息表';


DROP TABLE IF EXISTS `groups`;
CREATE TABLE IF NOT EXISTS  `groups` (
                                         `id` int NOT NULL AUTO_INCREMENT,
                                         `created_time` bigint not null,
                                         `updated_time` bigint not null,
                                         `is_delete` bool not null default false,
                                         `user_id` int DEFAULT NULL COMMENT '''群主ID''',
                                         `name` varchar(150) DEFAULT NULL COMMENT '''群名称',
                                         `notice` varchar(350) DEFAULT NULL COMMENT '''群公告',
                                         `uuid` varchar(150) NOT NULL COMMENT '''uuid''',
                                         PRIMARY KEY (`id`),
                                         KEY `idx_groups_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT '群组表';


DROP TABLE IF EXISTS `group_members`;
CREATE TABLE  IF NOT EXISTS `group_members` (
                                                `id` int NOT NULL AUTO_INCREMENT,
                                                `created_time` bigint not null,
                                                `updated_time` bigint not null,
                                                `is_delete` bool not null default false,
                                                `user_id` int DEFAULT NULL COMMENT '''用户ID''',
                                                `group_id` int DEFAULT NULL COMMENT '''群组ID''',
                                                `nickname` varchar(350) DEFAULT NULL COMMENT '''昵称',
                                                `mute` smallint DEFAULT NULL COMMENT '''是否禁言''',
                                                PRIMARY KEY (`id`),
                                                KEY `idx_group_members_user_id` (`user_id`),
                                                KEY `idx_group_members_group_id` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT '群组成员表';
