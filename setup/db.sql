DROP TABLE  IF EXISTS  bank;
create table bank (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"desc" TEXT,
	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT
);
INSERT INTO bank ("id", "name", "desc", "start_date", "end_date", "update_date")VALUES('c6f9c053-2ccd-4178-9d42-9853e950d500','台新銀行','',1624369053,1624369053,1624369053);

DROP TABLE  IF EXISTS  card;
create table card (
    "id" VARCHAR(36) PRIMARY KEY,
    bank_id VARCHAR(36), 
	"name" VARCHAR(100),
	"desc" TEXT,
	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT,
    FOREIGN KEY(bank_id) REFERENCES BANK("id")
);


INSERT INTO card ("id", "bank_id", "name", "desc", "start_date", "end_date", "update_date") VALUES('96d7ba6a-227a-45f2-8b78-f92223353316', 'c6f9c053-2ccd-4178-9d42-9853e950d500', 'Richart GoGo卡', '', 1624369053,1624369053,1624369053);

DROP TABLE  IF EXISTS reward;
create table reward (
    "id" VARCHAR(36) PRIMARY KEY,
    card_id VARCHAR(36), 
	"name" VARCHAR(100),
	"desc" TEXT,
    operator INTEGER,
	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT,
    cost JSON,
    constraints JSON,
    FOREIGN KEY(card_id) REFERENCES CARD("id")
);


DROP TABLE  IF EXISTS customization; 
CREATE TABLE customization (
    "id" VARCHAR(36) PRIMARY KEY,
    "reward_id" VARCHAR(36),
    "name" VARCHAR(100),
    "desc" TEXT,
	FOREIGN KEY(reward_id) REFERENCES reward("id")
);

/*
INSERT INTO customization ("id", "reward_id", "name", "desc")  VALUES ('aa', '4f84a5cb-d54d-4c83-541c-1dfaed707a8a', 'Richard帳戶自動扣繳於GoGo信用卡帳單', '');
*/


DROP TABLE  IF EXISTS mobilepay;
create table mobilepay (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"action" INTEGER,
    "desc" TEXT,
	PRIMARY KEY("id", "action")
);

INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('LinePay', 'LINE Pay', 0,  '連線支付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('GooglePay', 'GOOGLE Pay', 0,  'Google支付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('ApplePay', 'APPLE Pay', 0,  'Apple支付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('MyFamiPay', 'My FamiPay', 0,  '全家超商支付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('TaishinPay', 'Taishin PAY', 0,  '台新支付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('OpenWallet', 'Open Wallet', 0,  'OPEN錢包');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('EasyWallet', 'Easy Wallet', 0,  'Easy Wallet 悠遊付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('PXPay', 'PX Pay', 0,  '全聯支付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('PXPay', 'PX Pay', 1,  '全聯支付');





DROP TABLE  IF EXISTS ecommerce;
create table ecommerce (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"action" INTEGER,
    "desc" TEXT,
	PRIMARY KEY("id", "action")
);

INSERT INTO public.ecommerce(id, "name", "action", "desc")values ('Shopee', 'Shopee', 0,  '蝦皮購物');
INSERT INTO public.ecommerce(id, "name", "action", "desc")values ('MOMO', 'MOMO', 0,  'MOMO購物');
INSERT INTO public.ecommerce(id, "name", "action", "desc")values ('PChome', 'PChome', 0,  'PChome');
INSERT INTO public.ecommerce(id, "name", "action", "desc")values ('Amazon', 'Amazon', 0,  'Amazon');
INSERT INTO public.ecommerce(id, "name", "action", "desc")values ('Gmarket', 'Gmarket', 0,  'Gmarket');
INSERT INTO public.ecommerce(id, "name", "action", "desc")values ('DECATHLON', 'Decathlon', 0,  '迪卡儂線上購物');



DROP TABLE  IF EXISTS supermarket;
create table supermarket (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"action" INTEGER,
    "desc" TEXT,
	PRIMARY KEY("id", "action")
);

INSERT INTO public.supermarket(id, "name", "action", "desc")values ('PxMart', 'Px Mart', 0,  '全聯福利中心');





DROP TABLE  IF EXISTS onlinegame;
create table onlinegame (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"action" INTEGER,
    "desc" TEXT
);

INSERT INTO public.onlinegame(id, "name", "action", "desc")values ('Nintendo', 'Nintendo', 0,  '任天堂');
INSERT INTO public.onlinegame(id, "name", "action", "desc")values ('MyCard', 'MyCard', 0,  'MyCard');
INSERT INTO public.onlinegame(id, "name", "action", "desc")values ('Beanfun', 'Beanfun', 0,  '遊戲橘子');




DROP TABLE  IF EXISTS streaming;
create table streaming (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
	"action" INTEGER,
    "desc" TEXT
);

INSERT INTO public.streaming(id, "name", "action", "desc")values ('Spotify', 'Spotify', 0,  'Spotify');
INSERT INTO public.streaming(id, "name", "action", "desc")values ('Netflix', 'Netflix', 0,  'Netflix');
INSERT INTO public.streaming(id, "name", "action", "desc")values ('CATCHPLAY+', 'CATCHPLAY+', 0,  'CATCHPLAY+');
