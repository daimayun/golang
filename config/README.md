## 配置文件

#### 一、INI类型的文件

##### 1、config.ini
```ini
[mysql]
username='root'
password='123456'

[redis]
host = '127.0.0.1'
port = 3306
debug = false

[mongodb]
user='admin'
password='admin'
```
##### 2、使用
```go
package main

import (
	"fmt"
	"github.com/daimayun/golang/config"
)

func main() {
	_ = config.NewIniConf("./conf", "config.ini")
	host := config.Conf.GetString("redis.host")
	port := config.Conf.GetInt64("redis.port")
	debug := config.Conf.GetBool("redis.debug")
	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(debug)
}
```

#### 二、JSON类型的文件

##### 1、config.json
```json
{
  "version": "2.0",
  "secret": "hello world",
  "debug": true,
  "host": {
    "origin": "https://www.baidu.com",
    "port": 8080
  }
}
```
##### 2、使用
```go
package main

import (
	"fmt"
	"github.com/daimayun/golang/config"
)

func main() {
	_ = config.NewJsonConf("./conf", "config.json")

	version := config.Conf.GetString("version")
	origin := config.Conf.GetString("host.origin")

	fmt.Println(version)
	fmt.Println(origin)

	// 读取到map中
	host := config.Conf.GetStringMapString("host")
	fmt.Println(host)
	fmt.Println(host["origin"])
	fmt.Println(host["port"])

	allSettings := config.Conf.AllSettings()
	fmt.Println(allSettings)
}
```

#### 三、YAML类型的文件

##### 1、config.yaml
```yaml
database:
  host: 127.0.0.1
  user: root
  dbname: test
  pwd: 123456
```
##### 2、使用
```go
package main

import (
	"fmt"
	"github.com/daimayun/golang/config"
)

func main() {
	_ = config.NewYamlConf("./conf", "config.yaml")

	host := config.Conf.GetString("database.host")
	fmt.Println("viper load yml: ", host)

	allSettings := config.Conf.AllSettings()
	fmt.Println(allSettings)
}
```
