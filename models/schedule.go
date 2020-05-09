package models

import (
	"mserver/def"
)

func AddSchedule(dt *DeviceTask) (code int) {
	old := &DeviceTask{}
	db.First(old, DeviceTask{DevID: dt.DevID})
	if old.ID != 0 {
		return def.TASK_EXIST
	}
	db.Create(dt)
	return def.SUCCESS
}

func DeleteSchedule(devID string) (code int) {
	old := &DeviceTask{DevID: devID}
	db.First(old, old)
	if old.ID == 0 {
		return def.TASK_NOT_EXIST
	}
	//删除Tasks
	db.Delete(Task{}, Task{DeviceTaskID: old.ID})
	//删除DeviceTask
	db.Delete(old, old)
	return def.SUCCESS
}

func UpdateSchedule(dt *DeviceTask) (code int) {
	old := &DeviceTask{}
	db.First(old, DeviceTask{DevID: dt.DevID})
	if old.ID == 0 {
		return def.TASK_NOT_EXIST
	}
	db.Model(old).Updates(dt)
	dt.ID = old.ID
	//删除定时任务
	db.Delete(Task{}, Task{DeviceTaskID: dt.ID})
	for _, v := range dt.Tasks {
		v.DeviceTaskID = dt.ID
		db.Create(v)
	}
	return def.SUCCESS
}

// GetDevSchedule 接收设备ID和用户名返回设备下的定时任务
func GetDevSchedule(dID string) (deviceTask *DeviceTask) {
	dt := &DeviceTask{}
	db.Preload("Tasks").First(dt, DeviceTask{DevID: dID})
	return dt
}

// LoadAllTasks 返回所有定时任务
func LoadAllTasks() []DeviceTask {
	dTasks := make([]DeviceTask, 0)
	db.Find(&dTasks)
	return dTasks
}
