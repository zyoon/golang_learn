/*
 * @Author: your name
 * @Date: 2021-06-29 15:56:37
 * @LastEditTime: 2021-06-30 15:57:11
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Editpa
 * @FilePath: \learn\goroutine_learn\goroutine.go
 */
package goroutinelearn

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
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

//单例模式 原子操作配合互斥锁可以实现非常高效的单件模式。互斥锁的代价比普通整数的原子读写高很多，在性能敏感的地方可以增加一个数字型的标志位，通过原子检测标志位状态降低互斥锁的使用次数来提高性能。
type singleton struct{}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

func Instance() *singleton {
	//原子读（防止并发）
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	//互斥锁（保证初始化操作只有单个操作）
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}

//实现Once
type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

//基于sync.Once 实现单例模式

var (
	instance2 *singleton
	once      sync.Once
)

func Instance2() *singleton {
	once.Do(func() {
		instance2 = &singleton{}
	})
	return instance2
}

//atomic.Value 复杂对象的原子操作 生产者/消费者模型

var config atomic.Value

func InitCoinfig() {
	config.Store(loadConfig())
	go func() {
		for {
			time.Sleep(time.Second)
			config.Store(loadConfig())
		}
	}()
}
func loadConfig() map[string]string {
	mp := make(map[string]string)
	return mp
}
