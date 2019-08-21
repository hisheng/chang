SELECT symbol,close,pe,pb,ps,gather_day FROM chang.stock_chart where symbol = "SH601155" AND gather_day >= '2017-01-01' order BY pe ASC ;

SELECT * FROM chang.lirunbiao;

/*查看江苏的企业*/
SELECT symbol,name,exchange,area_name FROM chang.symbol where areacode = 320000 ;

/*查看江苏的企业 今天的股价*/
SELECT symbol.symbol,name,exchange,area_name,stock_chart.close,stock_chart.pe,stock_chart.pb,stock_chart.ps ,stock_chart.market_capital FROM chang.symbol
left join chang.stock_chart on symbol.symbol = stock_chart.symbol
where (chang.symbol.areacode = 320000 or chang.symbol.areacode = 330000) and stock_chart.gather_day = "2019-08-20";


SELECT symbol.symbol,name,exchange,area_name,stock_chart.close,stock_chart.pe,stock_chart.pb,stock_chart.ps ,stock_chart.market_capital FROM chang.symbol
left join chang.stock_chart on symbol.symbol = stock_chart.symbol
where (chang.symbol.areacode = 320000 or chang.symbol.areacode = 330000) and stock_chart.gather_day = "2019-08-20";




SELECT symbol.symbol,name,exchange,area_name,stock_chart.close,stock_chart.pe,stock_chart.pb,stock_chart.ps ,stock_chart.market_capital FROM chang.symbol
left join chang.stock_chart on symbol.symbol = stock_chart.symbol
where stock_chart.gather_day = "2019-08-20";


SELECT count(*) FROM chang.symbol;

show databases ;

show tables ;

SELECT * FROM chang.symbol;

SHOW CREATE TABLE chang.lirunbiao;


/* 检查 2019-08-09 pe 创新低的 股票*/
select a.symbol,a.close,a.pe,a.pb,a.ps,a.gather_day from chang.stock_chart a left join
              (select symbol,  min(pe) as pe from chang.stock_chart where pe > 1 group by symbol) b on a.symbol = b.symbol
 where gather_day = '2019-07-22' and a.pe=b.pe;


select symbol,close,pe from stock_chart where symbol = "SH600681" and gather_day = '2019-07-22';
select symbol,close,pe from stock_chart where symbol = "SH600681" and gather_day = '2019-08-20';



/**/
select symbol,close,pe from stock_chart where symbol = "SH600998";
select symbol,close,pe from stock_chart where symbol = "SH603583";
select symbol,close,pe from stock_chart where symbol = "SH603629";
select symbol,close,pe from stock_chart where symbol = "SZ002757";
select symbol,close,pe from stock_chart where symbol = "SZ002859";
select symbol,close,pe from stock_chart where symbol = "SZ300752";
select symbol,close,pe from stock_chart where symbol = "SZ300767";
select symbol,close,pe from stock_chart where symbol = "SZ300778";





select symbol,close,pe,pb,ps,gather_day from chang.stock_chart   where symbol = 'SH601155' and  pe > 1 order by pe asc;

explain select * from chang.stock_chart;

select COUNT(1) from chang.stock_chart;
select * from chang.stock_chart;


