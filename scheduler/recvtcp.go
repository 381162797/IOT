package scheduler

import (
	"fmt"
	"mserver/models"
	"net"
)

func receiveTCP(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(conn)
	}
}

// 试图读取设备ID
func handler(conn net.Conn) {
	data := make([]byte, 64)
	for {
		n, err := conn.Read(data)
		if err != nil {
			//断开连接
			return
		}
		//设备消息可能有16进制或者10进制
		connFilter(conn, data[:n])
		return
	}
}

func connFilter(conn net.Conn, data []byte) {
	nor := string(data)
	dHex := fmt.Sprintf("%x", data)
	if len(nor) == 8 && nor[:2] == "80" {
		//设备ID是否存在
		old := ID2ConnMap.Get(nor)
		if old != nil {
			conn.Write([]byte("0000"))
			conn.Close()
			return
		}
		//保存连接
		ID2ConnMap.Set(nor, conn)
		conn.Write([]byte("0001"))
		//加载定时任务
		task := models.GetDevSchedule(nor)
		TaskCH <- task
		return
	}
	if len(dHex) == 8 && dHex[:2] == "80" {
		old := ID2ConnMap.Get(dHex)
		if old != nil {
			conn.Write([]byte("0000"))
			conn.Close()
			return
		}
		//保存连接
		ID2ConnMap.Set(dHex, conn)
		conn.Write([]byte("0001"))
		//加载定时任务
		task := models.GetDevSchedule(dHex)
		TaskCH <- task
		return
	}
	conn.Write([]byte("0000"))
	conn.Close()
	return
}
