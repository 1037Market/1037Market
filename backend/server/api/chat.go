package api

/*
TODO:
chat API

新建两张表：CHAT_SESSIONS CHAT_MESSAGES

CHAT_SESSIONS: user1Id和user2Id需要保证 user1Id <= user2Id，这是需要server负责的，不然两个人的对话将会存在两个sessionId
				需要支持根据两个userId查询sessionId
CHAT_MESSAGES: 每个message属于一个session。
				type暂时分为文字和图片两种，图片存URI，文字直接存，后端不需要过多处理，只需要解析前端传来的类型
				需要支持在固定sessionId的情况下根据messageId的排序查询，因为前端需要查询从倒数第i条开始的n条消息

*/
