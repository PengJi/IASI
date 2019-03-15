# IASI
Insects Awaken SQL Instructor

# scenes
1. 如何防止业务的一条SQL把数据库打挂(cpu打满、I/O打满)
2. 特定查询的优化。现在所有的查询都是走的统一的框架，生成一个AST，以此执行，但实际对于一些特定查询，譬如 select count(*)，完全可以将AST压扁，让其直接跟 engine 交互，得到数据，快速返回。
3. 自动处理SQL qps达到某一阈值，自动kill。

# service
1. 启发式建议
2. 索引建议
3. explain分析
4. sql格式化

# references 
[tidb parser](https://github.com/pingcap/parser)  
[soar](https://github.com/XiaoMi/soar)  

[SQL解析在美团的应用](https://tech.meituan.com/2018/05/20/sql-parser-used-in-mtdp.html)  
[MySQL源代码：从SQL语句到MySQL内部对象](http://www.orczhou.com/index.php/2012/11/mysql-innodb-source-code-optimization-1/)  
[Queryparser，一款开源 SQL 解析工具](https://www.infoq.cn/article/uber-opensource-queryparser)  
[关于SQL解析，为何编程语言解析器ANTLR更胜一筹？](https://dbaplus.cn/news-155-2261-1.html)

