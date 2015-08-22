/*
 * Author: zheng-ji.info
 */

package logic

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/hoisie/web"
)

type Resp struct {
	Code     int    `json:"c"`
	TargetIp string `json:"targetip,omitempty"`
	Host     string `json:"host, omitempty"`
	Msg      string `json:"msg, omitempty"`
}

const (
	SUCC   = 0
	FAILED = -1
)

//Ping
func PingHandler(ctx *web.Context) string {
	ret := "ok"
	return ret
}

// 查询 DNS 真实IP
func ResolveHandler(ctx *web.Context) string {

	url := ctx.Params["url"]
	targetIp, host, err := DnsDecoder(url)
	if nil != err {
		resp := Resp{
			Code: FAILED,
			Msg:  fmt.Sprintf("%s", err),
		}
		seelog.Errorf("[ResolveHandler] error: %v", err)
		return resp.jsonString()
	} else {
		resp := Resp{
			Code:     SUCC,
			TargetIp: *targetIp,
			Host:     *host,
		}
		cacheResp(*host, *targetIp)
		seelog.Infof("[ResolveHandler] host:%s targetIp:%s", *host, *targetIp)
		fmt.Println("%s %s ", *targetIp, *host)
		return resp.jsonString()
	}
}
