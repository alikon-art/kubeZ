// 这个包定义通用的结构体
package models

// 返回数据的通用框架
type ReturnDataFrame struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 基础返回格式,拼接于ReturnDataFrame.Data内
// List方法仅填充基础字段,GET方法返回Item对象
type BasicReturn struct {
	ClusterID  string            `json:"clusterid"`
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace"`
	Labels     map[string]string `json:"labels"`
	CreateTime string            `json:"createtime"`
	Item       interface{}       `json:"item"`
}

// 基础返回的列表形式
type BasicReturnList struct {
	Items []BasicReturn `json:"items"`
}

// 基础返回的列表添加元素的方法
func (brl *BasicReturnList) AddBasicReturn(br BasicReturn) {
	brl.Items = append(brl.Items, br)
}

// 请求数据的通用框架
// List方法仅填充基础字段,Add.Updata方法携带Item对象
type RequestDataFrame struct {
	ClusterID string            `json:"clusterid"`
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
	Item      interface{}       `json:"item"`
}

// 处理用户登录信息的结构体
type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
