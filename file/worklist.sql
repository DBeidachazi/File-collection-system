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

