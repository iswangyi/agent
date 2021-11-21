package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// PK 合并部分监控指标信息
func PK(endpoint, metric string, tags map[string]string) string {
	if tags == nil || len(tags) == 0 {
		return fmt.Sprintf("%s/%s", endpoint, metric)
	}
	return fmt.Sprintf("%s/%s/%s", endpoint, metric, SortedTags(tags))
}

// IsExist 判断文件是否存在
func IsExist(fp string) bool {

	_, err := os.Stat(fp)

	return err == nil || os.IsExist(err)
}

//  toString 读取配置文件
func toString(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ToTrimString 解析配置文件
func ToTrimString(filePath string) (string, error) {
	str, err := toString(filePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(str), nil
}
