package test

import (
	"1037Market/dao"
	"1037Market/ds"
	"1037Market/logger"
	"testing"
)

var log = logger.GetInstance()

func TestLoginInjection(t *testing.T) {
	_, err := dao.Login(ds.LoginUser{
		StudentId:      "'; drop table USERS;",
		HashedPassword: "aaaaaa",
	})
	log.Info(err)

	if err == nil {
		t.Error("expect to fail")
	}

	_, err = dao.Login(ds.LoginUser{
		StudentId:      "\"; drop table USERS;",
		HashedPassword: "aaaaaa",
	})
	log.Info(err)
	if err == nil {
		t.Error("expect to fail")
	}
}

func TestProductInjection(t *testing.T) {
	product := makeProduct()
	product.Title = "111\"); drop table PRODUCTS"

	_, err := dao.PublishProduct("t1", product)
	log.Info(err)
	if err == nil {
		t.Error("expect to fail")
	}

	product.Title = "111'); drop table PRODUCTS"

	_, err = dao.PublishProduct("t1", product)
	log.Info(err)
	if err == nil {
		t.Error("expect to fail")
	}
}

func TestCommentXSS(t *testing.T) {
	
}
