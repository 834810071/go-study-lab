package demoredlock

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

var ErrReleaseResource = errors.New("error release resource")
var ErrAcquireResource = errors.New("error acquire resource")

type Locker interface {
	Lock(ctx context.Context) error
	Unlock(ctx context.Context) error
}

type locker struct {
	redisClients []*redis.Client // redis实例
	expiration   time.Duration   // 锁会根据给定的TTL自动释放
	drift        time.Duration   // 对redis客户过期的标记
	quorum       int             // 将被用于计算最多的失败次数。如果客户端获取锁失败次数小于N/2+1，我们将会尝试解锁所有实例上的资源。N表示redis实例的数量
	name         string          // name是用来作为redis 的键使用的，每个锁应该有唯一的名称
	value        string          // value是一个随机的字符串，因此只有设置他的客户端才有权限对其进行解锁
}

var scriptLock = `
if redis.call("EXISTS", KEYS[1]) == 1 then
return 0
end
return redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
`

// Lock 如果key被成功设置并且设置成功时没有到过期时间，我们应该增加设置成功的计数。
// 如果总的成功计数小于定义的quorum， 我们会调用unlock方法释放所有的锁。
// 并且对于unlock方法的实现，我们只需简单地循环所有客户端，并对每一个客户端进行解锁就可以了。
func (l *locker) Lock(ctx context.Context) error {
	totalSuccess := 0

	for _, rc := range l.redisClients {
		start := time.Now()

		status, err := rc.Eval(ctx, scriptLock, []string{l.name}, l.value, l.expiration.Milliseconds()).Result()
		if err != nil {
			return err
		}

		ok := status == "OK"

		now := time.Now()
		isTimeValid := (l.expiration - (now.Sub(start) - l.drift)) > 0

		if ok && isTimeValid {
			totalSuccess++
		}
	}

	if totalSuccess < l.quorum {
		l.Unlock(ctx)
		return ErrAcquireResource
	}

	return nil
}

var scriptUnlock = `
if redis.call("GET", KEYS[1]) == ARGV[1] then
 return redis.call("DEL", KEYS[1])
else
 return 0
end
`

func (l *locker) Unlock(ctx context.Context) error {
	totalSuccess := 0

	for _, rc := range l.redisClients {

		status, err := rc.Eval(ctx, scriptUnlock, []string{l.name}, l.value).Result()
		if err != nil {
			return err
		}

		if status != int64(0) {
			totalSuccess++
		}
	}

	if totalSuccess < l.quorum {
		return ErrReleaseResource
	}

	return nil
}

type DLM struct {
	redisClients []*redis.Client
	quorum       int
	expiration   time.Duration
	drift        time.Duration
}

func NewDLM(redisClients []*redis.Client, expiration time.Duration, drift time.Duration) *DLM {
	return &DLM{
		redisClients: redisClients,
		expiration:   expiration,
		drift:        drift,
		quorum:       len(redisClients)/2 + 1,
	}
}

func (dlm *DLM) NewLocker(name string) Locker {
	return newLocker(dlm, name)
}

func newLocker(dlm *DLM, name string) Locker {
	return &locker{
		redisClients: dlm.redisClients,
		quorum:       dlm.quorum,
		name:         name,
		value:        generateRandomString(),
		expiration:   dlm.expiration,
		drift:        dlm.drift,
	}
}

func generateRandomString() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune,
		time.Now().Unix()%64)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
