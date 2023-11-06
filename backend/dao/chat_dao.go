package dao

import (
	"1037Market/ds"
	"1037Market/mysqlDb"
	"strconv"
)

func GetSingleSessIdByStuIds(studentId1 string, studentId2 string) (sessionId int, err error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return 0, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	if studentId1 > studentId2 {
		studentId1, studentId2 = studentId2, studentId1
	}
	rows, err := db.Query("select sessionId from CHAT_SESSIONS where user1Id = ? and user2Id = ?",
		studentId1, studentId2)
	if err != nil {
		return 0, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	if !rows.Next() { // session不存在则创建一个session
		txn, err := db.Begin()
		if err != nil {
			return 0, NewErrorDao(ErrTypeDatabaseExec, err.Error())
		}
		defer txn.Rollback()
		result, err := txn.Exec("insert into CHAT_SESSIONS(user1Id, user2Id) values(?, ?)",
			studentId1, studentId2)
		if err != nil {
			return 0, NewErrorDao(ErrTypeDatabaseExec, err.Error())
		}
		if _, err := result.RowsAffected(); err != nil {
			return 0, NewErrorDao(ErrTypeDatabaseExec, err.Error())
		}
		rows.Close()
		rows, err = txn.Query("select sessionId from CHAT_SESSIONS where user1Id = ? and user2Id = ?",
			studentId1, studentId2)
		if err != nil {
			return 0, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
		}
		if !rows.Next() {
			return 0, NewErrorDao(ErrTypeAffectRows, "nothing affected")
		}
		var id int
		if err = rows.Scan(&id); err != nil {
			return 0, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		txn.Commit()
		return id, nil
	}

	if err = rows.Scan(&sessionId); err != nil {
		return 0, NewErrorDao(ErrTypeDatabaseScanRows, err.Error())
	}

	return sessionId, nil
}

func GetSessIdListBySingleStuId(studentId string) ([]int, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()
	rows, err := db.Query("select sessionId from CHAT_SESSIONS where user1Id = ? or user2Id = ?",
		studentId, studentId)
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	lst := make([]int, 0)
	for rows.Next() {
		var id int
		if err = rows.Scan(&id); err != nil {
			return nil, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		lst = append(lst, id)
	}
	return lst, nil
}

func GetTwoStuInfosBySessId(sessionId string) ([]ds.UserInfoGot, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select user1Id, user2Id from CHAT_SESSIONS where sessionId = ?", sessionId)
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, NewErrorDao(ErrTypeNoSuchSession, sessionId+" no such session")
	}

	var stuId1, stuId2 string
	if err = rows.Scan(&stuId1, &stuId2); err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseScanRows, err.Error())
	}

	lst := make([]ds.UserInfoGot, 0)
	stuId1Info, err := GetUserInfo(stuId1)
	if err != nil {
		return nil, err
	}
	lst = append(lst, stuId1Info)
	stuId2Info, err := GetUserInfo(stuId2)
	if err != nil {
		return nil, err
	}
	lst = append(lst, stuId2Info)
	return lst, nil
}

// GetNewestMsgIdBySessId find the max message id in the session
func GetNewestMsgIdBySessId(sessionId int) (int, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return 0, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select max(messageId) from CHAT_MESSAGES where sessionId = ?", sessionId)
	if err != nil {
		return 0, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()
	// note: "-1" means there is no message in the session
	if !rows.Next() {
		return -1, NewErrorDao(ErrTypeNoSuchMessage, strconv.Itoa(sessionId)+" no such message")
	}
	var messageId int
	if err = rows.Scan(&messageId); err != nil {
		return 0, NewErrorDao(ErrTypeDatabaseScanRows, err.Error())
	}
	return messageId, nil
}

func GetNMsgIdsFromKthLastBySessId(sessionId, k, n int) ([]int, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select messageId from CHAT_MESSAGES where sessionId = ? order by messageId desc limit ?, ?",
		sessionId, k-1, n)
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, NewErrorDao(ErrTypeNoSuchMessage, strconv.Itoa(sessionId)+" no such message")
	}

	lst := make([]int, 0)
	for rows.Next() {
		var id int
		if err = rows.Scan(&id); err != nil {
			return nil, NewErrorDao(ErrTypeDatabaseScanRows, err.Error())
		}
		lst = append(lst, id)
	}
	return lst, nil
}

func GetMsgInfoByMsgId(MessageId int) (ds.MsgGot, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return ds.MsgGot{}, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select * from CHAT_MESSAGES where messageId = ?", MessageId)
	if err != nil {
		return ds.MsgGot{}, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()
	if !rows.Next() {
		return ds.MsgGot{}, NewErrorDao(ErrTypeNoSuchMessage, "no such message")
	}

	var msg ds.MsgGot
	var isFromUser1 bool
	err = rows.Scan(&msg.MessageId, &msg.SessionId, &msg.MsgType, &msg.SendTime, &msg.Content, &isFromUser1)
	if err != nil {
		return ds.MsgGot{}, NewErrorDao(ErrTypeDatabaseScanRows, err.Error())
	}

	db.Close()
	rows, err = db.Query("select user1Id, user2Id from CHAT_SESSIONS where sessionId = ?", msg.SessionId)
	if err != nil {
		return ds.MsgGot{}, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	if !rows.Next() {
		return ds.MsgGot{}, NewErrorDao(ErrTypeNoSuchSession, strconv.Itoa(MessageId)+"no such session")
	}
	var user1Id, user2Id string
	err = rows.Scan(&user1Id, &user2Id)
	if err != nil {
		return ds.MsgGot{}, NewErrorDao(ErrTypeDatabaseScanRows, err.Error())
	}
	if isFromUser1 {
		msg.FromId = user1Id
	} else {
		msg.FromId = user2Id
	}
	return msg, nil
}