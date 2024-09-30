CREATE TABLE roles
(
    id   int not null primary key auto_increment,
    name varchar(50) unique
);

CREATE TABLE users
(
    id                 int          not null primary key auto_increment,
    email              varchar(40)  not null unique,
    encrypted_password varchar(100) not null,
    is_active          bool default 1,
    role_id            int,
    FOREIGN KEY (role_id) REFERENCES roles (id),
    first_name         varchar(40),
    last_name          varchar(40)
);