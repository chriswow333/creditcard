

DROP TABLE bank;
create table bank (
    "id" VARCHAR(36) PRIMARY KEY,
	"name" VARCHAR(100),
	"desc" TEXT,
	start_date INTEGER,
	end_date INTEGER,
	update_date INTEGER
)


DROP TABLE card;
create table card (
    "id" VARCHAR(36) PRIMARY KEY,
    bank_id VARCHAR(36), 
	"name" varchar(100),
	"desc" TEXT,
	start_date INTEGER,
	end_date INTEGER,
	update_date INTEGER,
    FOREIGN KEY(bank_id) REFERENCES BANK("id")
)


DROP TABLE privilage;
create table privilage (
    "id" VARCHAR(36) PRIMARY KEY,
    card_id VARCHAR(36), 
	"name" varchar(100),
	"desc" TEXT,
	start_date INTEGER,
	end_date INTEGER,
	update_date INTEGER,
    FOREIGN KEY(card_id) REFERENCES CARD("id")
)

DROP TABLE "constraint";
create table "constraint" (
    "id" VARCHAR(36) PRIMARY KEY,
    privilage_id VARCHAR(36), 
	"name" varchar(100),
	"desc" TEXT,
    "operator" INTEGER,
	start_date INTEGER,
	end_date INTEGER,
	update_date INTEGER,
    FOREIGN KEY(privilage_id) REFERENCES PRIVILAGE("id")
)