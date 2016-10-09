/*
 * Author: zheng-ji.info
 */

package logic

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"github.com/hoisie/web"
	"runtime"
)

var (
	redisPool *redis.Pool
)

// Initializer Func
func Initializer(conf string) bool {

	err := parseConfigFile(conf)
	if nil != err {
		fmt.Println("%v\n", err)
		return false
	}
	runtime.GOMAXPROCS(runtime.NumCPU())

	//seelog 配置
	logger, err := seelog.LoggerFromConfigAsFile(appConfig.Logconf)
	if nil != err {
		fmt.Println("%v\n", err)
		return false
	}
	seelog.UseLogger(logger)

	//redis 链接
	redisPool = NewRedisPool(appConfig.Redis.Host, appConfig.Redis.Db)
	return true
}

// Loop Func Register web interface
func Loop() {

	// 注册web接口
	web.Get("/ping", PingHandler)
	web.Get("/d", ResolveHandler)
	addr := fmt.Sprintf("%s:%s", appConfig.Listen, appConfig.Port)
	web.Run(addr)
}
