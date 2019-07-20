/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:stack__queue
 * @Description:使用go语言实现一个栈
 * @Date:2019/7/20 20:36
 * @Version:v1.0
*/
package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	items []interface{}
}
//为了保证并发安全这里最好定义一个互斥锁
var queuemutex sync.Mutex
//构造一个队列
func (queue *Queue) New() *Queue {
	return &Queue{[]interface{}{}}
}

//入队
func (queue *Queue) PUSH(item interface{}) {
	queuemutex.Lock()
	queue.items = append(queue.items, item)
	queuemutex.Unlock()
}

//出队
func (queue *Queue) POP() interface{} {
	queuemutex.Lock()
	if len(queue.items)==0{
		return nil
	}
	item := queue.items[0]
	queue.items = queue.items[1:]
	queuemutex.Unlock()
	return item
}
func main() {
	var queue Queue
	queue.New()
	queue.PUSH(1)
	queue.PUSH(2)
	queue.PUSH(3)
	fmt.Println(queue.POP())
	fmt.Println(queue.POP())
	fmt.Println(queue.POP())
	fmt.Println(queue.POP())
}

