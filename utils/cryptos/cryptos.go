package cryptos

import (
	"kubez_project/utils/base64s"
	"kubez_project/utils/logs"
)

// 整个程序使用的加解密函数

func Encrypt(in string) (out string) {
	out = base64s.EnCode(in)
	return
}

func Decrypt(in string) (out string, err error) {
	out, err = base64s.DeCode(in)
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "解密失败")
		return "", err
	}
	return out, err

}
