package api

import (
	"1037Market/dao"
	"1037Market/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(c *gin.Context, err error) {
	log := logger.GetInstance()
	if err == nil {
		return
	}

	errorDao, ok := err.(*dao.ErrorDao)
	if !ok {
		log.Error("unknown error", err.Error())
		c.String(http.StatusInternalServerError, "internal error")
		return
	}

	switch errorDao.Type {
	case dao.ErrTypeDatabaseConnection:
		log.Fatal("database connection error", errorDao.Message)
		c.String(http.StatusInternalServerError, "数据库错误\ndatabase error")
	case dao.ErrTypeDatabaseQuery:
		log.Fatal("database query error", errorDao.Message)
		c.String(http.StatusInternalServerError, "数据库错误\ndatabase error")
	case dao.ErrTypeDatabaseExec:
		log.Error("database exec error", errorDao.Message)
		c.String(http.StatusInternalServerError, "数据库错误\ndatabase error")
	case dao.ErrTypeInvalidStudentId:
		log.Error("invalid student id", errorDao.Message)
		c.String(http.StatusBadRequest, "invalid student id")
	case dao.ErrTypeScanRows:
		log.Error("database result scan error", errorDao.Message)
		c.String(http.StatusInternalServerError, "数据库错误\ndatabase error")
	case dao.ErrTypeAffectRows:
		log.Error("database affect rows error", errorDao.Message)
		c.String(http.StatusInternalServerError, "数据库错误\ndatabase error")
	case dao.ErrTypeNoSuchUser:
		log.Error("no such user", errorDao.Message)
		c.String(http.StatusBadRequest, "用户不存在\nno such user")
	case dao.ErrTypeNoSuchProduct:
		log.Error("no such product", errorDao.Message)
		c.String(http.StatusBadRequest, "商品不存在\nno such product")
	case dao.ErrTypeIntParse:
		log.Error("int parse error", errorDao.Message)
		c.String(http.StatusBadRequest, "not a int")
	case dao.ErrTypeSysOpenFile:
		log.Error("open file error", errorDao.Message)
		c.String(http.StatusInternalServerError, "服务错误\ninternal error")
	case dao.ErrTypeSysReadFile:
		log.Error("read file error", errorDao.Message)
		c.String(http.StatusInternalServerError, "服务错误\ninternal error")
	case dao.ErrTypeEmailSend:
		log.Error("send email error", errorDao.Message)
		c.String(http.StatusInternalServerError, "服务错误\ninternal error")
	case dao.ErrTypeWrongCaptcha:
		log.Error("wrong captcha", errorDao.Message)
		c.String(http.StatusBadRequest, "验证码错误\nwrong captcha")
	case dao.ErrTypeUserAlreadyExist:
		log.Error("user already exist", errorDao.Message)
		c.String(http.StatusBadRequest, "用户已存在\nuser already exist")
	case dao.ErrTypeWrongPassword:
		log.Error("wrong password", errorDao.Message)
		c.String(http.StatusBadRequest, "密码错误\nwrong password")
	case dao.ErrTypeWrongRequestFormat:
		log.Error("wrong request format", errorDao.Message)
		c.String(http.StatusBadRequest, "请求格式错误\nwrong request format")
	case dao.ErrTypeProductAlreadyExist:
		log.Error("product already exist", errorDao.Message)
		c.String(http.StatusBadRequest, "商品已存在\nproduct already exist")
	case dao.ErrTypeNoSuchComment:
		log.Error("no such comment", errorDao.Message)
		c.String(http.StatusBadRequest, "评论不存在\nno such comment")
	case dao.ErrTypeSysSaveFile:
		log.Error("save file error", errorDao.Message)
		c.String(http.StatusInternalServerError, "服务错误\ninternal error")
	case dao.ErrTypeNoSuchSession:
		log.Error("no such session", errorDao.Message)
		c.String(http.StatusBadRequest, "聊天不存在\nno such session")
	case dao.ErrTypeNoSuchMessage:
		log.Error("no such message", errorDao.Message)
		c.String(http.StatusBadRequest, "消息不存在\nno such message")
	}
}
