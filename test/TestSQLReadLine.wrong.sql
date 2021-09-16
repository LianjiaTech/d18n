-- insert into test_table
values (1, 2, 3, 4, 5, 6);
/* insert into test_table(a, 1b, c, d) */
values (1, 2, 3, 4),
       (1, 2, 3, 4),
       (1, 2, 3, 4),
       (1, 2, 3, 4);

#insert into test_table
values (1, 2),
       (3, 4, 5),
       (5, 6);
insert into test_table select 1,2,3,4,5;
select * from test_table;
update table test a=1,b=2 where id=10;