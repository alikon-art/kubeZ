package models

// secret列表的详细数据
type SecretListData struct {
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace"`
	Labels     map[string]string `json:"labels"`
	CreateTime string            `json:"createtime"`
}
type SecretList struct {
	Secrets []SecretListData
}

func (sl *SecretList) AddSecret(secret SecretListData) {
	sl.Secrets = append(sl.Secrets, secret)
}
