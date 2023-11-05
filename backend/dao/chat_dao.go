package dao

import (
	"1037Market/mysqlDb"
)

func GetSessionIdByStudentIds(studentId1 string, studentId2 string) (sessionId int, err error) {
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
