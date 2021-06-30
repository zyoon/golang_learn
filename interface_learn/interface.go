/*
 * @Author: your name
 * @Date: 2021-06-29 15:41:17
 * @LastEditTime: 2021-06-29 15:47:06
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \learn\interface_learn\interface.go
 */
package interfacelearn

import "fmt"

type notifer interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Println("sendNotify to user", u.name)
}
func (ad admin) notify() {
	fmt.Println("sendNotify to user", ad.name)
}

func sendNotify(n notifer) {
	n.notify()
}

func Test() {
	u := user{"张三", "245525399"}
	sendNotify(&u)
	ad := admin{"李四", "13298508293"}
	sendNotify(ad)
}
