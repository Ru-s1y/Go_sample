-- drop table posts cascade if exists;
-- drop table comments if exists;

create table post (
	id		serial primary key,
	content	text,
	author	varchar(255)
);

-- create table comments (
-- 	id		serial primary key,
-- 	content	text,
-- 	author	varchar(255),
-- 	post_id	integer references post(id)
-- );