package mqclient

import (
	"mserver/setting"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// var f = func(client mqtt.Client, msg mqtt.Message) {
// 	fmt.Printf("MSG: %s\n topic %s\n", msg.Payload(), msg.Topic())
// 	devID := msg.Topic()[strings.Index(msg.Topic(), "/"):]
// 	fmt.Println(devID)
// 	conn := scheduler.ID2ConnMap.Get(devID)
// 	conn.Write([]byte(msg.Topic()))
// }

// Client mqtt客户端
var Client mqtt.Client

func init() {
	// mqtt.DEBUG = log.New(os.Stdout, "", 0)
	// mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker(setting.Config.MqttAddr).SetClientID(setting.Config.MqttClientID)
	opts.SetKeepAlive(5 * time.Second)
	opts.SetDefaultPublishHandler(nil)
	// opts.SetPingTimeout(1 * time.Second)

	Client = mqtt.NewClient(opts)
	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// if token := Client.Subscribe(fmt.Sprintf("dt/#"), 0, f); token.Wait() && token.Error() != nil {
	// 	fmt.Println(token.Error())
	// }

	// time.Sleep(6 * time.Second)

	// if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
	// 	fmt.Println(token.Error())
	// 	os.Exit(1)
	// }

	// c.Disconnect(250)

	// time.Sleep(1 * time.Second)
}
