/*
 * Author: zheng-ji.info
 */

package logic

import (
	"github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"time"
)

func NewRedisPool(host string, db string) *redis.Pool {
	redisPool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 1 * time.Hour,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if nil != err {
				seelog.Errorf("[Redis] Dial erro: %v", err)
				return nil, err
			}
			if _, err := c.Do("SELECT", db); nil != err {
				c.Close()
				seelog.Errorf("[Redis] Select DB error: %v", err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return redisPool
}
