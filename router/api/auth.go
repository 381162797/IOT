package api

import (
	"mserver/def"
	"mserver/models"
	"mserver/tool"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthRes struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

// @Summary 获取token
// @Tags 认证
// @Description
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} api.AuthRes "返回token"
// @Failure 401 {object} api.AuthRes "Unauthorized"
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	code := def.INVALID_PARAMS
	data := ""
	username, password := c.Query("username"), c.Query("password")
	if username != "" && password != "" {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := tool.GenerateToken(username, password)
			if err != nil {
				code = def.ERROR_AUTH_TOKEN
			} else {
				data = token
				code = def.SUCCESS
			}
		} else {
			code = def.ERROR_AUTH
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  def.GetMsg(code),
		"data": data,
	})
}
