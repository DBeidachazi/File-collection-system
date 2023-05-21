create table file.class
(
    class_id   int(4) auto_increment
        primary key,
    name       varchar(255) not null,
    class_name varchar(255) not null
)
    charset = utf8mb4;

