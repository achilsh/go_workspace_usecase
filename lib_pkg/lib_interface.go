package lib_pkg

import (
	"fmt"
	"otherlib"
)

func CallLib() {
	fmt.Println("this is lib")
	otherlib.OtherLib()

}