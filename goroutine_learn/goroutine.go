/*
 * @Author: your name
 * @Date: 2021-06-29 15:56:37
 * @LastEditTime: 2021-06-29 16:07:23
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Editpa
 * @FilePath: \learn\goroutine_learn\goroutine.go
 */
package goroutinelearn

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//粗粒度的原子操作

var total struct {
	sync.Mutex //互斥锁
	value      int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		total.Lock()
		total.value++
		total.Unlock()
	}
}

func Test() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total.value)
}

//细粒度
var total2 uint32

func worker2(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint32
	for i = 0; i < 100; i++ {
		atomic.AddUint32(&total2, i)
	}
}

func Test2() {
	var wg *sync.WaitGroup
	wg.Add(2)
	go worker2(wg)
	go worker2(wg)
	wg.Wait()
	fmt.Println(total2)
}
