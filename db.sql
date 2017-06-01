CREATE TABLE users(
	id serial,
	name varchar(100) not null,
	lastName varchar(100) null,
	email varchar(100) null unique,
	active boolean default false,
	primary key(id)
);

CREATE TABLE activation_tokens(
	id serial,
	user_id int not null,
	token varchar not null,
	primary key(id),
	foreign key(user_id) references users(id) on delete cascade
);