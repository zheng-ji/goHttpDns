/*
 * Author: zheng-ji.info
 */

package logic

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/hoisie/web"
)

// Resp Type
type Resp struct {
	Code     int    `json:"c"`
	TargetIP string `json:"targetip,omitempty"`
	Host     string `json:"host, omitempty"`
	Msg      string `json:"msg, omitempty"`
}

const (
	SUCC   = 0
	FAILED = -1
	HTTP   = "http://"
)

// PingHandler Func
func PingHandler(ctx *web.Context) string {
	ret := "ok"
	return ret
}

// ResolveHandler Func
func ResolveHandler(ctx *web.Context) string {

	url := ctx.Params["url"]
	targetIPstr, hostStr, err := getResultFromCache(url)
	if err == nil {
		resp := Resp{
			Code:     SUCC,
			TargetIP: targetIPstr,
			Host:     hostStr,
		}
		return resp.jsonString()
	}

	targetIP, host, err := DnsDecoder(url)
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
			TargetIP: *targetIP,
			Host:     *host,
		}
		cacheResp(url, *host, *targetIP)
		seelog.Infof("[ResolveHandler] host:%s targetIp:%s", *host, *targetIP)
		return resp.jsonString()
	}
}
