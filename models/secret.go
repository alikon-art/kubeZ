package models

// 集群Add,Update传入数据的结构体
type ClusterDataWConfig struct {
	ClusterID     string `json:"clusterid"`
	ClusterName   string `json:"clustername"`
	Version       string `json:"version"`
	Annotations   string `json:"annotations"`
	ClusterConfig string `json:"clusterconfig"`
}

// Cluster List的详细数据
type ClusterListData struct {
	ClusterID   string `json:"clusterid"`
	ClusterName string `json:"clustername"`
	Version     string `json:"version"`
	Annotations string `json:"annotations"`
	CreateTime  string `json:"createtime"`
}

// Cluster List的数据表
type ClusterList struct {
	Clusters []ClusterListData `json:"clusters"`
}

func (cl *ClusterList) AddCluster(cluster ClusterListData) {
	cl.Clusters = append(cl.Clusters, cluster)
}

// secret List的详细数据
type SecretListData struct {
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace"`
	Labels     map[string]string `json:"labels"`
	CreateTime string            `json:"createtime"`
}

// secret List的详细数据
type SecretList struct {
	Secrets []SecretListData
}

func (sl *SecretList) AddSecret(secret SecretListData) {
	sl.Secrets = append(sl.Secrets, secret)
}
