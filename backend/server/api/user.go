package api

import (
	"1037Market/dao"
	"1037Market/ds"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		studentId := c.Query("studentId")
		err := dao.RegisterEmail(studentId)
		if err != nil {
			handleError(c, err)
			return
		}
		c.String(http.StatusOK, "OK")
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user ds.RegisterUser
		if err := c.ShouldBindJSON(&user); err != nil {
			handleError(c, dao.NewErrorDao(dao.ErrTypeWrongRequestFormat, err.Error()))
			return
		}
		err := dao.Register(user)
		if err != nil {
			handleError(c, err)
			return
		}
		c.String(http.StatusOK, "OK")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user ds.LoginUser
		if err := c.ShouldBindJSON(&user); err != nil {
			handleError(c, dao.NewErrorDao(dao.ErrTypeWrongRequestFormat, err.Error()))
			return
		}
		cookieString, err := dao.Login(user)
		if err != nil {
			handleError(c, err)
			return
		}
		c.String(http.StatusOK, cookieString)
	}
}

func UpdateUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.Query("user")

		type UserInfo struct {
			NickName string `json:"nickName"`
			Avatar   string `json:"avatar"`
			Contact  string `json:"contact"`
			Address  string `json:"address"`
		}
		var userInfo ds.UserInfoUpdated
		err := c.ShouldBindJSON(&userInfo)
		if err != nil {
			handleError(c, dao.NewErrorDao(dao.ErrTypeWrongRequestFormat, err.Error()))
			return
		}

		err = dao.UpdateUserInfo(cookie, userInfo)
		if err != nil {
			handleError(c, err)
			return
		}
		c.String(http.StatusOK, "OK")
	}
}

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("studentId")
		userInfo, err := dao.GetUserInfo(id)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, userInfo)
	}
}
