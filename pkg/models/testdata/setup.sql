CREATE TABLE codetables (
    id VARCHAR(10) NOT NULL PRIMARY KEY,
    category VARCHAR(128) NOT NULL,
    code VARCHAR(10) NOT NULL,
    display VARCHAR(128) NOT NULL
);

insert into codetables (id, category, code, display) values ("SD", "sex", "D", "Dog");
insert into codetables (id, category, code, display) values ("SB", "sex", "B", "Bitch");

create table litters (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    regnum VARCHAR(10) NOT NULL UNIQUE
);

insert into litters (regnum) values ("no-litter");

CREATE TABLE dogs (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    regnum VARCHAR(24) NOT NULL,
    nsdtrcregnum VARCHAR(24) NOT NULL,
    sequencenum VARCHAR(24) NOT NULL,
    litterid INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    callname VARCHAR(255)  NOT NULL,
    whelpdate DATETIME NOT NULL,
    sex VARCHAR(10) NOT NULL,
    nba BOOL DEFAULT true,
    alive BOOL DEFAULT true,
    intact BOOL DEFAULT true,
    sire INT,
    dam INT,
    FOREIGN KEY (litterid) REFERENCES litters(id),
    FOREIGN KEY (sex) REFERENCES codetables(id),
    FOREIGN KEY (sire) REFERENCES dogs(id),
    FOREIGN KEY (dam) REFERENCES dogs(id)
);


insert into dogs (regnum, nsdtrcregnum, sequencenum, litterid, name, callname, sex, whelpdate) values ("DOG1", "", "", 1, "SIRE", "", "SD", NOW());

CREATE TABLE IF NOT EXISTS `tollers` (
`A_REGISTRA` VARCHAR(14),
`A_NSDTRC_R` VARCHAR(14),
`A_SEQ` VARCHAR(6),
`A_DOGNAME` VARCHAR(70),
`A_TITLE_NA` VARCHAR(255),
`A_LITTERRE` VARCHAR(10),
`A_SEX` VARCHAR(5),
`A_COLOR` VARCHAR(30),
`A_PRA` VARCHAR(4),
`A_HIPCLEAR` VARCHAR(28),
`A_EYECLEAR` VARCHAR(18),
`A_HEART_CL` VARCHAR(18),
`A_ELBOW_CL` VARCHAR(14),
`A_WHELPDAT` DATE,
`A_NBA` VARCHAR(4),
`A_ALIVE` VARCHAR(6),
`A_OWNER` VARCHAR(70),
`A_ADDRESS1` VARCHAR(51),
`A_INTACT` VARCHAR(6),
`A_CITY` VARCHAR(19),
`A_STATE` VARCHAR(9),
`A_ZIP` VARCHAR(15),
`A_COUNTRY` VARCHAR(14),
`A_SIRE_REG` VARCHAR(16),
`A_SIRENAME` VARCHAR(80),
`A_DAM_REGN` VARCHAR(16),
`A_DAMNAME` VARCHAR(80),
`A_BREEDER` VARCHAR(60),
`A_BREEDERA` VARCHAR(51),
`A_BREEDERC` VARCHAR(19),
`A_BREEDERS` VARCHAR(9),
`A_BREEDERZ` VARCHAR(15),
`A_BREEDER0` VARCHAR(14),
`A_ROM_ROMX` VARCHAR(2),
`A_CALLNAME` VARCHAR(10),
`A_EMAIL` VARCHAR(20)
);

INSERT INTO tollers (`A_REGISTRA`,`A_NSDTRC_R`,`A_SEQ`,`A_DOGNAME`,`A_TITLE_NA`,`A_LITTERRE`,`A_SEX`,`A_COLOR`,
    `A_PRA`,`A_HIPCLEAR`,`A_EYECLEAR`,`A_HEART_CL`,`A_ELBOW_CL`,`A_WHELPDAT`,`A_NBA`,`A_ALIVE`,
    `A_OWNER`,`A_ADDRESS1`,`A_INTACT`,`A_CITY`,`A_STATE`,`A_ZIP`,`A_COUNTRY`,`A_SIRE_REG`,`A_SIRENAME`,
    `A_DAM_REGN`,`A_DAMNAME`,`A_BREEDER`,`A_BREEDERA`,`A_BREEDERC`,`A_BREEDERS`,`A_BREEDERZ`,`A_BREEDER0`,
    `A_ROM_ROMX`,`A_CALLNAME`,`A_EMAIL`) 
    VALUES ("regnum", "nsdtrcregnum","seqnum","name","titlename","litterrn","D","color","pra",
    "hipclear","eyeclear","heartclear","elbowclear",NOW(),"nba","Y","owner","address1","Y","city",
    "state","zip", "country","DOG1","sirename","DOG1","damname","breedername","breederaddress",
    "breedercity","brst","breederzip","breedercountry","","callname","");
