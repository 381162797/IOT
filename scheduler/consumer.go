package scheduler

import (
	"encoding/hex"
	"fmt"
	"log"
	"mserver/models"
	mqclient "mserver/mqtt/client"
	"mserver/tool"
	"net"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func consumer() {

	// 订阅MQTT
	f := func(client mqtt.Client, msg mqtt.Message) {
		topic := msg.Topic()
		idx := strings.Split(topic, "/")
		devID := idx[1]
		if devID != "" {
			conn := ID2ConnMap.Get(devID)
			if conn != nil {
				query, err := hex.DecodeString(string(msg.Payload()))
				if err != nil {
					log.Println(err)
					return
				}
				conn.Write(query)
			}
		}
	}

	if token := mqclient.Client.Subscribe(fmt.Sprintf("dt/#"), 0, f); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
	for {
		msg := <-QueryCH
		//启动协程发送消息
		fmt.Println(msg, "consumer 16")
		go msgHandler(msg)
	}
}

func msgHandler(msg *models.DeviceTask) {
	conn := ID2ConnMap.Get(msg.DevID)
	if conn == nil {
		//删除定时任务
		if stop := StopMap.Get(msg.DevID); stop != nil {
			stop <- true
		}
		return
	}
	ch := make(chan uint64)
	//开始读数据
	go readData(msg.DevID, ch)
	for _, v := range msg.Tasks {
		query, err := hex.DecodeString(v.Query)
		if err != nil {
			continue
		}
		ch <- v.PointID
		conn.Write(query)
	}
	close(ch)
	return
}

// 阻塞2s等待数据
func readData(devID string, point chan uint64) {
	conn := ID2ConnMap.Get(devID)
	b := make([]byte, 64)
	for v := range point {
		//设置2s读取时延
		conn.SetReadDeadline(time.Now().Add(time.Second * 2))
		n, err := conn.Read(b)
		if err != nil {
			//如果是Timeout进行下一次
			if e, ok := err.(net.Error); ok && e.Timeout() {
				continue
			}
			//可能是关闭了连接,关闭定时任务
			stop := StopMap.Get(devID)
			if stop != nil {
				stop <- true
			}
			//取完通道的消息
			continue
		}
		data := fmt.Sprintf("%x", b[:n])
		crc := tool.GetCRC16(b[:n-2])
		if crc == data[len(data)-4:] {
			//mqtt topic
			tpc := fmt.Sprintf("/test/%d", v)
			//如果PointID为0则为临时查询数据不保存
			if v == 0 {
				mqclient.Client.Publish(tpc, 0, false, data).Wait()
			}
			mqclient.Client.Publish(tpc, 0, false, data).Wait()
			models.SavePointData(v, devID, data)
		}
	}
	return
}
