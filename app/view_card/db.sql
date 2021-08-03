/**
CREATE DATABASE view_card
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;
**/


DROP TABLE  if exists bank;
create table bank (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	icon VARCHAR(100),
	update_date BIGINT
);

DROP TABLE  if exists card;
create table card (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	icon VARCHAR(100),
	bank_id VARCHAR(36),
	start_time BIGINT,
	end_time BIGINT,
	max_point float,
	feature_desc VARCHAR(255),
	applicant_qualifications JSON,
	update_date BIGINT
);


DROP TABLE  if exists reward;
create table reward (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	card_id VARCHAR(36),
	"desc" VARCHAR(255),
	reward_type INTEGER,
	operator_type INTEGER,
	total_point float,
	start_time BIGINT,
	end_time BIGINT,
	update_date BIGINT
);




DROP TABLE  if exists task;
create table task (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"desc" VARCHAR(255),
	reward_id VARCHAR(36),
	point float,
	update_date BIGINT
);



DROP TABLE  if exists feature;
create table feature (
    card_id VARCHAR(36),
	"type" smallint,
	PRIMARY KEY(card_id, "type")
);