drop database wserver;
create database if not exists wserver;
use wserver;
create table IF NOT EXISTS `users` (
  `id` int not null auto_increment,
  `username` varchar(40) not null,
  `password` varchar(400) not null,
  `nickname` varchar(40) not null,
  `email` varchar(40) not null,
  `phone` varchar(11) not null,
  `role` varchar(400),
  `create_time` bigint not null,
  `update_time` bigint not null,
  `is_delete` bool not null default false,
  PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO users ( username, password, nickname, email, phone, role, create_time, update_time ) VALUES ( 'sutter', 'password123', 'sutter', 'wu_yiqiang@outlook.com', '15770870823', 'admin, sys_admin', 123456799, 123456799);
INSERT INTO users ( username, password, nickname, email, phone, role, create_time, update_time ) VALUES ( 'altas', 'passwd123', 'atals', 'wu_yiqiang@aliyun.com', '15770870825', 'admin', 123456789, 123456789);

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
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO menus ( name, permission, description, create_time, update_time ) VALUES ( 'user:view', '查看', '查看权限', 112343,112343);
INSERT INTO menus ( name, permission, description, create_time, update_time ) VALUES ( 'user:update', '更新', '更新权限', 112353,112353);
INSERT INTO menus ( name, permission, description, create_time, update_time ) VALUES ( 'user:other', '特殊', '特殊权限', 112363,112363);