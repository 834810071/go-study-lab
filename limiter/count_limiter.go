package limiter

import (
	"sync/atomic"
	"time"
)

// CountLimiter 在单位时间内进行计数,如果请求数大于设置的最大值,则进行拒绝；如果过了单位时间,则重新进行计数
// 1. 优点：简单易实现
// 2. 缺点：突发流量会出现毛刺现象，限流不准确
type CountLimiter struct {
	counter      int64
	limit        int64
	intervalNano int64
	unixNano     int64
}

func NewCountLimiter(interval time.Duration, limit int64) *CountLimiter {
	return &CountLimiter{
		counter:      0,
		limit:        limit,
		intervalNano: int64(interval),
		unixNano:     time.Now().UnixNano(),
	}
}

func (c *CountLimiter) Allow() bool {
	now := time.Now().UnixNano()
	if now-c.unixNano > c.intervalNano {
		atomic.StoreInt64(&c.counter, 0)
		atomic.StoreInt64(&c.unixNano, now)
		return true
	}
	atomic.AddInt64(&c.counter, 1)
	return c.counter <= c.limit
}
