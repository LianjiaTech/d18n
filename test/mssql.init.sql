-- test raw data
CREATE TABLE rawdata (
  c1 VARBINARY(11)
);
INSERT INTO rawdata VALUES (0x1);
INSERT INTO rawdata VALUES (CAST('AB' AS VARBINARY));
go
