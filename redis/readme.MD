1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
```bigquery
dz0400819@MacBook-Pro   redis-benchmark -h 127.0.0.1 -p 6379 -q -d 10
PING_INLINE: 121212.12 requests per second
PING_BULK: 123609.39 requests per second
17067:M 20 Jun 2021 12:24:24.786 * 10000 changes in 60 seconds. Saving...
17067:M 20 Jun 2021 12:24:24.786 * Background saving started by pid 17163
17163:C 20 Jun 2021 12:24:24.787 * DB saved on disk
17067:M 20 Jun 2021 12:24:24.886 * Background saving terminated with success
SET: 121212.12 requests per second
GET: 123001.23 requests per second
INCR: 121951.22 requests per second
LPUSH: 123762.38 requests per second
RPUSH: 125313.29 requests per second
LPOP: 120918.98 requests per second
RPOP: 125156.45 requests per second
SADD: 125944.58 requests per second
HSET: 123762.38 requests per second
SPOP: 123609.39 requests per second
ZADD: 121802.68 requests per second
ZPOPMIN: 126582.27 requests per second
LPUSH (needed to benchmark LRANGE): 125470.52 requests per second
LRANGE_100 (first 100 elements): 30012.00 requests per second
LRANGE_300 (first 300 elements): 12156.58 requests per second
LRANGE_500 (first 450 elements): 8306.34 requests per second
LRANGE_600 (first 600 elements): 6207.71 requests per second
MSET (10 keys): 96805.42 requests per second
```
2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

