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

type ProductGot struct {
	ProductId    int      `json:"productId"`
	Title        string   `json:"title"`
	Content      string   `json:"content"`
	Publisher    string   `json:"publisher"`
	Price        float32  `json:"price"`
	PublishTime  string   `json:"publishTime"`
	UpdateTime   string   `json:"updateTime"`
	ImageURIs    []string `json:"imageURIs"`
	Categories   []string `json:"categories"`
	IsSoldOut    bool     `json:"soldout"`
	IsSubscribed bool     `json:"subscribed"`
	NickName     string   `json:"nickName"`
	Avatar       string   `json:"avatar"`
}

type ProductPublished struct {
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
	ImageURIs  []string `json:"imageURIs"`
	Price      float32  `json:"price"`
}

type ProductUpdated struct {
	ProductId  int      `json:"productId"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
	ImageURIs  []string `json:"imageURIs"`
	Price      float32  `json:"price"`
	IsSoldOut  bool     `json:"soldout"`
}

type LoginUser struct {
	StudentId      string `json:"studentId"`
	HashedPassword string `json:"hashedPassword"`
}

type UserInfoGot struct {
	UserId   string `json:"userId"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Contact  string `json:"contact"`
	Address  string `json:"address"`
}

type UserInfoUpdated struct {
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Contact  string `json:"contact"`
	Address  string `json:"address"`
}

type CommentGot struct {
	FromId   string `json:"fromId"`
	Content  string `json:"content"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Stars    int    `json:"stars"`
}

type CommentSent struct {
	ToId    string `json:"toId"`
	Content string `json:"content"`
	Stars   int    `json:"stars"`
}

type MsgGot struct {
	MessageId int    `json:"messageId"`
	SessionId int    `json:"sessionId"`
	SendTime  string `json:"sendTime"`
	Content   string `json:"content"`
	ImageURI  string `json:"imageURI"`
	FromId    string `json:"fromId"`
}

type MsgSent struct {
	SessionId int    `json:"sessionId"`
	Content   string `json:"content"`
	ImageURI  string `json:"imageURI"`
}
