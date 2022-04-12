

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
	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT,
	"image_path" TEXT,
	"link_url" TEXT,
    FOREIGN KEY(bank_id) REFERENCES BANK("id")
);
DROP TABLE  IF EXISTS card_reward;
create table card_reward (
    "id" VARCHAR(36) PRIMARY KEY,
	"card_id" VARCHAR(36), 
	"card_reward_desc" TEXT,
	"card_reward_operator" INT,
	"reward_type" INT,
	"constraint_pass_logic" TEXT, 
    FOREIGN KEY(card_id) REFERENCES card("id")
);



DROP TABLE  IF EXISTS reward;
create table reward (
    "id" VARCHAR(36) PRIMARY KEY,
    card_reward_id VARCHAR(36), 
	"order" INT, 
	"title" TEXT,
	"sub_title" TEXT,

	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT,
	
	payload_operator INT, 
    payload JSON,
	FOREIGN KEY(card_reward_id) REFERENCES card_reward("id")
);

DROP TABLE  IF EXISTS customization;
create table customization (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"desc" TEXT,
	"card_id" VARCHAR(36),
	"default_pass" BOOLEAN, 
	FOREIGN KEY(card_id) REFERENCES card("id")
);

INSERT INTO public.customization(id, "name", "desc", "card_id", "default_pass")values (uuid_generate_v4() , '任務一', '綁定電子或行動帳單且設定本行台外幣帳戶自動扣繳帳款\n※若您於本行數位帳戶開戶且同時申辦幣倍卡，則您已符合任務一', 'cfae77f6-4eff-4112-5053-b129889e3ebb', false);
INSERT INTO public.customization(id, "name", "desc", "card_id", "default_pass")values (uuid_generate_v4() , '任務二(1)', '與本行往來符合以下任一且達等值台幣金額1元以上~\n①外幣存款月平均餘額或\n②台外幣帳戶間換匯單筆金額或\n③臨櫃投保外幣保單月扣繳單筆金額\n※1元~10萬元屬【懂匯】資格', 'cfae77f6-4eff-4112-5053-b129889e3ebb', false);
INSERT INTO public.customization(id, "name", "desc", "card_id", "default_pass")values (uuid_generate_v4() , '任務二(2)', '與本行往來符合以下任一且達等值台幣金額1元以上~\n①外幣存款月平均餘額或\n②台外幣帳戶間換匯單筆金額或\n③臨櫃投保外幣保單月扣繳單筆金額\n※10萬元以上屬【超匯】資格', 'cfae77f6-4eff-4112-5053-b129889e3ebb', false);
INSERT INTO public.customization(id, "name", "desc", "card_id", "default_pass")values (uuid_generate_v4() , '任務三', '當期帳單之幣倍卡新增一般消費滿2,000元(含)以上', '60e45bac-61f5-4c6e-4d88-1f09e04599af', true);
 INSERT INTO public.customization(id, "name", "desc", "card_id", "default_pass")values (uuid_generate_v4() , '豐城海外村', '註：豐城網頁版進入海外村內任一商店，輸入”身分證字號+生日”就能作為登入依據，或可從豐城APP版(下載汗水不白流APP)登入後點選[豐城]再連結海外村內任一商店，APP登入且點選[豐城]紀錄就能作為導購流程的依據。幣倍卡持卡人須有豐城登入紀錄且登入後24小時內，透過點擊連結至海外村內任一商家並成功以幣倍卡完成刷卡消費，即可納入計算。', 'cfae77f6-4eff-4112-5053-b129889e3ebb', false);
INSERT INTO public.customization(id, "name", "desc", "card_id", "default_pass")values (uuid_generate_v4() , '基本回饋', '回饋無上限', 'cfae77f6-4eff-4112-5053-b129889e3ebb', true);
INSERT INTO public.customization(id, "name", "desc", "card_id", "default_pass")values (uuid_generate_v4() , '國外消費', '限非台灣且非新台幣之一般消費(含實體商店及網路)或商店收單行為國外銀行之一般消費。', 'cfae77f6-4eff-4112-5053-b129889e3ebb', false);

-- INSERT INTO public.customization(id, "name", "card_id", "default_pass")values (uuid_generate_v4() , '基本0.2%', '60e45bac-61f5-4c6e-4d88-1f09e04599af', true);
-- INSERT INTO public.customization(id, "name", "card_id", "default_pass")values (uuid_generate_v4() , '任務', '60e45bac-61f5-4c6e-4d88-1f09e04599af', false);


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
