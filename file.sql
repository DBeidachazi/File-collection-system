drop table if exists `user`;
create table `user`
(
    `stu_id`   int(11) not null,
    `name`     varchar(255) not null,
    `password` varchar(255) not null,
    primary key (`stu_id`)
) engine=innodb default charset=utf8mb4;

-- # 班级表
drop table if exists `class`;
create table `class`
(
    `class_id` int(4) not null auto_increment,
    `name`     varchar(255) not null,
    primary key (`class_id`)
) engine=innodb default charset=utf8mb4;

-- #
-- 班级用户表
drop table if exists `class_user`;
create table `class_user`
(
    `class_id`        int(4) not null,
    `stu_id`          int(11) not null,
    `name`            varchar(255) not null,
    `permission_type` int(11) not null comment '1:创建班级 2:加入班级 可以设置多个管理',
    primary key (`class_id`, `stu_id`)
) engine=innodb default charset=utf8mb4;

-- #
-- 仅管理员可发布
-- # 作业新增
drop table if exists `course`;
create table `course`
(
    `course_id`      int(11) not null auto_increment,
    `name`           varchar(255) not null,
    `class_id`       int(4) not null,
    `publisher_name` varchar(255) not null,
    primary key (`course_id`)
) engine=innodb default charset=utf8mb4;

-- 作业清单
drop table if exists `role`;
create table `role`(
    `user_id` int(11) not null,
    `role_id` int (11) not null auto_increment,
    `role_name` varchar(255) not null,
    `class_id` int(4) not null,
    `status` int(11) not null comment '1:已提交 2:未提交 3:发布者(不显示)'
) engine=innodb default charset=utf8mb4;

-- 作业管理
drop table if exists `worklist`;
create table `worklist`
(
    `work_id`   int(11) not null auto_increment,
    `course_id` int(11) not null,
    `name`      varchar(255) not null,
    `path`      varchar(255) not null,
    `status`    int(11) not null comment '提交人数',
    primary key (`work_id`)
) engine=innodb default charset=utf8mb4;
