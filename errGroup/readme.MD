调用127.0.0.1:8080/api/v1/stop
先调用g2中server.shutdown()停止g1，再调用g3中errGroupCtx.Done()停止g3，最后停止main
```
2021/05/06 23:37:06 g2 server stop,errGroup exit
[GIN] 2021/05/06 - 23:37:06 | 200 |       3.017µs |       127.0.0.1 | GET      "/api/v1/stop"
2021/05/06 23:37:06 g2 server shutdown
2021/05/06 23:37:06 g3 errGroupCtx.Done, errGroup exit, errGroupCtx.Err():context canceled
main errgroup exiting: http: Server closed

```
直接kill进程
g3中先返回err
g2中errGroupCtx.Done()不再阻塞，调用server.shutdown()停止g1
最后停止mian
```
2021/05/06 23:40:57 g3 get os signal: interrupt
2021/05/06 23:40:57 g2 errGroupCtx.Done,errGroup exit
2021/05/06 23:40:57 g2 server shutdown
main errgroup exiting: g3 get os signal: interrupt

```