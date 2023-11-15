package test

import (
	"1037Market/dao"
	"1037Market/ds"
	"testing"
)

func TestGetTwoStuInfosBySessId(t *testing.T) {
	user0 := "U202115224"
	user1 := "U202115228"
	// note that user0 < user1
	sessionId, err := dao.GetSingleSessIdByStuIds(user0, user1)
	if err != nil {
		t.Error(err)
	}
	got, err := dao.GetTwoStuInfosBySessId(sessionId)
	if err != nil {
		t.Error(err)
	}
	if len(got) != 2 {
		t.Error("expect 2, got", len(got))
	}
	if got[0].UserId != user0 {
		t.Error("expect", user0, "got", got[0].UserId)
	}
	if got[1].UserId != user1 {
		t.Error("expect", user1, "got", got[1].UserId)
	}
	// get user info
	userInfo0, err := dao.GetUserInfo(user0)
	if err != nil {
		t.Error(err)
	}
	userInfo1, err := dao.GetUserInfo(user1)
	if err != nil {
		t.Error(err)
	}
	if got[0].NickName != userInfo0.NickName || got[1].NickName != userInfo1.NickName {
		t.Error("expect", userInfo0.NickName, userInfo1.NickName, "got", got[0].NickName, got[1].NickName)
	}
	if got[0].Avatar != userInfo0.Avatar || got[1].Avatar != userInfo1.Avatar {
		t.Error("expect", userInfo0.Avatar, userInfo1.Avatar, "got", got[0].Avatar, got[1].Avatar)
	}
	if got[0].Contact != userInfo0.Contact || got[1].Contact != userInfo1.Contact {
		t.Error("expect", userInfo0.Contact, userInfo1.Contact, "got", got[0].Contact, got[1].Contact)
	}
	if got[0].Address != userInfo0.Address || got[1].Address != userInfo1.Address {
		t.Error("expect", userInfo0.Address, userInfo1.Address, "got", got[0].Address, got[1].Address)
	}
}

func TestGetSessIdListBySingleStuId(t *testing.T) {
	user := "U202115224"
	lst, err := dao.GetSessIdListBySingleStuId(user)
	if err != nil {
		t.Error(err)
	}
	for _, sessionId := range lst {
		got, err := dao.GetTwoStuInfosBySessId(sessionId)
		if err != nil {
			t.Error(err)
		}
		if len(got) != 2 {
			t.Error("expect 2, got", len(got))
		}
		if got[0].UserId != user && got[1].UserId != user {
			t.Error("expect", user, "got", got[0].UserId, got[1].UserId)
		}
	}
}

func TestGetMsgInfoByMsgId(t *testing.T) {
	user0 := "U202115224"
	user1 := "U202115228"
	// note that user0 < user1
	sessionId, err := dao.GetSingleSessIdByStuIds(user0, user1)
	if err != nil {
		t.Error(err)
	}
	send := ds.MsgSent{
		SessionId: sessionId,
		Content:   "content",
		ImageURI:  "image uri",
	}
	messageId, err := dao.SendMsg(user0, send)
	if err != nil {
		t.Error(err)
	}
	got, err := dao.GetMsgInfoByMsgId(messageId)
	if err != nil {
		t.Error(err)
	}
	if got.MessageId != messageId {
		t.Error("expect", messageId, "got", got.MessageId)
	}
	if got.SessionId != sessionId {
		t.Error("expect", sessionId, "got", got.SessionId)
	}
	if got.Content != send.Content {
		t.Error("expect", send.Content, "got", got.Content)
	}
	if got.ImageURI != send.ImageURI {
		t.Error("expect", send.ImageURI, "got", got.ImageURI)
	}
}
