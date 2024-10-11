create table users (
	id bigserial primary key,
	email varchar(255) unique not null,
	name varchar(100) not null,
	image_url text not null,
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);