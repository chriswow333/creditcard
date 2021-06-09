

DROP TABLE IF EXISTS BANK;
create table BANK (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"desc" TEXT,
	start_date INTEGER,
	end_date INTEGER,
	update_date INTEGER
)


DROP TABLE IF EXISTS CARD;
create table CARD (
    "id" VARCHAR(36) PRIMARY KEY,
    bank_id VARCHAR(36), 
	"name" varchar(100),
	"desc" TEXT,
	start_date INTEGER,
	end_date INTEGER,
	update_date INTEGER,
    FOREIGN KEY(bank_id) REFERENCES BANK("id")
)


DROP TABLE IF EXISTS PRIVILAGE;
create table PRIVILAGE (
    "id" VARCHAR(36) PRIMARY KEY,
    card_id VARCHAR(36), 
	"name" varchar(100),
	"desc" TEXT,
	start_date INTEGER,
	end_date INTEGER,
	update_date INTEGER,
    FOREIGN KEY(card_id) REFERENCES CARD("id")
)

DROP TABLE IF EXISTS "CONSTRAINT";
create table "CONSTRAINT" (
    "id" VARCHAR(36) PRIMARY KEY,
    privilage_id VARCHAR(36), 
	"name" varchar(100),
	"desc" TEXT,
    "operator" INTEGER,
	start_date INTEGER,
	end_date INTEGER,
	update_date INTEGER,
    limit_mx INTEGER,
    limit_mn INTEGER,
    FOREIGN KEY(privilage_id) REFERENCES PRIVILAGE("id")
)