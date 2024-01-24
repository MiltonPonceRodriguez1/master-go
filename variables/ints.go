package variables

import "fmt"

func ShowInts() {
	var comunInt int
	intOf32 := int32(10)
	intOf64 := int64(100)

	fmt.Println("comunInt: ", comunInt)
	fmt.Println("intOf32: ", intOf32)
	fmt.Println("intOf64: ", intOf64)

	fmt.Println("test: ", Name)

}
