create table file.user
(
    stu_id   int          not null
        primary key,
    name     varchar(255) not null,
    password varchar(255) not null,
    username varchar(255) not null
)
    charset = utf8mb4;

