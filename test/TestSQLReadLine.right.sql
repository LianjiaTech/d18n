insert into test_table(a, b, c, d)
values (1, '2"', "3'", 4), /* test " ; ' # --  ()s test*/(1, 2, 3, 4),
       -- fsdfsa # / *
       # dsadfas /*
       -- /*
       (1, 2, 3, 4),
       (1, 2, 3, 4),
       (1, 2, 3, 4);
replace
into test_table(a,b,c,d)
values
    (1,2,3,4)
    ,(2,3,4,5);