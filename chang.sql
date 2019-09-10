SELECT symbol,close,pe,pb,ps,gather_day FROM chang.stock_chart where symbol = "SH601155" AND gather_day >= '2017-01-01' order BY pe ASC ;

/* 营收增幅，利润增幅，pe增幅，和市值增幅的关系 */
select * from (
SELECT lirunbiao.symbol,name,total_revenue ,(last_revenue / total_revenue) as  revenue_percent,(last_op /net_profit_atsopc) as op_percent,(last_pe /pe) as pe_percent,(last_market  / market_capital) as market_capital_percent ,pe ,last_pe FROM chang.lirunbiao
left join (SELECT symbol,total_revenue as last_revenue, net_profit_atsopc as last_op FROM chang.lirunbiao where gather_day = '2018-12-31') a
on a.symbol = chang.lirunbiao.symbol
left join (SELECT symbol,market_capital as last_market ,pe as last_pe from stock_chart where gather_day = '2019-08-26') b
on b.symbol = chang.lirunbiao.symbol
left join (SELECT symbol,market_capital,pe from stock_chart where gather_day = '2016-01-29') c
on c.symbol = chang.lirunbiao.symbol
left join symbol on a.symbol = symbol.symbol
where gather_day = '2015-12-31' and (last_market  / market_capital) < 40 and pe < 20 and pe > 4 and last_pe > 4) d
where market_capital_percent > 0 and op_percent > 2.5 and op_percent < 10 ;



SELECT * FROM chang.lirunbiao where gather_day = '2015-12-31';
SELECT symbol,market_capital from stock_chart where gather_day = '2016-01-29' and symbol ='SH601155';



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
              (select symbol,  min(pe) as pe from chang.stock_chart where pe > 1 and gather_day <= '2019-07-22' group by symbol) b on a.symbol = b.symbol
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

/*某天到现在的股价涨幅*/
select count('*'),percent from (
select symbol,start_price,start_day,end_price,end_day,floor((end_price - start_price)/start_price) as percent from (
select stock_chart.symbol,close as start_price, gather_day as start_day ,end_price,end_day from chang.stock_chart
left join (select symbol,close as end_price, gather_day as end_day from chang.stock_chart where gather_day = '2019-08-26') a on a.symbol = chang.stock_chart.symbol
where gather_day = '2010-01-06') c where end_price is not  null) d group by percent;


/*月份波动规律*/
select symbol,avg(pb) as "平均pb",DATE_FORMAT(gather_day,'%m') as "月份" from chang.stock_chart   where symbol = 'SZ000002' and  pe > 1  and  gather_day<= '2019-02-01' group by DATE_FORMAT(gather_day,'%m');

select symbol,avg(pe) as "平均pb",DATE_FORMAT(gather_day,'%Y') as "年份" from chang.stock_chart   where symbol = 'SH601155' and  pe > 1 group by DATE_FORMAT(gather_day,'%Y');

select a.symbol,pe,a.year,month,yearpe,(pe - yearpe) from (
select symbol,avg(pe) as pe,DATE_FORMAT(gather_day,'%Y') as year ,DATE_FORMAT(gather_day,'%m') as month from chang.stock_chart
where symbol = 'SH601155' and  pe > 1
group by DATE_FORMAT(gather_day,'%Y') ,DATE_FORMAT(gather_day,'%m')) b
left join(
    select symbol,avg(pe) as yearpe,DATE_FORMAT(gather_day,'%Y') as year from chang.stock_chart where symbol = 'SH601155' and  pe > 1 group by DATE_FORMAT(gather_day,'%Y')
) a on a.year = b.year and a.symbol = b.symbol
;

select * from symbol where symbol = 'SH600346';


explain select * from chang.stock_chart;

select COUNT(1) from chang.stock_chart;
select * from chang.stock_chart order by id desc ;



select * from symbol
left join stock_chart on symbol.symbol = stock_chart.symbol
where areacode = 320000 and stock_chart.gather_day = '2019-09-05'
order by market_capital desc ;




/*查看 模拟交易的 日期 到 今天的 股价 增长幅度对比 */
select * from (
select a.symbol,a.start_price,close,(close-start_price)/start_price as profit,
       a.start_pe ,pe,start_day ,basic_eps_percent
from chang.stock_chart right join
     (select symbol,start_price,pe as "start_pe",gather_day as "start_day" from chang.moni where chang.moni.gather_day <= '2016-12-15' and pe <= 20) a
     on a.symbol = chang.stock_chart.symbol
left join (SELECT symbol,basic_eps_percent FROM chang.lirunbiao where gather_day = '2018-12-31') b on b.symbol = a.symbol
where chang.stock_chart.gather_day = '2019-08-21' and pe > 4 order by profit desc ) c;



select (sum(p) -1.5)/12 from (
select a.symbol,a.start_price,close,(close-start_price)/start_price as p
from chang.stock_chart right join
    (select symbol,start_price from chang.moni where chang.moni.gather_day = '2019-07-02') a
on a.symbol = chang.stock_chart.symbol
where chang.stock_chart.gather_day = '2019-08-20') b;

SELECT * FROM chang.moni order by id  desc ;
SELECT * FROM chang.moni where symbol = 'SZ300721'  order by id  desc ;
SELECT count(*) FROM chang.moni;


SELECT * FROM stock_quote  order by id  asc ;
SELECT * FROM stock_quote where symbol = 'SH601155' order by id  desc ;


/*报表的周期 与 股价的 反应情况*/
select a.symbol,a.year,a.report_type, a.qop,b.op,(a.qop/b.op) from (
select lirunbiao.symbol,date_format(gather_day,'%Y') as year,report_type,sum(op) as qop from  lirunbiao
group by year,symbol,report_type) a
left join (select symbol,date_format(gather_day,'%Y')as year,op from  lirunbiao where report_type = 'Q4') b
on a.symbol = b.symbol and a.year = b.year
where  a.symbol = 'SH601155'
;


/*按利润 倒序*/

select symbol.symbol,name,floor(op/100000000) from lirunbiao
left join symbol on symbol.symbol = lirunbiao.symbol
where lirunbiao.gather_day = '2018-12-31'
order by op desc ;

