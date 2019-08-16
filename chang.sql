SELECT symbol,close,pe,pb,ps,gather_day FROM chang.stock_chart where symbol = "SH601155" AND gather_day >= '2017-01-01' order BY pe ASC ;

SELECT * FROM chang.lirunbiao;

SELECT * FROM chang.symbol where areacode = 320000 ;

SELECT count(*) FROM chang.symbol;

show databases ;

show tables ;

SELECT * FROM chang.symbol;

SHOW CREATE TABLE chang.lirunbiao;


/* 检查 2019-08-09 创新低的 股票*/
select a.symbol,a.close,a.pe,a.pb,a.ps,a.gather_day from chang.stock_chart a left join
              (select symbol,  min(pe) as pe from chang.stock_chart where pe > 1 group by symbol) b on a.symbol = b.symbol
 where gather_day = '2019-08-09' and a.pe=b.pe;


select symbol,close,pe,pb,ps,gather_day from chang.stock_chart   where symbol = 'SH601155' order by pe asc
