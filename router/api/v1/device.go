package v1

import (
	"mserver/def"
	"mserver/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeviceRes struct {
	Code int             `json:"code"`
	Data []models.Device `json:"data"`
	Msg  string          `json:"msg"`
}

type DeviceReq struct {
	DevID      string       `json:"devID"`
	DeviceName string       `json:"deviceName"`
	Addr       string       `json:"addr"`
	Position   string       `json:"position"`
	Username   string       `json:"username" sql:"index"`
	Slavers    []slaverTemp `json:"slavers"`
}

type slaverTemp struct {
	SlaverName   string `json:"slaverName"`
	SlaverIndex  uint8  `json:"slaverIndex"`
	TemplateID   uint64 `json:"templateID"`
	TemplateName string `json:"templateName"`
	DeviceID     uint64 `json:"deviceID,omitempty"`
}

// @Summary 获取用户设备
// @Tags 设备
// @Description
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Success 200 {object} DeviceReq "返回设备"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/device [get]
func GetDevice(c *gin.Context) {
	username := c.Query("username")
	res := DeviceRes{
		Code: def.SUCCESS,
	}
	if len(username) < 6 {
		res.Code = def.INVALID_PARAMS
		res.Msg = def.GetMsg(def.INVALID_PARAMS)
		c.JSON(http.StatusOK, res)
		return
	}
	devices, code := models.GetUserDevices(username)
	res.Code = code
	res.Msg = def.GetMsg(code)
	res.Data = devices
	c.JSON(http.StatusOK, res)
}

// AddDevice 添加设备
// @Summary 添加设备
// @Tags 设备
// @Description
// @Accept json
// @Produce json
// @Param device body v1.DeviceReq true "设备"
// @Success 200 {object} models.Device "返回设备"
// @Failure 400 {string} string "ok"
// @Security ApiKeyAuth
// @Router /v1/device [post]
func AddDevice(c *gin.Context) {
	req := &models.Device{}
	req.Slavers = make([]models.Slaver, 0)
	code := def.SUCCESS
	err := c.ShouldBind(req)
	if err != nil {
		code = def.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": req,
			"msg":  def.GetMsg(code),
		})
		return
	}

	code = models.AddDevice(req)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": req,
		"msg":  def.GetMsg(code),
	})
}

// UpdateDevice 更新设备
// @Summary 更新设备
// @Tags 设备
// @Description 设备存在
// @Accept json
// @Produce json
// @Param device body v1.DeviceReq true "设备"
// @Success 200 {object} models.Device "返回设备"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/device [put]]
func UpdateDevice(c *gin.Context) {
	req := &models.Device{}
	code := def.SUCCESS
	err := c.ShouldBind(req)
	if err != nil {
		code = def.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": req,
			"msg":  def.GetMsg(code),
		})
		return
	}
	code = models.UpdateDevide(req)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": req,
		"msg":  def.GetMsg(code),
	})
}

// @Summary 删除设备以及从机
// @Tags 设备
// @Produce json
// @Param deviceID query int true "设备ID"
// @Success 200 "成功"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/device [delete]
func DeleteDevice(c *gin.Context) {
	deviceID := c.Query("deviceID")
	code := models.DeleteDevice(deviceID)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": "",
		"msg":  def.GetMsg(code),
	})
}
