create table users (
UserID serial primary key,
email text not null unique,
password text not null
);
Insert into users (email,password) values ('xyz@test.com', 'sfks');
Insert into users (email,password) values ('abc@test.com', 'test')

