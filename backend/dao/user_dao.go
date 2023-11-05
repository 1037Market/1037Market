package dao

import (
	"1037Market/ds"
	"1037Market/mysqlDb"
	"bufio"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"os"
	"time"
)

var studentId2Captcha = make(map[string]string)

func generateRandomDigits(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = digits[rand.Intn(10)]
	}
	return string(result)
}

func RegisterEmail(studentId string) error {
	file, err := os.Open("/var/EMAILPASSWORD")
	if err != nil {
		return NewErrorDao(ErrTypeSysOpenFile, err.Error())
	}

	scanner := bufio.NewScanner(file)
	var psw string
	if scanner.Scan() {
		psw = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return NewErrorDao(ErrTypeSysReadFile, err.Error())
	}
	e := email.NewEmail()
	e.From = "Franky <1255411561@qq.com>"
	e.To = []string{studentId + "@hust.edu.cn"}
	e.Subject = "1037Market Register"
	captcha := generateRandomDigits(6)
	e.Text = []byte("1037Market\nYour captcha is: " + captcha)
	err = e.Send("smtp.qq.com:25", smtp.PlainAuth("", "1255411561@qq.com", psw, "smtp.qq.com"))
	if err != nil {
		return NewErrorDao(ErrTypeEmailSend, err.Error())
	}
	studentId2Captcha[studentId] = captcha
	return nil
}

func Register(user ds.RegisterUser) error {
	if user.EmailCaptcha != studentId2Captcha[user.StudentId] {
		return NewErrorDao(ErrTypeWrongCaptcha, "wrong captcha")
	}
	if err := AddNewUser(user); err != nil {
		return err
	}
	return nil
}

func AddNewUser(user ds.RegisterUser) error {
	// get database connection
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	//start transaction
	txn, err := db.Begin()
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer txn.Rollback()

	// insert new user
	result, err := txn.Exec("insert into USERS values(?, ?)", user.StudentId, user.HashedPsw)
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}

	// check result
	affected, err := result.RowsAffected()
	if err != nil {
		return NewErrorDao(ErrTypeAffectRows, err.Error())
	}
	if affected < 1 {
		return NewErrorDao(ErrTypeUserAlreadyExist, "user "+user.StudentId+" already exist")
	}

	// insert into USER_INFOS
	result, err = txn.Exec("insert into USER_INFOS(userId, nickName, avatar, contact) values(?, ?, ?, ?)",
		user.StudentId, user.StudentId, "null", "null")
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}
	// check result
	affected, err = result.RowsAffected()
	if err != nil {
		return NewErrorDao(ErrTypeAffectRows, err.Error())
	}
	if affected < 1 {
		return NewErrorDao(ErrTypeUserAlreadyExist, "user info"+user.StudentId+" already exist")
	}
	txn.Commit()
	return nil
}

func Login(user ds.LoginUser) (string, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return "", NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select pswHash from USERS where userId = ?", user.StudentId)
	if err != nil {
		return "", NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()
	if !rows.Next() {
		return "", NewErrorDao(ErrTypeNoSuchUser, user.StudentId+" not found")
	}
	var realPsw string
	err = rows.Scan(&realPsw)
	if err != nil {
		return "", NewErrorDao(ErrTypeScanRows, err.Error())
	}
	if realPsw != user.HashedPassword {
		return "", NewErrorDao(ErrTypeWrongPassword, user.StudentId+"wrong password")
	}
	cookieString := generateRandomDigits(16)
	rows, err = db.Query("select cookie from COOKIES where userId = ?", user.StudentId)
	if err != nil {
		return "", NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()
	if rows.Next() { // already has a cookie in DB
		_, err = db.Exec("update COOKIES set cookie = ? where userId = ?", cookieString, user.StudentId)
	} else {
		_, err = db.Exec("insert into COOKIES values(?, ?)", user.StudentId, cookieString)
	}
	if err != nil {
		return cookieString, NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}
	return cookieString, nil
}

func UpdateUserInfo(cookie string, userInfo ds.UserInfoUpdated) error {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()
	rows, err := db.Query("select userId from COOKIES where cookie = ?", cookie)
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	if !rows.Next() {
		return NewErrorDao(ErrTypeNoSuchUser, "no such user")
	}
	defer rows.Close()

	var userId string
	if err = rows.Scan(&userId); err != nil {
		return NewErrorDao(ErrTypeScanRows, err.Error())
	}
	_, err = db.Exec("update USER_INFOS set nickName = ?, avatar = ?, contact = ? where userId = ?",
		userInfo.NickName, userInfo.Avatar, userInfo.Contact, userId)
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}
	return nil
}

func GetUserInfo(userId string) (ds.UserInfoGot, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return ds.UserInfoGot{}, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select * from USER_INFOS where userId = ?", userId)
	if err != nil {
		return ds.UserInfoGot{}, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	if !rows.Next() {
		return ds.UserInfoGot{}, NewErrorDao(ErrTypeNoSuchUser, "no such user")
	}
	defer rows.Close()

	var userInfo ds.UserInfoGot
	if err = rows.Scan(&userInfo.UserId, &userInfo.NickName, &userInfo.Avatar, &userInfo.Contact); err != nil {
		return ds.UserInfoGot{}, NewErrorDao(ErrTypeScanRows, err.Error())
	}
	return userInfo, nil
}

func GetUserIdByCookie(cookie string) (string, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return "", NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select userId from COOKIES where cookie = ?", cookie)
	if err != nil {
		return "", NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}

	if !rows.Next() {
		return "", NewErrorDao(ErrTypeNoSuchUser, "wrong cookie")
	}
	var userId string
	if err = rows.Scan(&userId); err != nil {
		return "", NewErrorDao(ErrTypeScanRows, err.Error())
	}
	return userId, nil
}
