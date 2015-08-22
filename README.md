## goHttpDns

A HttpDns Server Written by Go, In order to avoid Dns hijacking and cache resolve answer

一个用 Go 写的 HttpDns 服务, 为了抵抗Dns劫持污染，并带有缓存功能 。

### How To Compile

```
cd $GOPATH;
git clone http://github.com/zheng-ji/goHttpDns;
cd src;
make // httpDns will Generate in bin directory
```

### How To Configure

```
# redis connect config
redis:
  host: 127.0.0.1:6379
  db: 0

# seelog config 
log_config: ../etc/logger.xml

# ip & port & answer cache TTL
listen: 0.0.0.0
port: 9999
ttl: 100

# DnsServer lists
dnsservers:
    - 202.96.128.86
    - 202.96.128.166
    - 8.8.8.8
    - 8.8.4.4
```

### How To Run

```
zj@zheng-ji:$ ./httpDns --help
Usage of ./httpDns:
  -c="../etc/conf.yml": conf file，default is ../etc/conf.yml

./httpDns -c="your_conf_yaml_path"
```

You can also use `supervisor` to start your sever

### How To Use

```
$ curl http://127.0.0.1:9999/d?url=http://zheng-ji.info

Resp:
{
    "c":0,
    "targetip":"http://106.185.48.24",
    "host":"zheng-ji.info",
    "msg":""
}
```

### Dependece Third Part Lib

Thanks to:

* [launchpad/goyaml](https://launchpad.net/goyaml)
* [cihub/seelog](github.com/cihub/seelog)
* [miekg/dns](github.com/miekg/dns)
* [redisgo/redis](github.com/garyburd/redigo/redis")
* [hoisie/web](github.com/hoisie/web)

You need to go get the list above

----

MIT LICENSE 

Author [zheng-ji](http://zheng-ji.info)
