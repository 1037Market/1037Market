package server

import (
	"1037Market/mysqlDb"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"
)

var studentId2Captcha = make(map[string]string)
var studentId2Cookie = make(map[string]string)

func generateRandomDigits(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = digits[rand.Intn(10)]
	}
	return string(result)
}

func registerEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := email.NewEmail()
		e.From = "Franky <1255411561@qq.com>"

		to := c.Query("studentId")

		e.To = []string{to + "@hust.edu.cn"}
		e.Subject = "1037Market 注册"
		captcha := generateRandomDigits(6)
		e.Text = []byte("1037Market\n您的验证码是：" + captcha)
		err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "1255411561@qq.com", "replace it with your password", "smtp.qq.com"))

		if err != nil {
			c.String(400, err.Error())
			return
		}
		studentId2Captcha[to] = captcha
		c.String(200, "OK")
	}
}

func register() gin.HandlerFunc {
	return func(c *gin.Context) {

		type RegisterUser struct {
			StudentId    string `json:"studentId"`
			HashedPsw    string `json:"hashedPassword"`
			EmailCaptcha string `json:"emailCaptcha"`
		}
		var user RegisterUser
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(400, err.Error())
			return
		}
		if user.EmailCaptcha != studentId2Captcha[user.StudentId] {
			c.String(400, "验证码错误")
			fmt.Println(studentId2Captcha[user.StudentId] + "  " + user.EmailCaptcha)
			return
		}

		db, err := mysqlDb.GetDefault()
		if err != nil {
			c.String(500, err.Error())
			return
		}

		result, err := db.Exec("insert into USERS values(?, ?)", user.StudentId, user.HashedPsw)
		if err != nil {
			c.String(500, err.Error())
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.String(500, err.Error())
			return
		}
		if rowsAffected == 0 {
			c.String(400, "用户已存在")
			return
		}

		c.String(200, "OK")
	}
}

func login() gin.HandlerFunc {
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

		db, err := mysqlDb.GetDefault()
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
		studentId2Cookie[user.StudentId] = cookieString
		cookie := &http.Cookie{
			Name:     "user",
			Value:    cookieString,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(24 * time.Hour), // 设置 cookie 的过期时间
		}
		http.SetCookie(c.Writer, cookie)

		c.String(http.StatusOK, "Cookie has been set!")
	}
}