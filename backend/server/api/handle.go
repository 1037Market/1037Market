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
		c.String(http.StatusInternalServerError, "database error")
	case dao.ErrTypeDatabaseQuery:
		log.Println("database query error", errorDao.Message)
		c.String(http.StatusInternalServerError, "database error")
	case dao.ErrTypeDatabaseExec:
		log.Println("database exec error", errorDao.Message)
		c.String(http.StatusInternalServerError, "database error")
	case dao.ErrTypeInvalidStudentId:
		log.Println("invalid student id", errorDao.Message)
		c.String(http.StatusBadRequest, "invalid student id")
	case dao.ErrTypeScanRows:
		log.Println("database result scan error", errorDao.Message)
		c.String(http.StatusInternalServerError, "database error")
	case dao.ErrTypeAffectRows:
		log.Println("database affect rows error", errorDao.Message)
		c.String(http.StatusInternalServerError, "database error")
	case dao.ErrTypeNoSuchUser:
		log.Println("no such user", errorDao.Message)
		c.String(http.StatusBadRequest, "no such user")
	case dao.ErrTypeNoSuchProduct:
		log.Println("no such product", errorDao.Message)
		c.String(http.StatusBadRequest, "no such product")
	case dao.ErrTypeIntParse:
		log.Println("int parse error", errorDao.Message)
		c.String(http.StatusBadRequest, "not a int")
	}
}
