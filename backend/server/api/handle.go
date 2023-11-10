package api

import (
	"1037Market/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func handleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	errorDao, ok := err.(*dao.ErrorDao)
	if !ok {
		log.Println("unknown error", err.Error())
		c.String(http.StatusInternalServerError, "internal error")
		return
	}

	switch errorDao.Type {
	case dao.ErrTypeDatabaseConnection:
		log.Print("database connection error", errorDao.Message)
		c.String(http.StatusInternalServerError, "数据库错误\ndatabase error")
	case dao.ErrTypeDatabaseQuery:
		log.Println("database query error", errorDao.Message)
		c.String(http.StatusInternalServerError, "数据库错误\ndatabase error")
	case dao.ErrTypeDatabaseExec:
		log.Println("database exec error", errorDao.Message)
		c.String(http.StatusInternalServerError, "数据库错误\ndatabase error")
	case dao.ErrTypeInvalidStudentId:
		log.Println("invalid student id", errorDao.Message)
		c.String(http.StatusBadRequest, "invalid student id")
	case dao.ErrTypeScanRows:
		log.Println("database result scan error", errorDao.Message)
		c.String(http.StatusInternalServerError, "数据库错误\ndatabase error")
	case dao.ErrTypeAffectRows:
		log.Println("database affect rows error", errorDao.Message)
		c.String(http.StatusInternalServerError, "数据库错误\ndatabase error")
	case dao.ErrTypeNoSuchUser:
		log.Println("no such user", errorDao.Message)
		c.String(http.StatusBadRequest, "用户不存在\nno such user")
	case dao.ErrTypeNoSuchProduct:
		log.Println("no such product", errorDao.Message)
		c.String(http.StatusBadRequest, "商品不存在\nno such product")
	case dao.ErrTypeIntParse:
		log.Println("int parse error", errorDao.Message)
		c.String(http.StatusBadRequest, "not a int")
	case dao.ErrTypeSysOpenFile:
		log.Println("open file error", errorDao.Message)
		c.String(http.StatusInternalServerError, "服务错误\ninternal error")
	case dao.ErrTypeSysReadFile:
		log.Println("read file error", errorDao.Message)
		c.String(http.StatusInternalServerError, "服务错误\ninternal error")
	case dao.ErrTypeEmailSend:
		log.Println("send email error", errorDao.Message)
		c.String(http.StatusInternalServerError, "服务错误\ninternal error")
	case dao.ErrTypeWrongCaptcha:
		log.Println("wrong captcha", errorDao.Message)
		c.String(http.StatusBadRequest, "验证码错误\nwrong captcha")
	case dao.ErrTypeUserAlreadyExist:
		log.Println("user already exist", errorDao.Message)
		c.String(http.StatusBadRequest, "用户已存在\nuser already exist")
	case dao.ErrTypeWrongPassword:
		log.Println("wrong password", errorDao.Message)
		c.String(http.StatusBadRequest, "密码错误\nwrong password")
	case dao.ErrTypeWrongRequestFormat:
		log.Println("wrong request format", errorDao.Message)
		c.String(http.StatusBadRequest, "请求格式错误\nwrong request format")
	case dao.ErrTypeProductAlreadyExist:
		log.Println("product already exist", errorDao.Message)
		c.String(http.StatusBadRequest, "商品已存在\nproduct already exist")
	case dao.ErrTypeNoSuchComment:
		log.Println("no such comment", errorDao.Message)
		c.String(http.StatusBadRequest, "评论不存在\nno such comment")
	case dao.ErrTypeSysSaveFile:
		log.Println("save file error", errorDao.Message)
		c.String(http.StatusInternalServerError, "服务错误\ninternal error")
	case dao.ErrTypeNoSuchSession:
		log.Println("no such session", errorDao.Message)
		c.String(http.StatusBadRequest, "聊天不存在\nno such session")
	case dao.ErrTypeNoSuchMessage:
		log.Println("no such message", errorDao.Message)
		c.String(http.StatusBadRequest, "消息不存在\nno such message")
	}
}