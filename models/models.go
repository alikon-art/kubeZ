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

// 集群List返回数据的结构体
type ClusterData struct {
	ClusterID   string                 `json:"clusterid"`
	ClusterName string                 `json:"clustername"`
	Annotations map[string]interface{} `json:"annotations"`
	CreateTime  string                 `json:"createtime"`
}

// 集群Add,Update传入数据的结构体
type ClusterDataWConfig struct {
	ClusterID     string                 `json:"clusterid"`
	ClusterName   string                 `json:"clustername"`
	Annotations   map[string]interface{} `json:"annotations"`
	ClusterConfig string                 `json:"clusterconfig"`
}
