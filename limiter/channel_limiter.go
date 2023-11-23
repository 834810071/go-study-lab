package limiter

// ChannelLimiter 利用channel的缓冲设定，channel满了即阻塞
// 1. 优点：简单易实现，适合一次性限流
// 2. 缺点：阻塞无时间限制，无法自动解除限流
type ChannelLimiter struct {
	bufferChannel chan int
}

func NewChannelLimiter(limit int) *ChannelLimiter {
	return &ChannelLimiter{bufferChannel: make(chan int, limit)}
}

func (c *ChannelLimiter) Allow() bool {
	select {
	case c.bufferChannel <- 1:
		return true
	default:
		return false
	}
}

func (c *ChannelLimiter) Release() {
	<-c.bufferChannel
}
