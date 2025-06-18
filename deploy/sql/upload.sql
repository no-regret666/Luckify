drop database if exists `go-lottery-upload`;
create database `go-lottery-upload` character set = utf8mb4 collate = utf8mb4_general_ci;

use `go-lottery-upload`;
drop table if exists `upload_file`;
create table `upload_file`(
                              id int not null primary key auto_increment,
                              user_id int not null comment '上传用户id',
                              file_name varchar(100) character set utf8mb4 collate utf8mb4_general_ci not null comment '文件名',
    ext varchar(50) character set utf8mb4 collate utf8mb4_general_ci not null comment '扩展名',
    size int not null comment '文件大小',
    url varchar(255) character set utf8mb4 collate utf8mb4_general_ci not null comment '下载链接',
    create_time datetime not null default current_timestamp,
    update_time datetime not null default current_timestamp on update current_timestamp
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_0900_ai_ci row_format=dynamic comment='文件上传表';