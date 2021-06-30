package stringlearn

import "fmt"

/**
*底层结构 reflect.StringHeader
*type StringHeader struct {
*    Data uintptr
*    Len  int
*}
*
**/
func StringTest() {
	var str string = "12Hj好人"
	fmt.Printf("==================str=%s type=%T ==========================\n", str, str)
	fmt.Println(len(str))
	str2 := []byte("12Hj好人")
	fmt.Printf("==================str2=%s type=%T ==========================\n", str2, str2)
	fmt.Println(len(str2))
	str3 := []rune("12Hj好人")
	fmt.Printf("==================str3=%s type=%T ==========================\n", string(str3), str3)
	fmt.Println(len(str3))
}
