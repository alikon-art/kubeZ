package models

type ReturnDatas struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 默认状态为200,message为ok,data为空
func NewRetunData() ReturnDatas {
	returnData := ReturnDatas{
		Status:  200,
		Message: "ok",
		Data:    make(map[string]interface{}),
	}
	return returnData
}
