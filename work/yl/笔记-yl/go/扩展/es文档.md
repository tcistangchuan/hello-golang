- 安装参考文档：

```
https://blog.csdn.net/Vi_NSN/article/details/81546303
```

- 关于ElasticSearch和MySQL中的概念对应关系:

```
https://blog.csdn.net/qq_37230634/article/details/103583537
```

- 操作：

```
https://www.ruanyifeng.com/blog/2017/08/elasticsearch.html
https://juejin.im/post/6844903860457177101
```

- Elasticsearch 查询技巧( yl 文档)

```
ES在6.2之后的版本中，X-Pack新增SQL查询的插件。 由于ES原生查询的语法复杂且难于记忆， 所以通过SQL查询的方式可以极大的降低学习成本，提高工作效率。

1 . 将SQL翻译成ES原生查询语法

POST /_sql/translate

curl -XPOST http://localhost:9200/_sql/translate -H 'Content-Type:application/json' -d'
{
"query": "select items from test1 limit 3"
}'

2. 直接通过SQL查询
POST  /_sql?format=json

curl -XPOST http://localhost:9200/_sql?format=json -H 'Content-Type:application/json' -d'
{
"query": "select items from test1 limit 3"
}'
```

```
https://github.com/olivere/elastic 
https://github.com/alibaba/canal
```

