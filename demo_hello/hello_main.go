package main

import (
	lib_pkg "libpkg"
	loginterface "second_svr/log_interface"
)

func main() {
	lib_pkg.CallLib()
	loginterface.LogInfer()
}