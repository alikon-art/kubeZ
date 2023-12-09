package models

// 返回数据的基本结构体
type ReturnData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Add,Updata请求通用结构体
type RequestData struct {
	ClusterID string            `json:"clusterid"`
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
}

// 处理用户登录信息的结构体
type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// List 方法通用返回体
type ReturnList struct {
	Resources []ReturnListData `json:"resources"`
}

// List 方法通用返回体数据
type ReturnListData struct {
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace"`
	Labels     map[string]string `json:"labels"`
	CreateTime string            `json:"createtime"`
}

// List 结构体添加数据方法
func (rl *ReturnList) AddReturnListData(returnListData ReturnListData) {
	rl.Resources = append(rl.Resources, returnListData)
}
