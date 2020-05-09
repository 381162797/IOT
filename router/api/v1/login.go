package v1

import (
	"mserver/def"
	"mserver/models"
	"mserver/tool"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRes struct {
	Code int `exzample:"200"`
	Data LoginData
	Msg  string `example:"ok"`
}

type LoginData struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

//Login 用户登录接口
// @Summary 用户登录
// @Tags 用户
// @Accept  multipart/form-data
// @Produce  json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} LoginRes
// @Failure 401 {object} LoginRes
// @Router /login [POST]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	code := def.SUCCESS
	user := models.QueryUser(username, password)
	if user.ID == 0 {
		code = def.ERROR_AUTH
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": "",
			"msg":  def.GetMsg(code),
		})
		return
	}
	token, err := tool.GenerateToken(username, password)
	if err != nil {
		code = def.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": "",
			"msg":  def.GetMsg(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": LoginData{
			Token: token,
			User:  *user,
		},
		"msg": def.GetMsg(code),
	})
}
