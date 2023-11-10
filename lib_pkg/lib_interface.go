package libpkg

import (
	"fmt"
	"otherlib"
)

// DateDemo sidfad
type DateDemo struct {
	a int
	B string
}

// CallLib sidfadfa
func CallLib() {
	fmt.Println("this is lib")
	otherlib.OtherLib()

}
