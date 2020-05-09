package scheduler

import (
	"mserver/models"
	"time"
)

//loadTimer 从数据库加载定时任务
func loadTimer() {
	ts := models.LoadAllTasks()
	for _, v := range ts {
		TaskCH <- &v
	}
}

// Sender 从通道读取任务周期发送并储存结束任务管道进StopMap
//如果管道有值会关闭管道并结束定时任务
func producer() {
	for {
		msg := <-TaskCH
		if msg == nil || msg.Frequency == 0 {
			continue
		}
		//删除旧任务s
		st := StopMap.Get(msg.DevID)
		if st != nil {
			st <- true
		}
		stop := taskGen(msg)
		StopMap.Set(msg.DevID, stop)
	}
}

func taskGen(msg *models.DeviceTask) (stop chan bool) {
	stop = make(chan bool)
	go func(m *models.DeviceTask, s chan bool) {
		for {
			select {
			case <-s:
				close(s)
				//删除
				StopMap.Del(m.DevID)
				return
			default:
				QueryCH <- msg
			}
			time.Sleep(time.Duration(m.Frequency) * time.Second)
		}
	}(msg, stop)
	return
}

// 生成链表，返回头部指针
// func getList(n int, node []models.Task) *TaskList {
// 	if n < len(node)-1 {
// 		list := &TaskList{
// 			PointID: node[n].PointID,
// 			Query:   node[n].MSG,
// 			Next:    getList(n+1, node),
// 		}
// 		return list
// 	}
// 	return &TaskList{
// 		PointID: node[n].PointID,
// 		Query:   node[n].MSG,
// 		Next:    nil,
// 	}
// }
