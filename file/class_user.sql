create table file.class_user
(
    class_id        int(4)       not null,
    stu_id          int          not null,
    name            varchar(255) not null,
    permission_type int          not null comment '1:管理员 2:普通用户',
    username        varchar(255) not null
)
    charset = utf8mb4;

