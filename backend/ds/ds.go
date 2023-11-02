package ds

import "time"

type ProductInfo struct {
	ProductId   uint32    `json:"productId"`
	UserID      uint32    `json:"userID"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Images      []string  `json:"images"`
	Price       uint32    `json:"price"`
	CreateTime  time.Time `json:"publishTime"`
	UpdateTime  time.Time `json:"updateTime"`
}

type UserInfo struct {
	Email    string `json:"userId"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Contact  uint64 `json:"contact"`
}

type RegisterUser struct {
	StudentId    string `json:"studentId"`
	HashedPsw    string `json:"hashedPassword"`
	EmailCaptcha string `json:"emailCaptcha"`
}
