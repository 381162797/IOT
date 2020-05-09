package models

import (
	"time"
)

type User struct {
	ID        uint64    `json:"id" gorm:"AUTO_INCREMENT;primary_key"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	LastLogin string    `json:"lastLogin"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Device struct {
	ID         uint64    `json:"id" gorm:"AUTO_INCREMENT;primary_key"`
	DevID      string    `json:"devID"`
	DeviceName string    `json:"deviceName"`
	Status     uint8     `json:"status,omitempty" gorm:"default:0"`
	Addr       string    `json:"addr"`
	Position   string    `json:"position"`
	Username   string    `json:"username" sql:"index"`
	Slavers    []Slaver  `json:"slavers"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Slaver struct {
	ID           uint64    `json:"id" gorm:"AUTO_INCREMENT;primary_key"`
	SlaverName   string    `json:"slaverName"`
	SlaverIndex  uint8     `json:"slaverIndex"`
	TemplateID   uint64    `json:"templateID"`
	TemplateName string    `json:"templateName"`
	DeviceID     uint64    `json:"devID" gorm:"index"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Template struct {
	ID           uint64      `json:"id" gorm:"AUTO_INCREMENT;primary_key"`
	TemplateName string      `json:"templateName"`
	Username     string      `json:"username" gorm:"index"`
	DataPoints   []DataPoint `json:"dataPoints"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}

type DataPoint struct {
	ID         uint64    `json:"id" gorm:"AUTO_INCREMENT;primary_key"`
	Name       string    `json:"name"`
	Message    string    `json:"message"`
	DataType   uint8     `json:"dataType"`
	Unit       string    `json:"unit"`
	Formula    string    `json:"formula"`
	Frequency  uint64    `json:"frequency"`
	TemplateID uint64    `json:"templateID" gorm:"index"`
	CreatedAt  time.Time `json:"createdAt"`
}

type PointData struct {
	ID          uint64 `json:"id" gorm:"AUTO_INCREMENT;primary_key"`
	Data        string `json:"data"`
	CTime       int64  `json:"cTime"`
	PointID     uint64 `json:"pointID" sql:"index"`
	SlaverIndex string `json:"slaverIndex" sql:"index"`
	DevID       string `json:"devID" sql:"index"`
}

type DeviceTask struct {
	ID        uint64 `json:"id" gorm:"AUTO_INCREMENT;primary_key"`
	DevID     string `json:"devID" sql:"index"`
	Frequency uint64 `json:"frequency"`
	Tasks     []Task `json:"tasks"`
}

//Task 为从机下的单个数据点查询任务
type Task struct {
	ID           uint64 `json:"id" gorm:"AUTO_INCREMENT;primary_key"`
	PointID      uint64 `json:"pointID"`
	DeviceTaskID uint64 `json:"deviceTaskID" gorm:"index"`
	Query        string `json:"query"`
}
