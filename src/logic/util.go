/*
 * Author: zheng-ji.info
 */

package logic

import (
	"encoding/json"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
)

func substr(s string, from, to int) string {
	bytes := []byte(s)
	return string(bytes[from:to])
}

func (resp *Resp) jsonString() string {
	b, _ := json.Marshal(resp)
	return string(b)
}

func cacheResp(url, host, targetIP string) {
	var err error
	conn := redisPool.Get()
	if conn == nil {
		return
	}
	defer conn.Close()

	var key string

	key = fmt.Sprintf("%s_host", url)
	_, err = conn.Do("SETEX", key, appConfig.TTL, host)
	if err != nil {
		seelog.Errorf("[Redis][SETEX] error: %v", err)
	}

	key = fmt.Sprintf("%s_ip", url)
	_, err = conn.Do("SETEX", key, appConfig.TTL, targetIP)
	if err != nil {
		seelog.Errorf("[Redis][SETEX] error: %v", err)
	}
}

func getResultFromCache(url string) (targetIP string, host string, err error) {
	conn := redisPool.Get()
	if conn == nil {
		return
	}
	defer conn.Close()

	var key string

	key = fmt.Sprintf("%s_host", url)
	host, err = redis.String(conn.Do("GET", key))
	key = fmt.Sprintf("%s_ip", url)
	targetIP, err = redis.String(conn.Do("GET", key))

	if err != nil {
		seelog.Errorf("[Redis][GET] error: %v", err)
	}
	return
}
