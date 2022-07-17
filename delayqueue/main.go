package delayqueue

import (
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	redisCli := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	queue := NewQueue("example", redisCli, func(payload string) bool {
		// callback returns true to confirm successful consumption.
		// If callback returns false or not return within maxConsumeDuration, DelayQueue will re-deliver this message
		return true
	})
	// send delay message
	for i := 0; i < 10; i++ {
		err := queue.SendDelayMsg(strconv.Itoa(i), time.Hour, WithRetryCount(3))
		if err != nil {
			panic(err)
		}
	}
	// send schedule message
	for i := 0; i < 10; i++ {
		err := queue.SendScheduleMsg(strconv.Itoa(i), time.Now().Add(time.Hour))
		if err != nil {
			panic(err)
		}
	}
	// start consume
	done := queue.StartConsume()
	<-done
}
