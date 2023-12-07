package models

// 返回数据的基本结构体
type ReturnData struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 处理集群数据的基本结构体
type ClusterData struct {
	ClusterID   string                 `json:"clusterid"`
	ClusterName string                 `json:"clustername"`
	Annotations map[string]interface{} `json:"annotations"`
}

// 带有kubeconfig的结构体
type ClusterDataWConfig struct {
	ClusterID     string                 `json:"clusterid"`
	ClusterName   string                 `json:"clustername"`
	Annotations   map[string]interface{} `json:"annotations"`
	ClusterConfig string                 `json:"clusterconfig"`
}
