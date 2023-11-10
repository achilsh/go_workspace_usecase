package main

import (
	libpkg "libpkg"
	loginterface "second_svr/log_interface"
)

func main() {
	libpkg.CallLib()
	loginterface.LogInfer()
}
