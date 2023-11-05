package api

import (
	"1037Market/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
TODO:
chat实现思路

新建两张表：CHAT_SESSIONS CHAT_MESSAGES

CHAT_SESSIONS: user1Id和user2Id需要保证 user1Id <= user2Id，这是需要server负责的，不然两个人的对话将会存在两个sessionId
				需要支持根据两个userId查询sessionId
CHAT_MESSAGES: 每个message属于一个session。
				type暂时分为文字和图片两种，图片存URI，文字直接存，后端不需要过多处理，只需要解析前端传来的类型
				需要支持在固定sessionId的情况下根据messageId的排序查询，因为前端需要查询从倒数第i条开始的n条消息
*/

/*
TODO:
需要的API
	根据用户ID查询SessionID 即CHAT_SESSIONS表中user1Id或user2Id == studentID 的SessionID

	根据sessionID返回参与聊天的两个人的信息

	根据sessionID查询该session中最近一次的messageID

	查询同一session中的从 倒数第i条开始的n条消息的messageID

	根据messageID返回message信息

*/

func GetSessionIdByStudentIds() gin.HandlerFunc {
	return func(c *gin.Context) {
		studentId1 := c.Query("studentId1")
		studentId2 := c.Query("studentId2")
		sessionId, err := dao.GetSessionIdByStudentIds(studentId1, studentId2)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, sessionId)
	}
}
