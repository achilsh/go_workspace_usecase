package main

///////////////////////// 新增的 独立module 必须要在 workspace 目录里运行 go work user  this_module_dir_name, 这样此时该module 中的接口才能被其他外面独立的module 给使用到。
import (
	"fmt"

	loginterface "second_svr/log_interface"
)

func main() {
	fmt.Println("this is second svr")
	loginterface.LogInfer()
}
