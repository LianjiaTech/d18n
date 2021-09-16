alter session set current_schema = system;

-- test chinese
CREATE TABLE "nvarchar2_demo" (
	description NVARCHAR2(50)
);

INSERT INTO "nvarchar2_demo" VALUES (N'中文');
INSERT INTO "nvarchar2_demo" VALUES ('Hello World!');

-- test raw data
CREATE TABLE "rawdata" (
	c1 RAW(11)
);

INSERT INTO "rawdata" VALUES ('1');
INSERT INTO "rawdata" VALUES ('AB');

exit
