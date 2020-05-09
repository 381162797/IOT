package models

import "time"

import "strconv"

import "mserver/def"

//SavePointData 接收pointID和数据,保存数据
func SavePointData(pID uint64, devID, data string) {
	now := time.Now().Unix()
	pData := &PointData{
		CTime:       now,
		Data:        data[2 : len(data)-4],
		PointID:     pID,
		SlaverIndex: data[:2],
		DevID:       devID,
	}
	db.Create(pData)
}

//DeletePointData 接（收pointID 删除数据
func DeletePointData(pID, sIndex, devID string) (code int) {
	code = def.SUCCESS
	id, err := strconv.ParseUint(pID, 10, 64)
	if err != nil {
		return def.INVALID_PARAMS
	}
	db.Delete(PointData{}, PointData{PointID: id, SlaverIndex: sIndex, DevID: devID})
	return
}

//GetPointData 接收pointID
func GetPointData(pID, sIndex, devID, start, end string) (datas []PointData, code int) {
	datas = make([]PointData, 0)
	id, err := strconv.ParseUint(pID, 10, 64)
	if err != nil {
		return datas, def.INVALID_PARAMS
	}
	s, err := strconv.ParseUint(start, 10, 64)
	if err != nil {
		return datas, def.INVALID_PARAMS
	}
	e, err := strconv.ParseUint(end, 10, 64)
	if err != nil {
		return datas, def.INVALID_PARAMS
	}
	db.Where("point_id = ? AND slaver_index = ? AND dev_id = ? AND c_time BETWEEN ? AND ?", id, sIndex, devID, s, e).
		Find(&datas)
	return datas, def.SUCCESS
}
