create database chatroom;
use chatroom;
drop database chatroom;

create table `user`
(
    `account`  int auto_increment,
    `username` varchar(20),
    `password` varchar(20),
    `img`      varchar(100) default '/static/img/default.jpg',
    primary key (`account`)
) engine = InnoDB
  default charset = utf8mb4;
alter table user
    auto_increment = 101;

create table `private_messages`
(
    `id`               INT AUTO_INCREMENT PRIMARY KEY,
    `sender_account`   INT,
    `receiver_account` INT,
    `message`          TEXT,
    `created_at`       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) engine = InnoDB
  default charset = utf8mb4;

create table `user_relationships`
(
    `id`             INT AUTO_INCREMENT PRIMARY KEY,
    `user_account`   INT NOT NULL,
    `friend_account` INT NOT NULL,
    `status`         TINYINT   DEFAULT 0, -- 0 表示不是好友，1 表示是好友
    `created_at`     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) engine = InnoDB
  default charset = utf8mb4;

create table `groups`
(
    `id`   int auto_increment primary key,
    `name` varchar(20),
    `img`  varchar(100)
) engine = InnoDB
  default charset = utf8mb4;

create table `group_members`
(
    `id`           int auto_increment primary key,
    `group_id`     int,
    `user_account` int,
    `status`       tinyint -- 0 表示是普通成员，1 表示是群主
) engine = InnoDB
  default charset = utf8mb4;

create table `group_messages`
(
    `id`             int auto_increment primary key,
    `sender_account` int,
    `group_id`       int,
    `message`        text,
    `type`           tinyint, -- 0 为世界群聊，1 为普通群聊
    `created_at`     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) engine = InnoDB
  default charset = utf8mb4;

create table `info`
(
    `id`               int auto_increment primary key,
    `sender_account`   int,
    `receiver_account` int,
    `type`             tinyint, -- 0 表示加好友请求信息
    `finish`           tinyint -- 1 表示已完成，0 表示未完成
) engine = InnoDB
  default charset = utf8mb4;
