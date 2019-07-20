/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:main
 * @Description:用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型
 * @Date:2019/7/20 20:30
 * @Version:v1.0
*/
package main

import (
	"fmt"
	"sync"
)

//定义一个栈
type Stack struct {
	Items []int
}

//为了保证数据结构的并发安全，这里定义一个互斥锁
var stackmutex sync.Mutex

func (stack *Stack) New() *Stack {
	return &Stack{[]int{}}
}
func (stack *Stack) PUSH(item int) {
	stackmutex.Lock()
	defer stackmutex.Unlock()
	stack.Items = append(stack.Items, item)
}
func (stack *Stack) POP() int {
	stackmutex.Lock()
	defer stackmutex.Unlock()
	if len(stack.Items) == 0 {
		return -1
	}
	item := stack.Items[len(stack.Items)-1]
	stack.Items = stack.Items[:len(stack.Items)-1]
	return item
}

//判定当前栈是否为空
func (stack *Stack) IsEmpty() bool {
	if len(stack.Items) == 0 {
		return true
	}
	return false
}

//使用两个栈实现队列,队列中的数据为int类型
type Queue struct {
	stack1 Stack
	stack2 Stack
}

//初始化队列
func (queue *Queue) InitQueue() {
	//初始化两个栈
	queue.stack1.New()
	queue.stack2.New()
}

//入队
func (queue *Queue) QueueByStackPUSH(item int) {
	queue.stack1.PUSH(item)
}

//出队
func (queue *Queue) QueueByStackPOP() int {
	if queue.stack2.IsEmpty() {
		if queue.stack1.IsEmpty() {
			return -1
		}
		for !queue.stack1.IsEmpty(){
			queue.stack2.PUSH(queue.stack1.POP())
		}
	}
	return queue.stack2.POP()
}

func main() {
	var queue = Queue{}
	queue.InitQueue()
	queue.QueueByStackPUSH(1)
	fmt.Println(queue.QueueByStackPOP())
	queue.QueueByStackPUSH(2)
	fmt.Println(queue.QueueByStackPOP())
	queue.QueueByStackPUSH(3)
	fmt.Println(queue.QueueByStackPOP())
	fmt.Println(queue.QueueByStackPOP())
}
