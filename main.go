/*
 * @Author: your name
 * @Date: 2021-06-29 10:16:37
 * @LastEditTime: 2021-06-30 17:21:25
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \learn\main.go
 */
package main

import (
	"fmt"
)

func main() {
	// slicelearn.SliceTest()
	// stringlearn.StringTest()
	// funclearn.Num(1, 2, 3, 4, 5)
	// num := funclearn.Inc()
	// fmt.Println(num)
	// funclearn.Bibao()
	// funclearn.Bibao2()
	// interfacelearn.Test()
	// goroutinelearn.Test()
	// goroutinelearn.Test2()

	ch := GenerateNatural()
	for i := 0; i < 20; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		// time.Sleep(time.Second)
		ch = PrimeFilter(ch, prime)
	}
}

func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			// fmt.Println("??")
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}
