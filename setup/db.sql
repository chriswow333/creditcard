

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";  


DROP TABLE  IF EXISTS bank;
create table bank (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	update_date BIGINT,
	"image_path" TEXT,
	"link_url" TEXT
);
INSERT INTO bank ("id", "name", "update_date", "image_path", "link_url")VALUES(uuid_generate_v4(),'台新銀行', 1645533270, '', '');
INSERT INTO bank ("id", "name", "update_date", "image_path", "link_url")VALUES(uuid_generate_v4(),'永豐銀行', 1645533270, '', '');


DROP TABLE  IF EXISTS card;
create table card (
    "id" VARCHAR(36) PRIMARY KEY,
    bank_id VARCHAR(36), 
	"name" VARCHAR(100),
	update_date BIGINT,
	"image_path" TEXT,
	"link_url" TEXT,
	card_status INT,
	"other_reward" JSON,
    FOREIGN KEY(bank_id) REFERENCES BANK("id")
);

DROP TABLE  IF EXISTS card_reward;
create table card_reward (
    "id" VARCHAR(36) PRIMARY KEY,
	"card_id" VARCHAR(36), 
	"card_reward_operator" INT,
	"title" VARCHAR(36), 
	"descs" VARCHAR(36), 
	"start_date" BIGINT,
	"end_date" BIGINT,
	"reward_type" INT,
	"constraint_pass_logic" JSON, 
	"card_reward_bonus" JSON,
	"card_reward_limit_types" JSON,
    FOREIGN KEY(card_id) REFERENCES card("id")
);



DROP TABLE  IF EXISTS reward;
create table reward (
    "id" VARCHAR(36) PRIMARY KEY,
    card_reward_id VARCHAR(36), 
	"order" INT, 

	payload_operator INT, 
    payload JSON,
	FOREIGN KEY(card_reward_id) REFERENCES card_reward("id")
);


DROP TABLE  IF EXISTS task;
create table task (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"descs" JSON,
	"card_id" VARCHAR(36),
	"task_type" INT,
	"task_type_model" JSON,
	"default_pass" BOOLEAN, 
	FOREIGN KEY(card_id) REFERENCES card("id")
);


-- INSERT INTO public.customization(id, "name", "card_id", "default_pass")values (uuid_generate_v4() , '基本0.2%', '60e45bac-61f5-4c6e-4d88-1f09e04599af', true);
-- INSERT INTO public.customization(id, "name", "card_id", "default_pass")values (uuid_generate_v4() , '任務', '60e45bac-61f5-4c6e-4d88-1f09e04599af', false);


DROP TABLE  IF EXISTS mobilepay;
create table mobilepay (
    "id" VARCHAR(36)PRIMARY KEY,
	"name" VARCHAR(100),
	"image_path" TEXT
);

-- INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'LINE Pay', '');
-- INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'GOOGLE Pay',   '');
-- INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'APPLE Pay',  '');
-- INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'My FamiPay', '');
-- INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'Taishin PAY', '');
-- INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'Open Wallet',  '');
-- INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), '悠遊付', '');
-- INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'PX Pay', '');
-- INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'Samsung Pay', '');
-- INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'Garmin Pay', '');
-- INSERT INTO public.mobilepay(id, "name", "image_path")values (uuid_generate_v4(), 'Fitbit Pay', '');




DROP TABLE  IF EXISTS delivery;
create table delivery (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"image_path" TEXT
);

-- INSERT INTO public.delivery(id, "name", "image_path")values (uuid_generate_v4(), 'FOOD PANDA', '');
-- INSERT INTO public.delivery(id, "name", "image_path")values (uuid_generate_v4(), 'Uber Eats', '');





DROP TABLE  IF EXISTS ecommerce;
create table ecommerce (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"image_path" TEXT,
	PRIMARY KEY("id")
);

-- INSERT INTO public.ecommerce(id, "name", "image_path")values (uuid_generate_v4(), '蝦皮購物', '');
-- INSERT INTO public.ecommerce(id, "name", "image_path")values (uuid_generate_v4(), 'MOMO','');
-- INSERT INTO public.ecommerce(id, "name", "image_path")values (uuid_generate_v4(), 'PChome', '');
-- INSERT INTO public.ecommerce(id, "name", "image_path")values (uuid_generate_v4(), 'Amazon', '');
-- INSERT INTO public.ecommerce(id, "name", "image_path")values (uuid_generate_v4(), 'Gmarket', '');
-- INSERT INTO public.ecommerce(id, "name", "image_path")values (uuid_generate_v4(), 'Decathlon', '');



DROP TABLE  IF EXISTS supermarket;
create table supermarket (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"image_path" TEXT,
	PRIMARY KEY("id")
);

-- INSERT INTO public.supermarket(id, "name", "image_path")values (uuid_generate_v4(), '全聯福利中心', '');





DROP TABLE  IF EXISTS onlinegame;
create table onlinegame (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"image_path" TEXT
);

-- INSERT INTO public.onlinegame(id, "name", "image_path")values (uuid_generate_v4(), '任天堂', '');
-- INSERT INTO public.onlinegame(id, "name", "image_path")values (uuid_generate_v4(), 'MyCard',  '');
-- INSERT INTO public.onlinegame(id, "name", "image_path")values (uuid_generate_v4(), '遊戲橘子', '');




DROP TABLE  IF EXISTS streaming;
create table streaming (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"image_path" TEXT
);

-- INSERT INTO public.streaming(id, "name", "image_path")values (uuid_generate_v4(), 'Spotify','');
-- INSERT INTO public.streaming(id, "name", "image_path")values (uuid_generate_v4(), 'Netflix', '');
-- INSERT INTO public.streaming(id, "name", "image_path")values (uuid_generate_v4(), 'CATCHPLAY+', '');




DROP TABLE  IF EXISTS retail;
create table retail (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"image_path" TEXT
);


DROP TABLE  IF EXISTS reward_channel;
create table reward_channel (
    "id" VARCHAR(36),
	"order" INT,
	"card_id" VARCHAR(36),
	"card_reward_id" VARCHAR(36),
	"chaennel_id" VARCHAR(36),
	"channel_type" INT,
	PRIMARY KEY("id")
);



DROP TABLE  IF EXISTS food;
create table food (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100)
);




DROP TABLE  IF EXISTS cashback;
create table cashback (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"descs" JSON
);



DROP TABLE  IF EXISTS pointback;
create table pointback (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"descs" JSON
);


DROP TABLE  IF EXISTS redback;
create table redback (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"descs" JSON
);