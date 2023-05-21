create table file.class
(
    class_id   int(4) auto_increment
        primary key,
    name       varchar(255) not null,
    class_name varchar(255) not null
)
    charset = utf8mb4;

create table file.class_user
(
    class_id        int(4)       not null,
    stu_id          int          not null,
    name            varchar(255) not null,
    permission_type int          not null comment '1:管理员 2:普通用户',
    username        varchar(255) not null
)
    charset = utf8mb4;

create table file.course
(
    course_id      int auto_increment
        primary key,
    name           varchar(255) not null,
    class_id       int(4)       not null,
    publisher_name varchar(255) not null,
    deadline       datetime     null comment '截止时间',
    stu_id         int          not null,
    file_type      varchar(255) not null,
    username       varchar(255) not null,
    course_name    varchar(255) not null
)
    charset = utf8mb4;

create table file.role
(
    user_id   int          not null,
    role_id   int auto_increment
        primary key,
    role_name varchar(255) not null,
    class_id  int(4)       not null,
    status    int          not null comment '1:已提交 2:未提交 3:发布者(不显示)',
    stu_id    int          not null,
    deadline  datetime     not null,
    course_id int          not null
)
    charset = utf8mb4;

create table file.user
(
    stu_id   int          not null
        primary key,
    name     varchar(255) not null,
    password varchar(255) not null,
    username varchar(255) not null
)
    charset = utf8mb4;

create table file.worklist
(
    work_id     int auto_increment
        primary key,
    course_id   int          not null,
    name        varchar(255) not null,
    path        varchar(255) not null,
    status      int          not null comment '提交人数',
    stu_id      int          not null,
    course_name varchar(255) null,
    deadline    datetime     not null
)
    charset = utf8mb4;

