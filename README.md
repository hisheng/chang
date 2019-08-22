# chang

http://quotes.money.163.com/f10/hydb_601155.html#01g01

以对比为 核心的 分析体系


https://labs.zoe.im/Xueqiu-Spider/ 


一 性价比 对比


估值水平---成长比
pe  pb  p销   营收增长 利润增长   利润增长/pe

----这些都是 每天 都有值得，然后 会有一个历史 图

优先做这个，最重要的核心功能




1.自动采集 某只股票的 历史数据

1.1每天的收盘股价------
https://stock.xueqiu.com/v5/stock/chart/kline.json?symbol=SH601155&begin=1564744635504&period=day&type=before&count=-142&indicator=kline,pe,pb,ps,pcf,market_capital,agt,ggt,balance
这个接口有 pe的历史数据


1.2 每股收益是 每个季度会有一个值
然后 pe 可以根据最新的 四个季度计算出来-----
计算出pe
https://xueqiu.com/snowman/S/SH601155/detail#/ZYCWZB
这个接口，可以获取 营收 和 利润的增长----

把这两个数据 弄好，基本 第一个核心功能就ok 了

----这个需要 爬 html  https://github.com/PuerkitoBio/goquery 
goquery---下一步先走通这个 爬虫，试试数据，然后 数据有了，就是 
展示数据
----最后这些逻辑走通了，就是 
完善数据



//命令行 工具
https://studygolang.com/articles/7588



通过 问询函  学习分析财报
http://www.sse.com.cn/disclosure/credibility/supervision/inquiries/



股息
https://stock.xueqiu.com/v5/stock/quote.json?symbol=SH601155&extend=detail



查看代码行数
cloc ./
      11 text files.
      11 unique files.                              
       1 file ignored.

    github.com/AlDanial/cloc v 1.82  T=0.05 s (413.8 files/s, 40610.8 lines/s)
      -------------------------------------------------------------------------------
      Language                     files          blank        comment           code
      -------------------------------------------------------------------------------
      Go                              20            393             58           1764
      SQL                              1             45              6             55
      Markdown                         2             51              0             53
      TOML                             3              2              0             19
      -------------------------------------------------------------------------------
      SUM:                            26            491             64           1891
      -------------------------------------------------------------------------------








但是你这样后续可能有个问题，就是有可能只看到一时间的情况，看不到这个股这些年的规律
-----股价规律模型

高股息率，高增长， 这个模型怎么样

隐形冠军  这个也很好



选股的时候就要看好几年的业绩规律，而不是突然改变了再去算
业绩规律 就是 各个 季度 报表 利润占比 是吧







