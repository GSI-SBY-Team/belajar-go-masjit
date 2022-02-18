create table user_masjit (
	id varchar(36) primary key, 
	username varchar(100) unique,
	password varchar(50)
);
insert into user_masjit values('8080', 'Abdul Masjit Subekti', '123');
