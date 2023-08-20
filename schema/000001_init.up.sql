CREATE TABLE users
(
    id          serial       not null unique primary key ,
    login         varchar(255) not null,
    password         varchar(255) not null

);
drop table users;