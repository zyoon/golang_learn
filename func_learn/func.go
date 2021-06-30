/*
 * @Author: your name
 * @Date: 2021-06-29 14:18:09
 * @LastEditTime: 2021-06-29 14:33:54
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \learn\func_learn\func.go
 */
package funclearn

import "fmt"

// 可变数量的参数
// more 对应 []int 切片类型
func Num(more ...int) {
	fmt.Printf("type:%T, more=%v\n", more, more)
}

//如果返回值命名了，可以通过名字来修改返回值，也可以通过defer语句在return语句之后修改返回值
func Inc() (v int) {
	defer func() {
		v++
	}()
	return 42
}

//其中defer语句延迟执行了一个匿名函数，因为这个匿名函数捕获了外部函数的局部变量v，这种函数我们一般叫闭包。闭包对捕获的外部变量并不是传值方式访问，而是以引用的方式访问。
func Bibao() {
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}
}

//修复闭包穿引用:新定义变量赋值
func Bibao2() {
	for i := 0; i < 3; i++ {
		a := i // 定义一个循环体内局部变量i
		defer func() { println(a) }()
	}
}

//修复闭包穿引用:闭包函数传参
func Bibao3() {
	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int) { println(i) }(i)
	}
}
