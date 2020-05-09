package v1

import (
	"mserver/def"
	"mserver/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRes struct {
	Code int
	Data models.User
	Msg  string
}

//UserRegister 注册用户接口
// @Summary 用户注册
// @Tags 用户
// @Accept  multipart/form-data
// @Produce  json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} UserRes
// @Failure 401 {object} UserRes
// @Router /user [POST]
func UserRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, code := models.InsertUser(username, password)
	resp := UserRes{
		Code: code,
		Data: *user,
		Msg:  def.GetMsg(code),
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteUser 删除用户如果密码正确
// @Summary 删除用户
// @Tags 用户
// @Description
// @Accept json
// @Produce json
// @Param id query string true "用户id"
// @Param password query string true "密码"
// @Success 200 {object} UserRes "成功 code = 200"
// @Router /user [delete]
func DeleteUser(c *gin.Context) {
	code := def.SUCCESS
	res := UserRes{
		Code: code,
	}
	id := c.Query("id")
	uID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		code = def.INVALID_PARAMS
		res.Code = code
		res.Msg = def.GetMsg(code)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	password := c.Query("password")
	code = models.DeleteUser(uID, password)
	res.Code = code
	res.Msg = def.GetMsg(code)
	c.JSON(http.StatusOK, res)
	return
}

// PutUser 修改用户密码
// @Summary 修改密码
// @Tags 用户
// @Description
// @Produce json
// @Param id query string true "用户id"
// @Param old query string true "旧密码"
// @Param new query string true "新密码"
// @Success 200 {object} UserRes "成功"
// @Failure 400 {object} UserRes "失败"
// @Router /user [put]
func PutUser(c *gin.Context) {
	i := c.Query("id")
	old := c.Query("old")
	new := c.Query("new")
	id, err := strconv.ParseUint(i, 10, 0)
	code := def.INVALID_PARAMS
	if err != nil {
		c.JSON(http.StatusBadRequest, UserRes{
			Code: code,
			Msg:  def.GetMsg(code),
		})
		return
	}
	code = models.UpdateUser(id, old, new)
	c.JSON(http.StatusOK, UserRes{
		Code: code,
		Msg:  def.GetMsg(code),
	})
}
