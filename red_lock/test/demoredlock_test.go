package test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-unit-test/red_lock/demoredlock"
	"log"
	"testing"
	"time"
)

func TestRedLock(t *testing.T) {
	rc1 := redis.NewClient(&redis.Options{Addr: "0.0.0.0:7001"})
	rc2 := redis.NewClient(&redis.Options{Addr: "0.0.0.0:7002"})
	rc3 := redis.NewClient(&redis.Options{Addr: "0.0.0.0:7003"})

	dlm := demoredlock.NewDLM([]*redis.Client{rc1, rc2, rc3}, 10*time.Second, 2*time.Second)
	ctx := context.Background()
	locker := dlm.NewLocker("this-is-a-key-002")

	if err := locker.Lock(ctx); err != nil {
		log.Fatal("[main] Failed when locking, err:", err)
	}

	// Perform operation.
	someOperation()

	if err := locker.Unlock(ctx); err != nil {
		log.Fatal("[main] Failed when unlocking, err:", err)
	}

	log.Println("[main] Done")
}

func TestRedLockUnlock(t *testing.T) {
	log.SetFlags(log.Ltime)
	rc1 := redis.NewClient(&redis.Options{Addr: "0.0.0.0:7001"})
	rc2 := redis.NewClient(&redis.Options{Addr: "0.0.0.0:7002"})
	rc3 := redis.NewClient(&redis.Options{Addr: "0.0.0.0:7003"})

	dlm := demoredlock.NewDLM([]*redis.Client{rc1, rc2, rc3}, 10*time.Second, 2*time.Second)

	ctx := context.Background()
	locker := dlm.NewLocker("this-is-a-key-002")

	if err := locker.Lock(ctx); err != nil {
		log.Fatal("[main] Failed when locking, err:", err)
	}

	// Perform operation.
	someOperation()

	// Don't unlock

	log.Println("[main] Done")
}

func someOperation() {
	log.Println("[someOperation] Process has been started")
	time.Sleep(1 * time.Second)
	log.Println("[someOperation] Process has been finished")
}
