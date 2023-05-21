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

