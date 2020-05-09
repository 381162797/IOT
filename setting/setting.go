package setting

import (
	"encoding/json"
	"io/ioutil"
)

// Config 保存配置信息
var Config = &conf{}

// conf 配置结构体
type conf struct {
	DBURL        string
	DBType       string
	APIAddr      string
	TCPAddr      string
	JwtSecret    string
	MqttAddr     string
	MqttClientID string
}

func init() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, Config)
	if err != nil {
		panic(err)
	}
}
