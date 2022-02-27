

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";  


DROP TABLE  IF EXISTS  bank;
create table bank (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	update_date BIGINT,
	"image_path" TEXT,
	"link_url" TEXT
);
INSERT INTO bank ("id", "name", "update_date", "image_path", "link_url")VALUES(uuid_generate_v4(),'台新銀行', 1645533270, '', '');
INSERT INTO bank ("id", "name", "update_date", "image_path", "link_url")VALUES(uuid_generate_v4(),'永豐銀行', 1645533270, '', '');




DROP TABLE  IF EXISTS  card;
create table card (
    "id" VARCHAR(36) PRIMARY KEY,
    bank_id VARCHAR(36), 
	"name" VARCHAR(100),
	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT,
	"image_path" TEXT,
	"link_url" TEXT,
    FOREIGN KEY(bank_id) REFERENCES BANK("id")
);



DROP TABLE  IF EXISTS reward;
create table reward (
    "id" VARCHAR(36) PRIMARY KEY,
    card_id VARCHAR(36), 
	"order" INT, 
	"title" TEXT,
	"sub_title" TEXT,

	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT,
	
	"reward_type" INT,
	payload_operator INT, 
    payload JSON,
	feedback JSON,
	FOREIGN KEY(card_id) REFERENCES card("id")
);


DROP TABLE  IF EXISTS mobilepay;
create table mobilepay (
    "id" VARCHAR(36)PRIMARY KEY,
	"name" VARCHAR(100),
	"image_path" TEXT
);

INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'LINE Pay', '');
INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'GOOGLE Pay',   '');
INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'APPLE Pay',  '');
INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'My FamiPay', '');
INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'Taishin PAY', '');
INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'Open Wallet',  '');
INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), '悠遊付', '');
INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'PX Pay', '');
INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'Samsung Pay', '');
INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'Garmin Pay', '');
INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'Fitbit Pay', '');




DROP TABLE  IF EXISTS delivery;
create table delivery (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"image_path" TEXT
);

INSERT INTO public.delivery(id, "name", "image_path")values (uuid_generate_v4(), 'FOOD PANDA', '');
INSERT INTO public.delivery(id, "name", "image_path")values (uuid_generate_v4(), 'Uber Eats', '');





DROP TABLE  IF EXISTS ecommerce;
create table ecommerce (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"image_path" TEXT,
	PRIMARY KEY("id")
);

INSERT INTO public.ecommerce(id, "name", "image_path")values (uuid_generate_v4(), '蝦皮購物', '');
INSERT INTO public.ecommerce(id, "name", "image_path")values (uuid_generate_v4(), 'MOMO','');
INSERT INTO public.ecommerce(id, "name", "image_path")values (uuid_generate_v4(), 'PChome', '');
INSERT INTO public.ecommerce(id, "name", "image_path")values (uuid_generate_v4(), 'Amazon', '');
INSERT INTO public.ecommerce(id, "name", "image_path")values (uuid_generate_v4(), 'Gmarket', '');
INSERT INTO public.ecommerce(id, "name", "image_path")values (uuid_generate_v4(), 'Decathlon', '');



DROP TABLE  IF EXISTS supermarket;
create table supermarket (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"image_path" TEXT,
	PRIMARY KEY("id")
);

INSERT INTO public.supermarket(id, "name", "image_path")values (uuid_generate_v4(), '全聯福利中心', '');





DROP TABLE  IF EXISTS onlinegame;
create table onlinegame (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"image_path" TEXT
);

INSERT INTO public.onlinegame(id, "name", "image_path")values (uuid_generate_v4(), '任天堂', '');
INSERT INTO public.onlinegame(id, "name", "image_path")values (uuid_generate_v4(), 'MyCard',  '');
INSERT INTO public.onlinegame(id, "name", "image_path")values (uuid_generate_v4(), '遊戲橘子', '');




DROP TABLE  IF EXISTS streaming;
create table streaming (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"image_path" TEXT
);

INSERT INTO public.streaming(id, "name", "image_path")values (uuid_generate_v4(), 'Spotify','');
INSERT INTO public.streaming(id, "name", "image_path")values (uuid_generate_v4(), 'Netflix', '');
INSERT INTO public.streaming(id, "name", "image_path")values (uuid_generate_v4(), 'CATCHPLAY+', '');
