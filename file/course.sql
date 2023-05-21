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

