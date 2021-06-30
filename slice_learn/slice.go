/*
 * @Author: your name
 * @Date: 2021-06-29 11:38:29
 * @LastEditTime: 2021-06-29 14:14:41
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \learn\slice_learn\slice.go
 */
package slicelearn

import "fmt"

/**
*reflect.SliceHeader
*type SliceHeader struct {
*   Data uintptr
*	Len  int
*   Cap  int
*}
*
 */
func SliceTest() {
	var arr1 = [...]int{1, 2, 3, 4, 5, 6}
	arr2 := [...]string{"a", "b", "c", "d", "e", "f"}
	slice1 := arr1[:3]
	slice2 := arr2[:3]
	fmt.Println("======================================================")
	fmt.Printf(" arr1:%v\n slice1:%v\n slice2:%v\n", arr1, slice1, slice2)

	//修改源数组，切片数据会跟随改变
	arr1[1] = 10
	fmt.Println("======================================================")
	fmt.Printf(" arr1:%v\n slice1:%v\n slice2:%v\n", arr1, slice1, slice2)

	//修改切片，源数组数据会跟随改变
	slice1[2] = 100
	fmt.Println("======================================================")
	fmt.Printf(" arr1:%v\n slice1:%v\n slice2:%v\n", arr1, slice1, slice2)

	//追加切片超出容量则新开辟内存，未超出内存则直接覆盖
	slice1 = append(slice1, 8, 9)
	fmt.Println("======================================================")
	fmt.Printf(" arr1:%v\n slice1:%v\n slice2:%v\n", arr1, slice1, slice2)
	slice1 = append(slice1, 11, 12)
	fmt.Println("======================================================")
	fmt.Printf(" arr1:%v\n slice1:%v\n slice2:%v\n", arr1, slice1, slice2)
	arr1[0] = 101
	fmt.Println("======================================================")
	fmt.Printf(" arr1:%v\n slice1:%v\n slice2:%v\n", arr1, slice1, slice2)
}
