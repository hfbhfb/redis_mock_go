package main

import (
	"errors"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

// newTestRedis returns a redis.Cmdable.
func newTestRedis() *redismock.ClientMock {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return redismock.NewNiceMock(client)
}

// This would be your production code.
func RedisIsAvailable(client redis.Cmdable) bool {
	return client.Ping().Err() == nil
}

// Test Redis is down.
func TestRedisCannotBePinged(t *testing.T) {
	r := newTestRedis()
	r.On("Ping").
		Return(redis.NewStatusResult("", errors.New("server not available")))

	assert.False(t, RedisIsAvailable(r))
}

func TestRedisLoopSet(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "117.44.31.104:15011",
	})
	client.Set("mm", "hello world!!", time.Minute)
}
