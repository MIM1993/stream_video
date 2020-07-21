#创建user表
create table users(
    id INT unsigned not null auto_increment,
    login_name VARCHAR (64),
    pwd TEXT not null,
    unique key(login_name),
    primary key(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;



#创建sessions表
create table sessions(
         session_id VARCHAR(64) not null,
         TTL TinyText,
         login_name VARCHAR(64),
         primary key(session_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;