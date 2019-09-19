# 核心概念
- `database`
- `retention policy`: measurement 维度数据的保留时间和集群中的副本数，默认 autogen 策略，数据不过期且副本数为1
- `series`: 一个系列由同一个 retention policy、同一个 measurement、同一个 tag set 组成
- `point`: 一个点由以下四个部分组成
- `measurement`
- `tag set` : `tag key` = `tag value`
- `field set` : `field key` = `field value`
- `timestamp`

# 词汇表
- `aggregation`: 聚合
- `batch`
- `continuous query (CQ)`: 自动定期运行查询，连续查询必须有 `GROUP BY time()`
- `database`
- `duration`: 数据存储时长
- `field`
- `field key`
- `field set`
- `field value`
- `function`
- `identifier`: token
- `InfluxDB line protocol`
- `measurement`
- `metastore`: 包含了系统内部信息
- `node`: 一个独立的 influxd 进程
- `now()`: 服务器本地时间 纳秒
- `point`
- `points per second`
- `query`
- `replication factor`
- `retention policy（RP）`
- `schema`
- `selector`
- `series`
- `series cardinality`
- `server`
- `shard`: 每个分片对应一个磁盘上的 TSM 文件，每个分片从属于唯一一个 shard group 
- `shard duration`
- `shard group`
- `subscription`
- `tag`
- `tag key`
- `tag set`
- `tag value`
- `timestamp`
- `transformation`
- `tsm (Time Structured Merge tree)`: 特定存储格式
- `user`
- `values per second`
- `wal (Write Ahead Log)`: 预写缓存

# 比较 InfluxDB 和 SQL 数据库
InfluxDB 就是被设计用于处理时间序列的数据。SQL 数据库虽然可以处理时间序列的数据，但并不是专门以此为目标。InfluxDB 可以更高效快速的存储大量时间序列数据并实时分析这些数据。

时间是绝对的主角，InfluxDB 不需要预先定义数据格式，支持 continuous queries 和 retention policies ，不支持跨 measurement 的 JOIN 查询，时间戳必须是 UNIX epoch（GMT）或者 RFC3339 格式字符串。

InfluxQL 和 SQL 非常相似。

InfluxDB 不是 CRUD：
- 更新 point 只需要插入相同 measurement、tag set、timestamp
- 你可以删除 series，但是不能基于 field 值去删除独立的 points，作为解决方法，你需要查询 field 值的时间，然后根据时间删除
- 你无法更新或重命名 tags ，只能创建新的并导入数据然后删除老的。
- 你无法通过 tag key 删除 tag

# InfluxDB 设计之道和权衡
https://docs.influxdata.com/influxdb/v1.7/concepts/insights_tradeoffs/


# 数据格式设计

## 好的设计
Tags 会被索引，fields 不会索引，tags 查询性能更高。
- 常用的查询数据存储为 tags
- 计划使用 `GROUP BY()` 的数据存储为 tags
- 计划使用 InfluxQL function 的数据存储为 fields
- 数据不只是 string 类型的存储为 fields ，tags 总是 string

避免使用 InfluxQL keywords 作为识别名称

## 不好的设计