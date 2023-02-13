package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"service_app/components"
	model "service_app/model"
	"service_app/utils"
)

// RegisterPaths ShowAccount godoc
//
//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	model.Account
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/accounts/{phone} [get]
func RegisterPaths(c *gin.Context) {
	phone := c.Param("phone")

	reg := regexp.MustCompile(utils.RegRuler)

	if !reg.MatchString(phone) {
		err := model.HTTPError{}
		err.Error = "error"
		err.Message = "手机号码不合法"
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	err := components.CreateUser(c)

	if err != nil {
		res := model.HTTPError{}
		res.Error = err.Error()
		res.Message = "注册失败"
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	var res model.Account
	res.Uuid = components.User.Uuid
	res.Phone = phone
	c.JSON(200, res)
}

// LoginAppPaths
// 登陆信息
func LoginAppPaths(c *gin.Context) {

}

// AccountsGetCodePaths
// 获取账号验证码
func AccountsGetCodePaths() {

}

// AccountsResetPasswordPaths
// 重制密码
func AccountsResetPasswordPaths() {

}

// AccountsUpdatePaths
// 更新账号信息
func AccountsUpdatePaths() {

}
