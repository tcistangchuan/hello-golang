- type字段说明：连接类型，查找数据所用索引使用的情况。

  ```
  查询效率排序：
  all<index<range<ref<eq_ref<const<system
  
  all:全扫描
  index:索引树遍历扫描
  range:范围内索引树扫描
  ref:非主键非唯一索引等值扫描或者唯一索引的前缀扫描
  eq_ref:唯一索引或者主键索引树等值扫描
  const/system:这个是通过索引很快技能定位到一行匹配的数据
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
  典型的，order by 上的字符没有命中索引，就会触发filesort，常见的优化方案是，在order by的列上添加索引，避免每次查询都全量排序。
  
  五、【Using temporary】[需要优化]
  Extra为Using temporary说明，需要建立临时表(temporary table)来暂存中间结果。
  这类SQL语句性能较低，往往也需要进行优化。
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

  