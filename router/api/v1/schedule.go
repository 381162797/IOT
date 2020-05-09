package v1

import (
	"mserver/def"
	"mserver/models"
	"mserver/scheduler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddScheReq struct {
	DevID     string `json:"devID"`
	Frequency uint64 `json:"frequency" exzample:"整型(单位:秒)"`
	Tasks     []TaskReq
}
type TaskReq struct {
	PointID uint64 `json:"pointID"`
	Query   string `json:"query"`
}

// AddSchedule 添加定时任务，需要数据模板ID和设备ID
// @Summary 添加设备定时任务
// @Tags 定时任务
// @Accept json
// @Produce json
// @Param device body v1.AddScheReq true "添加定时任务"
// @Success 200 {object} models.DeviceTask "返回设备"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/schedule [post]
func AddSchedule(c *gin.Context) {
	req := &models.DeviceTask{}
	err := c.ShouldBind(req)
	code := def.SUCCESS
	if err != nil {
		code = def.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": err.Error(),
			"msg":  def.GetMsg(code),
		})
		return
	}
	//存数据库
	code = models.AddSchedule(req)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": req,
		"msg":  def.GetMsg(code),
	})
	if code == def.SUCCESS {
		//发送新任务
		scheduler.TaskCH <- req
	}

}

// DeleteSchedule 删除定时任务
// @Summary 删除设备定时任务
// @Tags 定时任务
// @Produce json
// @Param devID query string true "设备的devID 8位"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/schedule [delete]
func DeleteSchedule(c *gin.Context) {
	devID := c.Query("devID")
	code := models.DeleteSchedule(devID)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": "",
		"msg":  def.GetMsg(code),
	})
	//停止定时任务
	if code == def.SUCCESS {
		stop := scheduler.StopMap.Get(devID)
		if stop != nil {
			stop <- true
		}
	}

}

// UpdateSchedule 更新并执行定时任务
// @Summary 更新并执行设备定时任务
// @Tags 定时任务
// @Accept json
// @Produce json
// @Param device body v1.AddScheReq true "更新后的定时任务"
// @Success 200 {object} models.DeviceTask "返回定时任务"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/schedule [put]
func UpdateSchedule(c *gin.Context) {
	code := def.SUCCESS
	req := &models.DeviceTask{}
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": err.Error(),
			"msg":  def.GetMsg(code),
		})
		return
	}
	code = models.UpdateSchedule(req)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": req,
		"msg":  def.GetMsg(code),
	})
	//更新任务
	if code == def.SUCCESS {
		scheduler.TaskCH <- req
	}
}

// GetSchedule 获取设备下的定时任务
// @Summary 获取设备定时任务
// @Tags 定时任务
// @Produce json
// @Param devID query string true "添加定时任务"
// @Success 200 {object} models.DeviceTask "返回设备"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/schedule [get]
func GetSchedule(c *gin.Context) {
	devID := c.Query("devID")
	res := models.GetDevSchedule(devID)
	c.JSON(http.StatusOK, gin.H{
		"code": def.SUCCESS,
		"data": res,
		"msg":  def.GetMsg(def.SUCCESS),
	})
}
