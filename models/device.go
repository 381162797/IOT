package models

import (
	"mserver/def"
	"strconv"
)

//GetUserDevices 接收用户名返回包含从机的设备切片
func GetUserDevices(username string) (devices []Device, code int) {
	devices = make([]Device, 0)
	db.Preload("Slavers").Find(&devices, Device{Username: username})
	if len(devices) == 0 {
		code = def.DEVICE_NOT_EXIST
		return
	}
	code = def.SUCCESS
	return
}

//AddDevice 接收用户名，设备地址，从机数组
func AddDevice(device *Device) (code int) {
	//设备存在则不保存
	old := &Device{}
	if len(device.DevID) != 8 {
		return def.INVALID_PARAMS
	}
	db.First(old, Device{DevID: device.DevID})
	if old.ID > 0 {
		code = def.DEVICE_EXIST
		return
	}
	db.Create(device)
	code = def.SUCCESS
	return
}

// UpdateDevide 更改设备
func UpdateDevide(device *Device) (code int) {
	old := &Device{}
	if len(device.DevID) != 8 {
		return def.INVALID_PARAMS
	}
	db.First(old, Device{Username: device.Username, DevID: device.DevID})
	if old.ID == 0 {
		return def.DEVICE_NOT_EXIST
	}
	db.Model(old).Updates(device)
	device.ID = old.ID
	//删除从机
	db.Delete(Slaver{}, Slaver{DeviceID: device.ID})
	for _, v := range device.Slavers {
		//存入deviceID
		v.DeviceID = device.ID
		db.Create(v)
	}
	return def.SUCCESS
}

// DeleteDevice 删除设备以及从机
func DeleteDevice(deviceID string) (code int) {
	device := &Device{}
	code = def.SUCCESS
	id, err := strconv.ParseUint(deviceID, 10, 64)
	if err != nil {
		return def.INVALID_PARAMS
	}
	db.Delete(device, Device{ID: id})
	db.Delete(Slaver{}, Slaver{DeviceID: id})
	return
}
