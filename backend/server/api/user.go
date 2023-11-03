package api

import (
	"1037Market/dao"
	"1037Market/ds"
	"1037Market/mysqlDb"
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/http"
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

func RegisterEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := os.Open("/var/EMAILPASSWORD")
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		// 创建一个Scanner用于读取文件
		scanner := bufio.NewScanner(file)

		// 读取第一行
		var psw string
		if scanner.Scan() {
			psw = scanner.Text() // Text方法返回不带换行符的当前行
		}

		// 检查是否有读取错误
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
		e := email.NewEmail()
		e.From = "Franky <1255411561@qq.com>"

		to := c.Query("studentId")

		e.To = []string{to + "@hust.edu.cn"}
		e.Subject = "1037Market 注册"
		captcha := generateRandomDigits(6)
		e.Text = []byte("1037Market\n您的验证码是：" + captcha)
		err = e.Send("smtp.qq.com:25", smtp.PlainAuth("", "1255411561@qq.com", psw, "smtp.qq.com"))

		if err != nil {
			c.String(400, err.Error())
			return
		}
		studentId2Captcha[to] = captcha
		c.String(200, "OK")
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user ds.RegisterUser
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		if user.EmailCaptcha != studentId2Captcha[user.StudentId] {
			c.String(400, "验证码错误")
			fmt.Println(studentId2Captcha[user.StudentId] + "  " + user.EmailCaptcha)
			return
		}

		if err := dao.AddNewUser(user); err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}

		c.String(200, "OK")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		type LoginUser struct {
			StudentId      string `json:"studentId"`
			HashedPassword string `json:"hashedPassword"`
		}
		var user LoginUser
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(400, "参数错误")
			return
		}

		db, err := mysqlDb.GetConnection()
		if err != nil {
			c.String(500, err.Error())
			return
		}

		rows, err := db.Query("select pswHash from USERS where userId = ?", user.StudentId)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		defer rows.Close()
		if !rows.Next() {
			c.String(400, "用户不存在")
			return
		}
		var realPsw string
		err = rows.Scan(&realPsw)
		if err != nil {
			c.String(500, err.Error())
			return
		}

		if realPsw != user.HashedPassword {
			c.String(400, "用户名或密码错误")
			fmt.Println(realPsw + " " + user.HashedPassword)
			return
		}
		cookieString := generateRandomDigits(16)
		cookie := &http.Cookie{
			Name:     "user",
			Value:    cookieString,
			Path:     "/",
			HttpOnly: false,
			Expires:  time.Now().Add(24 * time.Hour), // 设置 cookie 的过期时间
		}
		http.SetCookie(c.Writer, cookie)

		// check whether already has a cookie
		rows, err = db.Query("select cookie from COOKIES where userId = ?", user.StudentId)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		defer rows.Close()
		if rows.Next() { // already has a cookie in DB
			_, err = db.Exec("update COOKIES set cookie = ? where userId = ?", cookieString, user.StudentId)
		} else { // no cookie
			_, err = db.Exec("insert into COOKIES values(?, ?)", user.StudentId, cookieString)
		}

		if err != nil {
			c.String(500, err.Error())
			return
		}

		c.String(http.StatusOK, "Cookie has been set!")
	}
}

func UpdateUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("user")
		if err != nil {
			c.String(http.StatusBadRequest, "no cookie is set")
			return
		}

		type UserInfo struct {
			NickName string `json:"nickName"`
			Avatar   string `json:"avatar"`
			Contact  string `json:"contact"`
		}
		var userInfo UserInfo
		err = c.ShouldBindJSON(&userInfo)
		if err != nil {
			c.String(http.StatusBadRequest, "bind error: %s", err)
			return
		}

		db, err := mysqlDb.GetConnection()
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}

		rows, err := db.Query("select userId from COOKIES where cookie = ?", cookie)
		defer rows.Close()
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		if !rows.Next() {
			c.String(http.StatusBadRequest, "no such user")
			return
		}

		var userId string
		if err = rows.Scan(&userId); err != nil {
			c.String(http.StatusInternalServerError, "scan error: %s", err)
			return
		}

		_, err = db.Exec("update USER_INFOS set nickName = ?, avatar = ?, contact = ? where userId = ?",
			userInfo.NickName, userInfo.Avatar, userInfo.Contact, userId)
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		c.String(http.StatusOK, "OK")
	}
}

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		type UserInfo struct {
			UserId   string `json:"userId"`
			NickName string `json:"nickName"`
			Avatar   string `json:"avatar"`
			Contact  string `json:"contact"`
		}
		id := c.Query("studentId")
		db, err := mysqlDb.GetConnection()
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		rows, err := db.Query("select * from USER_INFOS where userId = ?", id)
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		if !rows.Next() {
			c.String(http.StatusBadRequest, "no such user")
			return
		}
		var userInfo UserInfo
		if err = rows.Scan(&userInfo.UserId, &userInfo.NickName, &userInfo.Avatar, &userInfo.Contact); err != nil {
			c.String(http.StatusInternalServerError, "scan error: %s", err)
			return
		}

		c.JSON(http.StatusOK, userInfo)
	}

}
