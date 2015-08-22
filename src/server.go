// Author: zheng-ji.info

package main

import (
	"flag"
	"httpDns/src/logic"
)

var (
	configFile = flag.String("c", "../etc/conf.yml", "配置文件路径，默认etc/conf.yml")
)

func main() {

	flag.Parse()

	if !logic.Initializer(*configFile) {
		return
	}
	logic.Loop()
}
