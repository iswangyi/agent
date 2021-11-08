package settings

import "fmt"

const version = "0.0.1"

func GetVersion() string {
	fmt.Printf("当前版本：%v \n", version)
	return version
}
