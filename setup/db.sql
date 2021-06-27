

DROP TABLE bank;
create table bank (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"desc" TEXT,
	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT
);
INSERT INTO bank ("id", "name", "desc", "start_date", "end_date", "update_date")VALUES('c6f9c053-2ccd-4178-9d42-9853e950d500','台新銀行','',1624369053,1624369053,1624369053);

DROP TABLE card;
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

DROP TABLE reward;
create table reward (
    "id" VARCHAR(36) PRIMARY KEY,
    card_id VARCHAR(36), 
	"name" VARCHAR(100),
	"desc" TEXT,
    operator INTEGER,
	start_date BIGINT,
	end_date BIGINT,
	update_date BIGINT,
    bonus JSON,
    constraints JSON,
    FOREIGN KEY(card_id) REFERENCES CARD("id")
);


DROP TABLE customization; 
CREATE TABLE customization (
    "id" VARCHAR(36) PRIMARY KEY,
    "reward_id" VARCHAR(36),
    "name" VARCHAR(100),
    "descs" JSON,
	FOREIGN KEY(reward_id) REFERENCES reward("id")
)

INSERT INTO customization ("id", "rewardi_id", "name", "descs")  VALUES ('', '4f84a5cb-d54d-4c83-541c-1dfaed707a8a', 'Richard帳戶自動扣繳於GoGo信用卡帳單', );



DROP TABLE mobilepay;
create table mobilepay (
    "id" VARCHAR(36) PRIMARY KEY,
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





DROP TABLE ecommerce;
create table ecommerce (
    "id" VARCHAR(36) PRIMARY KEY,
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



DROP TABLE supermarket;
create table supermarket (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"action" INTEGER,
    "desc" TEXT,
	PRIMARY KEY("id", "action")
);

INSERT INTO public.supermarket(id, "name", "action", "desc")values ('PxMart', 'Px Mart', 0,  '全聯福利中心');





DROP TABLE onlinegame;
create table onlinegame (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"action" INTEGER,
    "desc" TEXT
);

INSERT INTO public.onlinegame(id, "name", "action", "desc")values ('Nintendo', 'Nintendo', 0,  '任天堂');
INSERT INTO public.onlinegame(id, "name", "action", "desc")values ('MyCard', 'MyCard', 0,  'MyCard');
INSERT INTO public.onlinegame(id, "name", "action", "desc")values ('Beanfun', 'Beanfun', 0,  '遊戲橘子');




DROP TABLE streaming;
create table streaming (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"action" INTEGER,
    "desc" TEXT
);

INSERT INTO public.streaming(id, "name", "action", "desc")values ('Spotify', 'Spotify', 0,  'Spotify');
INSERT INTO public.streaming(id, "name", "action", "desc")values ('Netflix', 'Netflix', 0,  'Netflix');
INSERT INTO public.streaming(id, "name", "action", "desc")values ('CATCHPLAY+', 'CATCHPLAY+', 0,  'CATCHPLAY+');





/*

[
    {
        "rewardID":"178e34d3-f2c2-4b76-60f3-cbfce5878edb",
        "name":"任務0.8%",
        "desc":"完成任務0.8%",
        "start_date":1625097600,
        "end_date":1643587200,
        "update_date":1624428666,
        "constraintPayload":{
            "name":"",
            "operator":0,
            "descs":["當期帳單「成功以Richart帳戶自動扣繳@GoGo卡信用卡帳單，且@GoGo卡消費金額滿NT$5,000(含)」"],
            "constraintType":0,
            "constraintPayloads":[
                 {
                    "name":"最高回饋200元",
                    "operator":0,
                    "descs":["200元"],
                    "constraintType":4,
                    "bonusLimit":{
                        "id":"richard-gogo-2021",
                        "richart":"最高回饋200元",
                        "bonusType":0,
                        "atLeast":0,
                        "atMost":200
                    }
                 },
                 {
                    "name":"完成任務0.8%",
                    "operator":0,
                    "descs":["完成任務0.8%"],
                    "constraintType":0,
                    "constraintPayloads":[
                        {
                            "name":"以Richart帳戶自動扣繳@GoGo卡信用卡帳單",
                            "operator":0,
                            "descs":[],
                            "constraintType":1,
                            "customizations":[
                                {
                                    "ID":"10232919-3d2a-41e0-4ef2-166c1159a42d",
                                    "rewardID":"4f84a5cb-d54d-4c83-541c-1dfaed707a8a",
                                    "name":"以Richart帳戶自動扣繳@GoGo卡信用帳單",
                                    "descs":["消費採同品牌項下(含@GoGo悠遊卡/icash/虛擬卡)正附卡合併計算(四捨五入計)，分期以每期入帳金額計算，不含信用卡年費/預借現金/利息/違約金/手續費等非消費交易，其餘消費皆可認列(例：學費/稅/公用事業代扣繳費用(水/電/瓦斯/電信費)…等)。"
                                    ]
                                },
                                {
                                    "ID":"8dc2dc0a-e3a1-474d-63b3-5ae35ee1a12b",
                                    "rewardID":"4f84a5cb-d54d-4c83-541c-1dfaed707a8a",
                                    "name":"GoGo卡消費金額滿NT$5,000(含)",
                                    "descs":["消費採同品牌項下(含@GoGo悠遊卡/icash/虛擬卡)正附卡合併計算(四捨五入計)，分期以每期入帳金額計算，不含信用卡年費/預借現金/利息/違約金/手續費等非消費交易，其餘消費皆可認列(例：學費/稅/公用事業代扣繳費用(水/電/瓦斯/電信費)…等)。"
                                    ]
                                }
                            ]
                        }
                    ]
                }
            ]
        }
    }
]


{

    "rewardID":"4f84a5cb-d54d-4c83-541c-1dfaed707a8a",
    "name":"以Richart帳戶自動扣繳@GoGo卡信用帳單",
    "descs":[
        "消費採同品牌項下(含@GoGo悠遊卡/icash/虛擬卡)正附卡合併計算(四捨五入計)，分期以每期入帳金額計算，不含信用卡年費/預借現金/利息/違約金/手續費等非消費交易，其餘消費皆可認列(例：學費/稅/公用事業代扣繳費用(水/電/瓦斯/電信費)…等)。"
    ]
}


{
    "name":"8大精選，週六3%",
    "cardID":"96d7ba6a-227a-45f2-8b78-f92223353316",
    "desc":"@GoGo卡12大數位通路(網購/Pay/影音)最高6%，指定餐飲/健身/按摩等最高3%！, 綁Pay是標配！網購/健身/享樂也hen會！",
    "startDate":1625097600,
    "endDate":1643587200,
    "updateDate":1624412547,
    "operator":1
}

{
    "name":"基本0.2%",
    "cardID":"96d7ba6a-227a-45f2-8b78-f92223353316",
    "desc":"無任何限制條件皆0.2%",
    "startDate":1625097600,
    "endDate":1643587200,
    "updateDate":1624412547,
    "operator":0
}

*/

/*


[
    {
        "rewardID":"178e34d3-f2c2-4b76-60f3-cbfce5878edb",
        "name":"10大指定2%",
        "desc":"10大指定購物2%",
        "start_date":1625097600,
        "end_date":1643587200,
        "update_date":1624428666,
        "constraintPayload":{
            "name":"",
            "operator":0,
            "descs":["使用行動支付在指定網購平台消費"],
            "constraintType":0,
            "constraintPayloads":[
                 {
                    "name":"最高回饋200元",
                    "operator":1,
                    "descs":["200元"],
                    "constraintType":4,
                    "bonusLimit":{
                        "id":"richard-gogo-2021",
                        "richart":"最高回饋200元",
                        "bonusType":0,
                        "atLeast":0,
                        "atMost":200
                    }
                 },
                 {
                    "name":"10大指定2%",
                    "operator":0,
                    "descs":["在指定平台消費享有2%"],
                    "constraintType":1,
                    "constraintPayloads":[
                        {
                            "name":"網購",
                            "operator":1,
                            "descs":["Amazon", "Gmarket", "迪卡儂線上購物"],
                            "constraintType":6,
                            "ecommerces":[
                                {
                                    "id":"Amazon-0",
                                    "name":"Amazon",
                                    "actionType":0,
                                    "desc":"Amazon"
                                },
                                {
                                    "id":"Gmarket-0",
                                    "name":"Gmarket",
                                    "actionType":0,
                                    "desc":"Gmarket"
                                },
                                {
                                    "id":"DECATHLON-0",
                                    "name":"Decathlon",
                                    "actionType":0,
                                    "desc":"迪卡儂線上購物"
                                }
                            ]
                        },
                        {
                            "name":"超市",
                            "operator":1,
                            "descs":["全聯超市"],
                            "constraintType":7,
                            "supermarkets":[
                                {
                                    "id":"PxMart-0",
                                    "name":"Px Mart",
                                    "actionType":0,
                                    "desc":"於全聯福利中心消費, 待修"
                                }
                            ]  
                        },
                        {
                            "name":"影音",
                            "operator":1,
                            "descs":["Spotify", "Netflix", "CATCHPLAY+"],
                            "constraintType":7,
                            "streamings":[
                                {
                                    "id":"Spotify-0",
                                    "name":"Spotify",
                                    "actionType":0,
                                    "desc":"Spotify"
                                },
                                {
                                    "id":"Netflix-0",
                                    "name":"Netflix",
                                    "actionType":0,
                                    "desc":"Netflix"
                                },
                                {
                                    "id":"CATCHPLAY+-0",
                                    "name":"CATCHPLAY+",
                                    "actionType":0,
                                    "desc":"CATCHPLAY+"
                                }
                            ] 
                        }
                    ]
                }
            ]
        }
    }
]










[
    {
        "rewardID":"178e34d3-f2c2-4b76-60f3-cbfce5878edb",
        "name":"週六3%",
        "desc":"8大精選，週六3%",
        "start_date":1625097600,
        "end_date":1643587200,
        "update_date":1624428666,
        "constraintPayload":{
            "name":"",
            "operator":0,
            "descs":["使用行動支付在指定網購平台消費"],
            "constraintType":0,
            "constraintPayloads":[
                 {
                    "name":"最高回饋200元",
                    "operator":1,
                    "descs":["200元"],
                    "constraintType":4,
                    "bonusLimit":{
                        "id":"richard-gogo-2021",
                        "richart":"最高回饋200元",
                        "bonusType":0,
                        "atLeast":0,
                        "atMost":200
                    }
                 },
                 {
                    "name":"8大精選，週六3%",
                    "operator":0,
                    "descs":["使用行動支付在指定網購平台消費"],
                    "constraintType":0,
                    "constraintPayloads":[
                        {
                            "name":"網購",
                            "operator":1,
                            "descs":["蝦皮", "momo", "PChome"],
                            "constraintType":6,
                            "ecommerces":[
                                {
                                    "id":"Shopee-0",
                                    "name":"Shopee",
                                    "actionType":0,
                                    "desc":"蝦皮購物"
                                },
                                {
                                    "id":"MOMO-0",
                                    "name":"MOMO",
                                    "actionType":0,
                                    "desc":"MOMO購物"
                                },
                                {
                                    "id":"PCHOME-0",
                                    "name":"PCHOME",
                                    "actionType":0,
                                    "desc":"PCHOME"
                                }
                            ]
                        },
                        {
                            "name":"行動支付",
                            "operator":1,
                            "descs":["LINE PAY", "MyFamiPay", "Open錢包", "悠遊付", "台新Pay"],
                            "constraintType":5,
                            "mobilepays":[
                                {
                                    "id":"LINE-PAY-0",
                                    "name":"LINE PAY",
                                    "actionType":0,
                                    "desc":"連線支付"
                                },
                                {
                                    "id":"MyFamiPay-0",
                                    "name":"My FamiPay",
                                    "actionType":0,
                                    "desc":"全家超商支付"
                                },
                                {
                                    "id":"OpenWallet-0",
                                    "name":"Open Wallet",
                                    "actionType":0,
                                    "desc":"OPEN錢包"
                                },
                                {
                                    "id":"EasyWallet-0",
                                    "name":"Easy Wallet",
                                    "actionType":0,
                                    "desc":"Easy Wallet 悠遊付"
                                }
                            ]  
                        }
                    ]
                }
            ]


        }

    }
]
*/