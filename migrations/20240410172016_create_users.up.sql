CREATE TABLE users(
    id int not null primary key auto_increment,
    email varchar(40) not null unique,
    encrypted_password varchar(100) not null,
    is_active bool default 1,
    role_ varchar(20) default "superuser", 
    first_name varchar(40),
    last_name varchar(40) 
);