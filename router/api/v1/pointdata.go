package v1

import "github.com/gin-gonic/gin"

import "mserver/models"

import "net/http"

import "mserver/def"

type PointDataRes struct {
	Code int                `json:"code"`
	Data []models.PointData `json:"data"`
	Msg  string             `json:"msg"`
}

//GetPointData 获取数据点数据
// @Summary 获取数据点数据
// @Tags 数据
// @Produce json
// @Param pID query string true "pointID"
// @Param slaverIndex query string true "从机地址"
// @Param devID query string true "devID"
// @Param start query string true "开始时间(unix)"
// @Param end query string true "结束时间(unix)"
// @Success 200 {object} v1.PointDataRes "返回数据"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/data [get]
func GetPointData(c *gin.Context) {
	s := c.Query("start")
	sIndex := c.Query("slaverIndex")
	e := c.Query("end")
	id := c.Query("pID")
	devID := c.Query("devID")
	res, code := models.GetPointData(id, sIndex, devID, s, e)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": res,
		"msg":  def.GetMsg(code),
	})
}

// DeletePointData 删除数据点下的数据
// @Summary 删除数据点下的数据
// @Tags 数据
// @Produce json
// @Param pID query string true "pointID"
// @Param slaverIndex query string true "从机地址"
// @Param devID query string true "devID"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/data [delete]
func DeletePointData(c *gin.Context) {
	id := c.Query("pID")
	sIndex := c.Query("slaverIndex")
	devID := c.Query("devID")
	code := models.DeletePointData(id, sIndex, devID)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": "",
		"msg":  def.GetMsg(code),
	})
}
