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

	if !rows.Next() { // TODO: session不存在则创建一个session
		return 0, NewErrorDao(ErrTypeInvalidStudentId, "no corresponding session")
	}

	if err = rows.Scan(&sessionId); err != nil {
		return 0, NewErrorDao(ErrTypeDatabaseScanRows, err.Error())
	}

	return sessionId, nil
}
