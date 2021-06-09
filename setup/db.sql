

DROP TABLE bank;
create table bank (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"desc" TEXT,
	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT
);


DROP TABLE card;
create table card (
    "id" VARCHAR(36) PRIMARY KEY,
    bank_id VARCHAR(36), 
	"name" varchar(100),
	"desc" TEXT,
	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT,
    FOREIGN KEY(bank_id) REFERENCES BANK("id")
);


DROP TABLE privilage;
create table privilage (
    "id" VARCHAR(36) PRIMARY KEY,
    card_id VARCHAR(36), 
	"name" varchar(100),
	"desc" TEXT,
	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT,
	score float8,
    FOREIGN KEY(card_id) REFERENCES CARD("id")
);

DROP TABLE "constraint";
create table "constraint" (
    "id" VARCHAR(36) PRIMARY KEY,
    privilage_id VARCHAR(36), 
	"name" varchar(100),
	"desc" TEXT,
    "operator" INTEGER,
	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT,
    FOREIGN KEY(privilage_id) REFERENCES PRIVILAGE("id")
);






DROP TABLE mobilepay;
create table mobilepay (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" varchar(100),
	"action" INTEGER,
    "desc" TEXT
);

INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('LinePay-0', 'LINE Pay', 0,  '連線支付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('GooglePay-0', 'GOOGLE Pay', 0,  'Google支付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('ApplePay-0', 'APPLE Pay', 0,  'Apple支付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('MyFamiPay-0', 'My FamiPay', 0,  '全家超商支付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('TaishinPay-0', 'Taishin PAY', 0,  '台新支付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('OpenWallet-0', 'Open Wallet', 0,  'OPEN錢包');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('EasyWallet-0', 'Easy Wallet', 0,  'Easy Wallet 悠遊付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('PXPay-0', 'PX Pay', 0,  '全聯支付');
INSERT INTO public.mobilepay(id, "name", "action", "desc")values ('PXPay-1', 'PX Pay', 1,  '全聯支付');





DROP TABLE ecommerce;
create table ecommerce (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" varchar(100),
	"action" INTEGER,
    "desc" TEXT
);

INSERT INTO public.ecommerce(id, "name", "action", "desc")values ('Shopee-0', 'Shopee', 0,  '蝦皮購物');
INSERT INTO public.ecommerce(id, "name", "action", "desc")values ('MOMO-0', 'MOMO', 0,  'MOMO購物');
INSERT INTO public.ecommerce(id, "name", "action", "desc")values ('PCHOME-0', 'PCHOME', 0,  'PCHOME');
INSERT INTO public.ecommerce(id, "name", "action", "desc")values ('Amazon-0', 'Amazon', 0,  'Amazon');
INSERT INTO public.ecommerce(id, "name", "action", "desc")values ('Gmarket-0', 'Gmarket', 0,  'Gmarket');
INSERT INTO public.ecommerce(id, "name", "action", "desc")values ('DECATHLON-0', 'Decathlon', 0,  '迪卡儂線上購物');



DROP TABLE supermarket;
create table supermarket (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" varchar(100),
	"action" INTEGER,
    "desc" TEXT
);

INSERT INTO public.supermarket(id, "name", "action", "desc")values ('PxMart-0', 'Px Mart', 0,  '全聯福利中心');





DROP TABLE onlinegame;
create table onlinegame (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" varchar(100),
	"action" INTEGER,
    "desc" TEXT
);

INSERT INTO public.onlinegame(id, "name", "action", "desc")values ('Nintendo-0', 'Nintendo', 0,  '任天堂');
INSERT INTO public.onlinegame(id, "name", "action", "desc")values ('MyCard-0', 'MyCard', 0,  'MyCard');
INSERT INTO public.onlinegame(id, "name", "action", "desc")values ('Beanfun-0', 'Beanfun', 0,  '遊戲橘子');




DROP TABLE streaming;
create table streaming (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" varchar(100),
	"action" INTEGER,
    "desc" TEXT
);

INSERT INTO public.streaming(id, "name", "action", "desc")values ('Spotify-0', 'Spotify', 0,  'Spotify');
INSERT INTO public.streaming(id, "name", "action", "desc")values ('Netflix-0', 'Netflix', 0,  'Netflix');
INSERT INTO public.streaming(id, "name", "action", "desc")values ('CATCHPLAY+-0', 'CATCHPLAY+', 0,  'CATCHPLAY+');


