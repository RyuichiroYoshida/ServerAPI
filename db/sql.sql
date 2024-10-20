create user 'devtest'@'localhost' identified by 'devtest';
create database if not exists dev;
use dev;
create table user (
    json json
);
