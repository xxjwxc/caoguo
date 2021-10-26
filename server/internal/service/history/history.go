package history

// Add 添加一个历史记录
func Add(userID, gpid, icon string) {
	getLevelDB()
	_que.Push(queInfo{UserID: userID, Gpid: gpid, Icon: icon})
}

// Get 获取
func Get(userID string) []queInfo {
	return gets(userID)
}
