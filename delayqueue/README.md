# 流程转化
```shell
                                                       ——————————unack2Retry———————————————
                                                       ⬆️                                 ⬇️
Pending ——pending2Ready——> Ready --ready2Unack--> Unack -——unack2Retry——> Retry       Garbage
                                          ⬇️           ⬆️                     ⬇️
                                        Client          —————————————————————
                                          ⬇️
                                         del
                                           
```
# [用 Redis 做一个可靠的延迟队列](https://mp.weixin.qq.com/s/rZIIL1TtAgwliYI0GpVMMg)