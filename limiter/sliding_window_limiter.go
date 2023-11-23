package limiter

import (
	"sync"
	"sync/atomic"
	"time"

	ratelimit_kit "github.com/ulovecode/ratelimit-kit"
)

var once sync.Once

// 参考TCP的滑动窗口协议
// 1. 优点：将固定时间段分块，时间比“计数器”复杂，适用于稍微精准的场景
// 2. 缺点：时间区间的精度越高,算法所需的空间容量就越大；实现稍微复杂，还是不能彻底解决“计数器”存在的边界问题
type slidingWindowLimiter struct {
	curRequests      int32         // 当前请求值
	durationRequests chan int32    // 请求队列
	accuracy         time.Duration // 时间窗口最小单元
	snippet          time.Duration // 时间窗口时间跨度
	curRequestsSum   int32         // 当前请求数之和
	allowRequests    int32         // 最大允许数量
}

func NewSlidingWindowLimiter(accuracy, snippet time.Duration, allowRequests int32) *slidingWindowLimiter {
	return &slidingWindowLimiter{
		durationRequests: make(chan int32, snippet/accuracy),
		accuracy:         accuracy,
		snippet:          snippet,
		allowRequests:    allowRequests,
	}
}

func (s *slidingWindowLimiter) Take() error {
	once.Do(func() { // 往前划动一个最小时间窗口
		go sliding(s)
		go calculate(s)
	})
	curRequest := atomic.LoadInt32(&s.curRequestsSum)
	if curRequest >= s.allowRequests {
		return ratelimit_kit.ErrExceededLimit
	}
	if !atomic.CompareAndSwapInt32(&s.curRequestsSum, curRequest, curRequest+1) { //cas
		return ratelimit_kit.ErrExceededLimit
	}
	atomic.AddInt32(&s.curRequests, 1)
	return nil
}

func sliding(s *slidingWindowLimiter) {
	for {
		select {
		case <-time.After(s.accuracy):
			s.durationRequests <- atomic.SwapInt32(&s.curRequests, 0)
		}
	}
}

func calculate(s *slidingWindowLimiter) {
	for {
		<-time.After(s.accuracy)
		if len(s.durationRequests) == cap(s.durationRequests) { // channel满了
			break
		}
	}
	for {
		<-time.After(s.accuracy)
		t := <-s.durationRequests
		if t != 0 {
			atomic.AddInt32(&s.curRequestsSum, -t)
		}
	}
}
