

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";  


DROP TABLE  IF EXISTS  bank;
create table bank (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"desc" TEXT,
	update_date BIGINT,
	"link_url" TEXT
);
INSERT INTO bank ("id", "name", "desc", "update_date", "link_url")VALUES(uuid_generate_v4(),'台新銀行','', 1624369053, '');
INSERT INTO bank ("id", "name", "desc", "update_date", "link_url")VALUES(uuid_generate_v4(),'永豐銀行','', 1624369053, '');




DROP TABLE  IF EXISTS  card;
create table card (
    "id" VARCHAR(36) PRIMARY KEY,
    bank_id VARCHAR(36), 
	"name" VARCHAR(100),
	"desc" TEXT,
	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT,
	"link_url" TEXT,
    FOREIGN KEY(bank_id) REFERENCES BANK("id")
);



DROP TABLE  IF EXISTS reward;
create table reward (
    "id" VARCHAR(36) PRIMARY KEY,
    card_id VARCHAR(36), 
	"name" VARCHAR(100),
	"desc" TEXT,
	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT,
    constraint_payload JSON,
    FOREIGN KEY(card_id) REFERENCES CARD("id")
);


DROP TABLE  IF EXISTS customization; 
CREATE TABLE customization (
    "id" VARCHAR(36) PRIMARY KEY,
    "name" VARCHAR(100),
    "desc" TEXT,
	"link_url" TEXT
);

INSERT INTO customization("id", "name", "desc", "link_url") 
	VALUES(uuid_generate_v4(), '回饋無上限', '', '');

INSERT INTO customization("id", "name", "desc", "link_url") 
	VALUES(uuid_generate_v4(), 'Richart帳戶自動扣繳@GoGo卡信用卡帳單，且@GoGo卡消費金額滿NT$5,000(含)', '', '');


INSERT INTO customization("id", "name", "desc", "link_url") 
	VALUES(uuid_generate_v4(), '永豐幣倍卡-任務一', '綁定本行新台幣帳戶(含數位帳戶)自動扣繳信用卡帳款設定完成且扣款成功。及②使用電子或行動帳單完成設定且寄送成功並取消實體帳單。前述兩項須同時必備達成始符合【任務一】；若僅達成其中一項，恕無法提供加碼回饋。', '');

INSERT INTO customization("id", "name", "desc", "link_url") 
	VALUES(uuid_generate_v4(), '永豐幣倍卡-任務二(懂匯)', '於本行①台外幣帳戶間換匯單筆金額或②外幣存款月平均餘額或③臨櫃投保外幣保單月扣繳單筆金額達等值台幣：1元~10萬元 屬【懂匯】資格；10萬元以上屬【超匯】資格。', '');

INSERT INTO customization("id", "name", "desc", "link_url") 
	VALUES(uuid_generate_v4(), '永豐幣倍卡-任務二(超匯)', '於本行①台外幣帳戶間換匯單筆金額或②外幣存款月平均餘額或③臨櫃投保外幣保單月扣繳單筆金額達等值台幣：1元~10萬元 屬【懂匯】資格；10萬元以上屬【超匯】資格。', '');


INSERT INTO customization("id", "name", "desc", "link_url") 
	VALUES(uuid_generate_v4(), '豐城海外村', '豐城網頁版進入海外村內任一商店，輸入”身分證字號+生日”就能作為登入依據，或可從豐城APP版(下載汗水不白流APP)登入後點選【豐城】再連結海外村內任一商店，APP登入且點選【豐城】紀錄就能作為導購流程的依據。幣倍卡持卡人須有豐城登入紀錄且登入後24小時內，透過點擊連結至海外村內任一商家並成功以幣倍卡完成刷卡消費，即可納入計算。', '');


DROP TABLE  IF EXISTS mobilepay;
create table mobilepay (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
    "desc" TEXT,
	"link_url" TEXT,
	PRIMARY KEY("id")
);

INSERT INTO public.mobilepay(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'LINE Pay',  '連線支付', '');
INSERT INTO public.mobilepay(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'GOOGLE Pay',  'Google支付', '');
INSERT INTO public.mobilepay(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'APPLE Pay', 'Apple支付', '');
INSERT INTO public.mobilepay(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'My FamiPay', '全家超商支付', '');
INSERT INTO public.mobilepay(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Taishin PAY', '台新支付', '');
INSERT INTO public.mobilepay(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Open Wallet', 'OPEN錢包', '');
INSERT INTO public.mobilepay(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Easy Wallet', 'Easy Wallet 悠遊付', '');
INSERT INTO public.mobilepay(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'PX Pay',  '全聯支付', '');
INSERT INTO public.mobilepay(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Samsung Pay',  '三星支付', '');
INSERT INTO public.mobilepay(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Garmin Pay',  'Garmin Pay', '');
INSERT INTO public.mobilepay(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Fitbit Pay',  'Fitbit Pay', '');





DROP TABLE  IF EXISTS ecommerce;
create table ecommerce (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
    "desc" TEXT,
	"link_url" TEXT,
	PRIMARY KEY("id")
);

INSERT INTO public.ecommerce(id, "name", "desc", "link_url")values (uuid_generate_v4(), '蝦皮購物',  '蝦皮購物', '');
INSERT INTO public.ecommerce(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'MOMO',  'MOMO購物', '');
INSERT INTO public.ecommerce(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'PChome',  'PChome', '');
INSERT INTO public.ecommerce(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Amazon',  'Amazon', '');
INSERT INTO public.ecommerce(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Gmarket', 'Gmarket', '');
INSERT INTO public.ecommerce(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Decathlon', '迪卡儂線上購物', '');



DROP TABLE  IF EXISTS supermarket;
create table supermarket (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
    "desc" TEXT,
	"link_url" TEXT,
	PRIMARY KEY("id")
);

INSERT INTO public.supermarket(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Px Mart', '全聯福利中心', '');





DROP TABLE  IF EXISTS onlinegame;
create table onlinegame (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
    "desc" TEXT,
	"link_url" TEXT
);

INSERT INTO public.onlinegame(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Nintendo',  '任天堂', '');
INSERT INTO public.onlinegame(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'MyCard',  'MyCard', '');
INSERT INTO public.onlinegame(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Beanfun', '遊戲橘子', '');




DROP TABLE  IF EXISTS streaming;
create table streaming (
    "id" VARCHAR(36),
	"name" VARCHAR(100),
    "desc" TEXT,
	"link_url" TEXT
);

INSERT INTO public.streaming(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Spotify', 'Spotify', '');
INSERT INTO public.streaming(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'Netflix', 'Netflix', '');
INSERT INTO public.streaming(id, "name", "desc", "link_url")values (uuid_generate_v4(), 'CATCHPLAY+', 'CATCHPLAY+', '');
