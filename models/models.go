package models

// 返回数据的基本结构体
type ReturnData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 处理用户登录信息的结构体
type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 以下暂无用处======

// 请求操作多个资源的列表结构体
type RequestList struct {
	Resources []RequestListData `json:"resources"`
}

// 请求操作多个资源的详情结构体
type RequestListData struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
}

// 返回操作多个资源的列表结构体
type ReturnList struct {
	Resources []RequestListData `json:"resources"`
}

// 返回操作多个资源的详情结构体
type ReturnListData struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
}
