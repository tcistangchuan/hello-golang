- SQL的执行顺序：

  from---where--group by---having---select---order by--limit
  
- sql慢常见原因

  ```
  1.索引未匹配或者没有建索引。
  2.join连表太多。
  3.锁或者死锁。
  4.没有limit限制，查询出来的数据量太大。
  5.返回了不必要的字段。
  ```
  
- sql语句优化

  ```
  1.创建并使用正确的索引
  2.只返回需要的字段
  3.合理使用批量插入
  4.每次查询返回合理的limit数据
  ```

- 大数据分页优化

  ```
  SELECT * FROM product WHERE ID >=(select id from product limit 866613, 1) limit 20
  ```

  



- type字段说明：连接类型，查找数据所用索引使用的情况。

  ```
  查询效率排序：
  all<index<range<ref<eq_ref<const<system
  
  all:全扫描
  index:索引树遍历扫描。（虽然where条件中没有用到索引，但select中的字段有索引，所以只要全表扫描索引即可，直接使用索引树查找数据）
  range:范围内索引树扫描
  ref:非主键非唯一索引等值扫描或者唯一索引的前缀扫描
  eq_ref:join连表使用了唯一索引或者主键索引字段
  const/system:唯一索引或者主键索引树等值扫描
  ```

- extra字段

  ```
  一、【Using where】 [不一定需要优化]
  	Extra为Using where说明，SQL使用了where条件过滤数据。结合type返回值查询，是否需要优化，优化方法就是在where里的字段添加索引
  	
  二、【Using index】
  Extra为Using index说明，SQL所需要返回的所有列数据均在一棵索引树上，而无需访问实际的行记录。这类SQL语句往往性能较好。
  
  三、【Using index condition】
  Extra为Using index condition说明，确实命中了索引，但不是所有的列数据都在索引树上，还需要访问实际的行记录。这类SQL语句性能也较高，但不如Using index。
  
  四、【Using filesort】[需要优化]
  Extra为Using filesort说明，得到所需结果集，需要对所有记录进行文件排序。
  这类SQL语句性能极差，需要进行优化。
  典型的，order by 上的字段没有命中索引，就会触发filesort，常见的优化方案是，在order by的列上添加索引，避免每次查询都全量排序。
  
  五、【Using temporary】[需要优化]
  Extra为Using temporary说明，需要建立临时表(temporary table)来暂存中间结果。
  这类SQL语句性能较低，往往也需要进行优化。
  
  原因：group by 和 order by 的字段不一样或者是字段没加索引
  
  典型的，group by和order by同时存在，且作用于不同的字段时，就会建立临时表，以便计算出最终的结果集。
  产生的条件：
  	如果GROUP BY 的列没有索引,产生临时表.
  　如果GROUP BY时,SELECT的列不止GROUP BY列一个,并且GROUP BY的列不是主键 ,产生临时表.
  　如果GROUP BY的列有索引,ORDER BY的列没索引.产生临时表.
  　如果GROUP BY的列和ORDER BY的列不一样,即使都有索引也会产生临时表.
  　如果GROUP BY或ORDER BY的列不是来自JOIN语句第一个表.会产生临时表.
  　如果DISTINCT 和 ORDER BY的列没有索引,产生临时表.
    如果group by 的列没有索引,必产生内部临时表,
    如果order by 与group by为不同列时,或多表联查时order by ,group by 包含的列不是第一张表的列,将会产生临时表

  ```
  
  
  
  
  
  
  
  
  
  <!--以下是bak部分（来源：https://github.com/tcistangchuan/LeetCode/blob/master/Rocket.md）-->
  
  ### explain命令概要
  
  1. id:select选择标识符
  2. select_type:表示查询的类型。
  3. table:输出结果集的表
  4. partitions:匹配的分区
  5. type:表示表的连接类型
  6. possible_keys:表示查询时，可能使用的索引
  7. key:表示实际使用的索引
  8. key_len:索引字段的长度
  9. ref:列与索引的比较
  10. rows:扫描出的行数(估算的行数)
  11. filtered:按表条件过滤的行百分比
  12. Extra:执行情况的描述和说明
  
  ### explain 中的 select_type（查询的类型）
  
  1. SIMPLE(简单SELECT，不使用UNION或子查询等)
  2. PRIMARY(子查询中最外层查询，查询中若包含任何复杂的子部分，最外层的select被标记为PRIMARY)
  3. UNION(UNION中的第二个或后面的SELECT语句)
  4. DEPENDENT UNION(UNION中的第二个或后面的SELECT语句，取决于外面的查询)
  5. UNION RESULT(UNION的结果，union语句中第二个select开始后面所有select)
  6. SUBQUERY(子查询中的第一个SELECT，结果不依赖于外部查询)
  7. DEPENDENT SUBQUERY(子查询中的第一个SELECT，依赖于外部查询)
  8. DERIVED(派生表的SELECT, FROM子句的子查询)
  9. UNCACHEABLE SUBQUERY(一个子查询的结果不能被缓存，必须重新评估外链接的第一行)
  
  ### explain 中的 type（表的连接类型）
  
  1. system：最快，主键或唯一索引查找常量值，只有一条记录，很少能出现
  2. const：PK或者unique上的等值查询
  3. eq_ref：PK或者unique上的join查询，等值匹配，对于前表的每一行(row)，后表只有一行命中
  4. ref：非唯一索引，等值匹配，可能有多行命中
  5. range：索引上的范围扫描，例如：between/in
  6. index：索引上的全集扫描，例如：InnoDB的count
  7. ALL：最慢，全表扫描(full table scan)
  
  ### explain 中的 Extra（执行情况的描述和说明）
  
  1. Using where:不用读取表中所有信息，仅通过索引就可以获取所需数据，这发生在对表的全部的请求列都是同一个索引的部分的时候，表示mysql服务器将在存储引擎检索行后再进行过滤
  2. Using temporary：表示MySQL需要使用临时表来存储结果集，常见于排序和分组查询，常见 group by ; order by
  3. Using filesort：当Query中包含 order by 操作，而且无法利用索引完成的排序操作称为“文件排序”
  4. Using join buffer：改值强调了在获取连接条件时没有使用索引，并且需要连接缓冲区来存储中间结果。如果出现了这个值，那应该注意，根据查询的具体情况可能需要添加索引来改进能。
  5. Impossible where：这个值强调了where语句会导致没有符合条件的行（通过收集统计信息不可能存在结果）。
  6. Select tables optimized away：这个值意味着仅通过使用索引，优化器可能仅从聚合函数结果中返回一行
  7. No tables used：Query语句中使用from dual 或不含任何from子句