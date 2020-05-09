package v1

import (
	"mserver/def"
	"mserver/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddTempReq struct {
	TemplateName string          `json:"templateName"`
	Username     string          `json:"username" sql:"index"`
	DataPoints   []TempDataPoint `json:"dataPoints"`
}

type TempDataPoint struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	DataType  uint64 `json:"dataType"`
	Frequency uint64 `json:"frequency"`
	Unit      string `json:"unit"`
	Formula   string `json:"formula"`
}

// GetTemplate 获取用户模板
// @Summary 获取用户模板
// @Tags 模板
// @Produce json
// @Param username query string true "用户名"
// @Success 200 {object} models.Template "返回模板"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/template [get]
func GetTemplate(c *gin.Context) {
	username := c.Query("username")
	res, code := models.GetTemplate(username)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": res,
		"msg":  def.GetMsg(code),
	})
}

// AddTemplate 添加数据模板及其数据点
// @Summary 添加数据模板及其数据点
// @Tags 模板
// @Description 返回添加好的数据
// @Accept json
// @Produce json
// @Param template body v1.AddTempReq true "数据模板"
// @Success 200 {object} models.Template "返回模板"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/template [post]
func AddTemplate(c *gin.Context) {
	req := &models.Template{}
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
	code = models.AddTemplate(req)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": req,
		"msg":  def.GetMsg(code),
	})
}

// UpdateTemplate 更新数据模板
// @Summary 更新数据模板
// @Tags 模板
// @Accept json
// @Produce json
// @Param template body v1.AddTempReq true "模板参数"
// @Success 200 {object} models.Template "返回模板"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/template [put]
func UpdateTemplate(c *gin.Context) {
	req := &models.Template{}
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
	code = models.UpdateTemplate(req)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": req,
		"msg":  def.GetMsg(code),
	})
}

// DeleteTemplate 删除数据模板及其数据点
// @Summary 删除数据模板及其数据点
// @Tags 模板
// @Produce json
// @Param username query string true "用户名"
// @Param id query string true "模板ID"
// @Success 200 {string} string "删除成功"
// @Failure 401 {string} string "Unauthorized"
// @Security ApiKeyAuth
// @Router /v1/template [delete]
func DeleteTemplate(c *gin.Context) {
	username := c.Query("username")
	id := c.Query("id")
	code := def.SUCCESS
	code = models.DeleteTemplate(username, id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": "",
		"msg":  def.GetMsg(code),
	})
}
