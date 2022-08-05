alter session set current_schema = system;

-- test chinese
CREATE TABLE test_language (
	description NVARCHAR2(50)
);

INSERT INTO test_language VALUES (N'中文');
INSERT INTO test_language VALUES ('Hello World!');

-- test raw data
CREATE TABLE test_raw (
	c1 RAW(11)
);

INSERT INTO test_raw VALUES ('1');
INSERT INTO test_raw VALUES ('AB');

-- test timestamp
create table test_ts(
  id number primary key,
  cname varchar2(32),
  balance number,
  amount  number(32,3),
  change_date date,
  create_time timestamp,
  update_time timestamp
);

insert into test_ts values (4, 'shanghai', 2427519.35, 4.520, to_date('2022-08-03', 'yyyy-mm-dd'), sysdate, current_timestamp);

exit
