drop schema tollerdata;
create schema tollerdata;
use tollerdata;

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

CREATE TABLE sessions (
    token CHAR(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);


CREATE TABLE codetables (
    id VARCHAR(10) NOT NULL PRIMARY KEY,
    category VARCHAR(128) NOT NULL,
    code VARCHAR(10) NOT NULL,
    display VARCHAR(128) NOT NULL
);

create table litters (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    regnum VARCHAR(10) NOT NULL
);

CREATE table addresses (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    address1 VARCHAR(128),
    address2 VARCHAR(128),
    city VARCHAR(128),
    state VARCHAR(10),
    zip VARCHAR(10),
    FOREIGN KEY (state) REFERENCES codetables(id)
);

CREATE TABLE people (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    firstname VARCHAR(128),
    middlename VARCHAR(128),
    lastname VARCHAR(128),
    address INTEGER,
    email VARCHAR(128),
    FOREIGN KEY (address) REFERENCES addresses(id)
);

CREATE TABLE dogs (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    regnum VARCHAR(14) NOT NULL UNIQUE,
    nsdtrcregnum VARCHAR(14),
    sequencenum VARCHAR(14),
    litterid INT,
    name VARCHAR(255) NOT NULL,
    callname VARCHAR(255),
    whelpdate DATETIME NOT NULL,
    sex VARCHAR(10),
    nba BIT,
    alive BOOL DEFAULT true,
    intact BOOL DEFAULT true,
    sire INT,
    dam INT,
    FOREIGN KEY (litterid) REFERENCES litters(id),
    FOREIGN KEY (sex) REFERENCES codetables(id),
    FOREIGN KEY (sire) REFERENCES dogs(id),
    FOREIGN KEY (dam) REFERENCES dogs(id)
);

CREATE TABLE owners (
    personid integer,
    dogid integer,
    FOREIGN key (personid) references people(id),
    FOREIGN key (dogid) references dogs(id)
);

CREATE TABLE breeders (
    personid integer,
    dogid integer,
    FOREIGN key (personid) references people(id),
    FOREIGN key (dogid) references dogs(id)
);

Create TABLE health (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    dogid INT,
    healthtype VARCHAR(10),
    certid VARCHAR(14),
    FOREIGN key (dogid) references dogs(id),
    FOREIGN KEY (healthtype) REFERENCES codetables(id)
);

Create TABLE colors (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    dogid INT,
    colortype VARCHAR(10),
    certid VARCHAR(14),
    FOREIGN KEY (dogid) REFERENCES dogs(id),
    FOREIGN KEY (colortype) REFERENCES codetables(id)
);

insert into codetables (id, category, code, display) values ("SD", "sex", "D", "Dog");
insert into codetables (id, category, code, display) values ("SB", "sex", "B", "Bitch");

insert into codetables (id, category, code, display) values ("CB", "color", "B", "Buff");
insert into codetables (id, category, code, display) values ("CR", "color", "R", "Red");
insert into codetables (id, category, code, display) values ("CDR", "color", "DR", "Dark Red");
insert into codetables (id, category, code, display) values ("CLR", "color", "LR", "Light Red");
insert into codetables (id, category, code, display) values ("CG", "color", "G", "Gold");
insert into codetables (id, category, code, display) values ("CW", "color", "W", "White");
insert into codetables (id, category, code, display) values ("CWM", "color", "WM", "White Markings");
insert into codetables (id, category, code, display) values ("CF", "color", "F", "Fawn");
insert into codetables (id, category, code, display) values ("CS", "color", "S", "Straw");
insert into codetables (id, category, code, display) values ("CO", "color", "O", "Orange");

insert into codetables (id, category, code, display) values ("HHIP", "health", "HIP", "Hip");
insert into codetables (id, category, code, display) values ("HHRT", "health", "HRT", "Heart");
insert into codetables (id, category, code, display) values ("HEYE", "health", "EYE", "Eye");
insert into codetables (id, category, code, display) values ("HELB", "health", "ELB", "Elbow");


insert into codetables (id, category, code, display) values ("SAK", "state", "AK", "Alaska");
insert into codetables (id, category, code, display) values ("SAL", "state", "AL", "Alabama");
insert into codetables (id, category, code, display) values ("SAR", "state", "AR", "Arkansas");
insert into codetables (id, category, code, display) values ("SAZ", "state", "AZ", "Arizona");
insert into codetables (id, category, code, display) values ("SCA", "state", "CA", "California");
insert into codetables (id, category, code, display) values ("SCO", "state", "CO", "Colorado");
insert into codetables (id, category, code, display) values ("SCT", "state", "CT", "Connecticut");
insert into codetables (id, category, code, display) values ("SDE", "state", "DE", "Delaware");
insert into codetables (id, category, code, display) values ("SDC", "state", "DC", "District of Columbia");
insert into codetables (id, category, code, display) values ("SFL", "state", "FL", "Florida");
insert into codetables (id, category, code, display) values ("SGA", "state", "GA", "Georgia");
insert into codetables (id, category, code, display) values ("SHI", "state", "HI", "Hawaii");
insert into codetables (id, category, code, display) values ("SID", "state", "ID", "Idaho");
insert into codetables (id, category, code, display) values ("SIL", "state", "IL", "Illinois");
insert into codetables (id, category, code, display) values ("SIN", "state", "IN", "Indiana");
insert into codetables (id, category, code, display) values ("SIA", "state", "IA", "Iowa");
insert into codetables (id, category, code, display) values ("SKS", "state", "KS", "Kansas");
insert into codetables (id, category, code, display) values ("SKY", "state", "KY", "Kentucky");
insert into codetables (id, category, code, display) values ("SME", "state", "ME", "Maine");
insert into codetables (id, category, code, display) values ("SMD", "state", "MD", "Maryland");
insert into codetables (id, category, code, display) values ("SMA", "state", "MA", "Massachusetts");
insert into codetables (id, category, code, display) values ("SMI", "state", "MI", "Michigan");
insert into codetables (id, category, code, display) values ("SMN", "state", "MN", "Minnesota");
insert into codetables (id, category, code, display) values ("SMS", "state", "MS", "Mississippi");
insert into codetables (id, category, code, display) values ("SMO", "state", "MO", "Missouri");
insert into codetables (id, category, code, display) values ("SMT", "state", "MT", "Montana");
insert into codetables (id, category, code, display) values ("SNE", "state", "NE", "Nebraska");
insert into codetables (id, category, code, display) values ("SNV", "state", "NV", "Nevada");
insert into codetables (id, category, code, display) values ("SNH", "state", "NH", "New Hampshire");
insert into codetables (id, category, code, display) values ("SNJ", "state", "NJ", "New Jersey");
insert into codetables (id, category, code, display) values ("SNM", "state", "NM", "New Mexico");
insert into codetables (id, category, code, display) values ("SNY", "state", "NY", "New York");
insert into codetables (id, category, code, display) values ("SNC", "state", "NC", "North Carolina");
insert into codetables (id, category, code, display) values ("SND", "state", "ND", "North Dakota");
insert into codetables (id, category, code, display) values ("SOH", "state", "OH", "Ohio");
insert into codetables (id, category, code, display) values ("SOK", "state", "OK", "Oklahoma");
insert into codetables (id, category, code, display) values ("SOR", "state", "OR", "Oregon");
insert into codetables (id, category, code, display) values ("SPA", "state", "PA", "Pennsylvania");
insert into codetables (id, category, code, display) values ("SRI", "state", "RI", "Rhode Island");
insert into codetables (id, category, code, display) values ("SSC", "state", "SC", "South Carolina");
insert into codetables (id, category, code, display) values ("STN", "state", "TN", "Tennesee");
insert into codetables (id, category, code, display) values ("STX", "state", "TX", "Texas");
insert into codetables (id, category, code, display) values ("SUT", "state", "UT", "Utah");
insert into codetables (id, category, code, display) values ("SVT", "state", "VT", "Vermont");
insert into codetables (id, category, code, display) values ("SVA", "state", "VA", "Virginia");
insert into codetables (id, category, code, display) values ("SWA", "state", "WA", "Washington");
insert into codetables (id, category, code, display) values ("SWV", "state", "WV", "West Virginia");
insert into codetables (id, category, code, display) values ("SWI", "state", "WI", "Wisconsin");
insert into codetables (id, category, code, display) values ("SWY", "state", "WY", "Wyoming");
insert into codetables (id, category, code, display) values ("SUN", "state", "UN", "Unknown");

insert into codetables (id, category, code, display) values ("CAALB", "state", "ALB", "Alberta");
insert into codetables (id, category, code, display) values ("CABC", "state", "BC", "British Columbia");
insert into codetables (id, category, code, display) values ("CANS", "state", "NS", "Nova Scotia");
insert into codetables (id, category, code, display) values ("CAQUE", "state", "QUE", "Quebec");
insert into codetables (id, category, code, display) values ("CAYT", "state", "YT", "Yukon Territory");

insert into users(email, name, hashed_password, created) values ('louis.garwood@gmail.com', "Lou Garwood", "$2a$12$0joxuXNvL5Q02IBH9SZu/OH6kAS7M8PkwCo.MjK.EncG8bYHsb/oW", Now());