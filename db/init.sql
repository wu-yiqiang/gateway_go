drop database wserver;
create database if not exists wserver;
use wserver;
drop table if exists `user`;
create table `user` (
  `id` int not null auto_increment,
  `username` varchar(40) not null,
  `password` varchar(40) not null,
  `nickname` varchar(40) not null,
  `email` varchar(40) not null,
  `phone` varchar(40) not null,
  `create_time` timestamp not null default CURRENT_TIMESTAMP,
  `update_time` timestamp not null default CURRENT_TIMESTAMP,
  `is_delete` bool not null default false,
  PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO user ( username, password, nickname, email, phone ) VALUES ( 'sutter', 'password123', 'sutter', 'wu_yiqiang@outlook.com', '15770870823');
INSERT INTO user ( username, password, nickname, email, phone ) VALUES ( 'altas', 'passwd123', 'atals', 'wu_yiqiang@aliyun.com', '15770870825');