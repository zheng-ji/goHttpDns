/*
 * Author: zheng-ji.info
 */

package logic

import (
	"encoding/json"
	"github.com/cihub/seelog"
)

func substr(s string, from, to int) string {
	bytes := []byte(s)
	return string(bytes[from:to])
}

func (resp *Resp) jsonString() string {
	b, _ := json.Marshal(resp)
	return string(b)
}

func cacheResp(host, targetIp string) {
	var err error
	conn := redisPool.Get()
	if conn == nil {
		return
	}
	defer conn.Close()

	_, err = conn.Do("SETEX", host, appConfig.Ttl, targetIp)
	if err != nil {
		seelog.Errorf("[Redis][SETEX] error: %v", err)
	}
}
