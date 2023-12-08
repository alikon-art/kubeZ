package jsons

import (
	"encoding/json"
	"fmt"
)

// 将cluster结构体转为map[string]string的方法
func Struct2MapStr(s interface{}) (map[string]string, error) {
	// 结构体数据转为byte
	bytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	// byte数据转为map[string]interface{}
	var data map[string]interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	// 将map[string]interface{}转换为map[string]string
	result := make(map[string]string)
	for k, v := range data {
		switch v := v.(type) {
		case string:
			result[k] = v
		default:
			// 如果字段的类型不是字符串，可以在此处进行适当的转换
			result[k] = fmt.Sprintf("%v", v)
		}
	}
	return result, err
}

func Struct2MapInterface(s interface{}) (map[string]interface{}, error) {
	// 结构体数据转为byte
	bytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	// byte数据转为map[string]interface{}
	var data map[string]interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, err
}
