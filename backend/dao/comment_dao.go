package dao

import (
	"1037Market/ds"
	"1037Market/mysqlDb"
)

func CreateComment(fromId string, comment ds.CommentSent) (int, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return 0, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select userId from USERS where userId = ?", comment.ToId)
	if err != nil {
		return 0, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, NewErrorDao(ErrTypeNoSuchUser, "from "+fromId+" to "+comment.ToId+" no such user")
	}

	txn, err := db.Begin()
	if err != nil {
		return 0, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer txn.Rollback()
	result, err := txn.Exec("insert into COMMENTS(publisherId, receiverId, content, stars) values(?, ?, ?, ?)",
		fromId, comment.ToId, comment.Content, comment.Stars)
	if err != nil {
		return 0, NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return 0, NewErrorDao(ErrTypeAffectRows, err.Error())
	}
	if affected < 1 {
		return 0, NewErrorDao(ErrTypeDatabaseExec, "insert comment failed")
	}

	rows.Close()
	rows, err = txn.Query("select LAST_INSERT_ID();")
	if err != nil {
		return 0, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	if !rows.Next() {
		return 0, NewErrorDao(ErrTypeAffectRows, "no last insert id")
	}
	var commentId int
	if err = rows.Scan(&commentId); err != nil {
		return 0, NewErrorDao(ErrTypeScanRows, err.Error())
	}
	
	txn.Commit()
	return commentId, nil
}

func QueryCommentList(userId string) ([]int, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select commentId from COMMENTS where receiverId = ?", userId)
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	list := make([]int, 0)
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		list = append(list, id)
	}
	return list, nil
}

func GetCommentById(commentId string) (ds.CommentGot, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return ds.CommentGot{}, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query(
		"select publisherId, content, stars, nickName, avatar from COMMENTS, USER_INFOS "+
			"where commentId = ? and COMMENTS.publisherId = USER_INFOS.userId",
		commentId)
	defer rows.Close()

	if err != nil {
		return ds.CommentGot{}, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	if !rows.Next() {
		return ds.CommentGot{}, NewErrorDao(ErrTypeNoSuchComment, commentId+" no such comment")
	}
	var comment ds.CommentGot
	err = rows.Scan(&comment.FromId, &comment.Content, &comment.Stars, &comment.NickName, &comment.Avatar)
	if err != nil {
		return ds.CommentGot{}, NewErrorDao(ErrTypeScanRows, err.Error())
	}
	return comment, nil
}
