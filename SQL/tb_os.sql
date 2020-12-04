/usr/local/firebird/bin/isql -user sysdba

DROP USER 'SYSDBA';
CREATE USER  'SYSDBA' PASSWORD 'masterkey';
ALTER USER 'SYSDBA' GRANT ADMIN ROLE;

connect 'CLIPP.FDB';

user 'SYSDBA' password 'masterkey';


CREATE TABLE 'TB_OS' (
   ID_OS int,
   ID_CLIENTE int,
   ID_STATUS int
);

SELECT * FROM RDB$RELATIONS;

/usr/local/firebird/bin/isql -user sysdba clipp_ara.fdb

create user sysdba password 'masterkey';
commit;

-- ARA
connect 'clipp_ara.fdb';
delete from TB_OS;
insert into TB_OS values (1001, 6001, 0);
insert into TB_OS values (1002, 6002, 1);
insert into TB_OS values (1003, 6003, 1);
select * from TB_OS;
commit;

-- CWB
connect 'clipp_cwb.fdb';
delete from TB_OS;
insert into TB_OS values (2001, 7001, 0);
insert into TB_OS values (2002, 7002, 1);
insert into TB_OS values (2003, 7003, 1);
select * from TB_OS;
commit;

-- LON
connect 'clipp_lon.fdb';
delete from TB_OS;
insert into TB_OS values (3001, 8001, 0);
insert into TB_OS values (3002, 8002, 1);
insert into TB_OS values (3003, 8003, 1);
insert into TB_OS values (9999, 7002, 1);
select * from TB_OS;
commit;

-- NAT
connect 'clipp_nat.fdb';
delete from TB_OS;
insert into TB_OS values (4001, 9001, 0);
insert into TB_OS values (4002, 9002, 1);
insert into TB_OS values (4003, 9003, 1);
select * from TB_OS;
commit;

-- REC
connect 'clipp_REC.fdb';
delete from TB_OS;
insert into TB_OS values (5001, 3001, 0);
insert into TB_OS values (5002, 3002, 1);
insert into TB_OS values (5003, 3003, 1);
select * from TB_OS;
commit;
