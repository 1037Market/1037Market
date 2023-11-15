package test

import (
	"1037Market/dao"
	"1037Market/ds"
	"testing"
)

func newRegisterUser() ds.RegisterUser {
	return ds.RegisterUser{
		StudentId:    "U202115228",
		HashedPsw:    "123456",
		EmailCaptcha: "123456",
	}
}

func newLoginUser() ds.LoginUser {
	return ds.LoginUser{
		StudentId:      "U202115228",
		HashedPassword: "123456",
	}
}

func TestAddNewUser(t *testing.T) {
	user := newRegisterUser()
	if err := dao.AddNewUser(user); err != nil {
		t.Error(err)
	}
}

func TestLogin(t *testing.T) {
	user := newLoginUser()
	_, err := dao.Login(user)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateUserInfo(t *testing.T) {
	user := newLoginUser()
	cookie, err := dao.Login(user)
	if err != nil {
		t.Error(err)
	}
	userId, err := dao.GetUserIdByCookie(cookie)
	if err != nil {
		t.Error(err)
	}
	update := ds.UserInfoUpdated{
		NickName: "nick name",
		Avatar:   "avatar",
		Contact:  "contact",
		Address:  "address",
	}
	if err = dao.UpdateUserInfo(userId, update); err != nil {
		t.Error(err)
	}
	got, err := dao.GetUserInfo(userId)
	if err != nil {
		t.Error(err)
	}
	if got.NickName != update.NickName {
		t.Error("nick name not match")
	}
	if got.Avatar != update.Avatar {
		t.Error("avatar not match")
	}
	if got.Contact != update.Contact {
		t.Error("contact not match")
	}
	if got.Address != update.Address {
		t.Error("address not match")
	}
}

func TestUpdateUserInfoInvalidStuId(t *testing.T) {
	user := newLoginUser()
	cookie, err := dao.Login(user)
	if err != nil {
		t.Error(err)
	}
	_, err = dao.GetUserIdByCookie(cookie)
	if err != nil {
		t.Error(err)
	}
	update := ds.UserInfoUpdated{
		NickName: "nick name",
		Avatar:   "avatar",
		Contact:  "contact",
		Address:  "address",
	}
	err = dao.UpdateUserInfo("not exist", update)
	if err != nil {
		t.Error("expect to fail to update, but succeeded")
	}
}
