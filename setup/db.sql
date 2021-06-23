

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

{
    "name":"台新銀行",
    "cardID":"96d7ba6a-227a-45f2-8b78-f92223353316",
    "name":"GoGo卡 不只網購，更懂生活",
    "desc":"@GoGo卡12大數位通路(網購/Pay/影音)最高6%，指定餐飲/健身/按摩等最高3%！, 綁Pay是標配！網購/健身/享樂也hen會！",
    "startDate":1625097600,
    "endDate":1643587200,
    "updateDate":1624412547,
    "operator":1
}

*/

/*
DROP TABLE base;
create table base (
    "id" VARCHAR(36) PRIMARY KEY,
	"bank_id" VARCHAR(36),
	"name" VARCHAR(100),
	"desc" TEXT,
	"target_from" VARCHAR(100),
	"target_to" VARCHAR(100),
	"base_type" INTEGER,
	"unit" VARCHAR(100),
	"action" INTEGER,
	FOREIGN KEY(bank_id) REFERENCES BANK("id")
);

INSERT INTO public.base("id", "bank_id", "name", "desc", "target_from", "target_to", "base_type", "unit", "action") 
values (gen_random_uuid(), '0ac94dfe-9604-472e-4738-fecdeac91ef1', '榜定Richart帳戶',  '以Richart帳戶自動扣繳@GoGo卡信用卡帳單', 
		'Richart', '', 2, 1, 2);

INSERT INTO public.base("id", "bank_id", "name", "desc", "target_from", "target_to", "base_type", "unit", "action") 
values (gen_random_uuid(), '0ac94dfe-9604-472e-4738-fecdeac91ef1', '消費滿5000元',  '', 
		'5000', '', 1, 2, 0);

*/


/**

{
    "privilageID":"ad82febb-8d0c-4ef5-5e29-3909cd2771dc",
    "name":"週六回饋",
    "desc":"",
    "startDate":1623204991,
    "endDate":1623204991,
    "limit":{
        "max":200,
        "min":0
    },
    "constraintBody":{
        "constraintPayloads":[
            {
                "operator":0,
                "base":[
                    {
                        "name":"週六回饋",
                        "desc":"僅只於週六提供",
                        "actDay":"Sat"
                    }
                ],
                "constraintPayloads":[
                    {
                        "operator":1,
                        "ecommerces":[
                            {
                                "id":"Shopee",
                                "name":"蝦皮購物",
                                "action":0,
                                "desc":"蝦皮購物"
                            },
                            {
                                "id":"MOMO",
                                "name":"MOMO購物",
                                "action":0,
                                "desc":"MOMO購物"
                            },
                            {
                                "id":"PChome",
                                "name":"PChome",
                                "action":0,
                                "desc":"PChome"
                            }
                        ]
                    },
                    {
                        "operator":1,
                        "mobilepays":[
                            {
                                "id":"LinePay",
                                "name":"LINE Pay",
                                "action":0,
                                "desc":"連線支付"
                            },
                            {
                                "id":"MyFamiPay",
                                "name":"My FamiPay",
                                "action":0,
                                "desc":"全家超商支付"
                            },
                            {
                                "id":"TaishinPay",
                                "name":"Taishin Pay",
                                "action":0,
                                "desc":"台新支付"
                            },
                            {
                                "id":"OpenWallet",
                                "name":"Open Wallet",
                                "action":0,
                                "desc":"OPEN錢包"
                            },
                            {
                                "id":"EasyWallet",
                                "name":"Easy Wallet",
                                "action":0,
                                "desc":"Easy Wallet 悠遊付"
                            }
                        ]
                    }
                ]
            }

        ]
        
    }
}

**/



/**

{
    "privilageID":"667256fb-4e6f-43d9-6a26-981a6d41ad68",
    "name":"精選回饋",
    "desc":"",
    "startDate":1623204991,
    "endDate":1623204991,
    "limit":{
        "max":200,
        "min":0
    },
    "constraintBody":{
        "constraintPayloads":[
            {
                "operator":0,
                "constraintPayloads":[
                    {
                        "operator":1,
                        "ecommerces":[
                            {
                                "id":"Shopee",
                                "name":"蝦皮購物",
                                "action":0,
                                "desc":"蝦皮購物"
                            },
                            {
                                "id":"MOMO",
                                "name":"MOMO購物",
                                "action":0,
                                "desc":"MOMO購物"
                            },
                            {
                                "id":"PChome",
                                "name":"PChome",
                                "action":0,
                                "desc":"PChome"
                            }
                        ]
                    },
                    {
                        "operator":1,
                        "mobilepays":[
                            {
                                "id":"LinePay",
                                "name":"LINE Pay",
                                "action":0,
                                "desc":"連線支付"
                            },
                            {
                                "id":"MyFamiPay",
                                "name":"My FamiPay",
                                "action":0,
                                "desc":"全家超商支付"
                            },
                            {
                                "id":"TaishinPay",
                                "name":"Taishin Pay",
                                "action":0,
                                "desc":"台新支付"
                            },
                            {
                                "id":"OpenWallet",
                                "name":"Open Wallet",
                                "action":0,
                                "desc":"OPEN錢包"
                            },
                            {
                                "id":"EasyWallet",
                                "name":"Easy Wallet",
                                "action":0,
                                "desc":"Easy Wallet 悠遊付"
                            }
                        ]
                    }
                ]
            }

        ]
        
    }
}

**/


/**

{
    "privilageID":"b02a75eb-08cb-482e-6755-cb72e2daaedd",
    "name":"任務回饋",
    "desc":"",
    "startDate":1623204991,
    "endDate":1623204991,
    "limit":{
        "max":200,
        "min":0
    },
    "constraintBody":{
        "constraintPayloads":[
            {
                "operator":0,
                "base":[
                    {
                        "id":1,
                        "name":"榜定Richart帳戶",
                        "desc":"以Richart帳戶自動扣繳@GoGo卡信用卡帳單",
                        "targetFrom":"Richart",
                        "targetTo":"",
                        "baseType":2,
                        "unitType":1,
                        "actionType":2
                    },
                    {
                        "id":2,
                        "name":"消費滿5000元",
                        "desc":"",
                        "targetFrom":"5000",
                        "targetTo":"",
                        "baseType":1,
                        "unitType":2,
                        "actionType":0
                    }
                ]
            }
        ]
    }
}

**/