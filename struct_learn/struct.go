/*
 * @Author: your name
 * @Date: 2021-06-29 14:46:20
 * @LastEditTime: 2021-06-29 14:49:01
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \learn\struct_learn\struct.go
 */

package structlearn

import (
	"fmt"
	"image/color"
)

//我们这里将Point嵌入到ColoredPoint来提供X和Y这两个字段
type Point struct {
	X, Y float64
}
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func Test() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"
}
