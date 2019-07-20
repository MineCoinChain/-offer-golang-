/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:main
 * @Description:使用go语言实现栈
 * @Date:2019/7/20 20:56
 * @Version:v1.0
*/
package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	Items []interface{}
}
//为了保证数据结构的并发安全，这里定义一个互斥锁
var stackmutex sync.Mutex
func (stack *Stack) New() *Stack {
	return &Stack{[]interface{}{}}
}
func (stack *Stack) PUSH(item interface{}) {
	stackmutex.Lock()
	stack.Items = append(stack.Items, item)
	stackmutex.Unlock()
}
func (stack *Stack) POP() interface{} {
	stackmutex.Lock()
	if len(stack.Items)==0{
		return nil
	}
	item := stack.Items[len(stack.Items)-1]
	stack.Items = stack.Items[:len(stack.Items)-1]
	stackmutex.Unlock()
	return item
}
func main() {
	var stack Stack
	stack.New()
	stack.PUSH(1)
	stack.PUSH(2)
	stack.PUSH(3)
	fmt.Println(stack.POP())
	fmt.Println(stack.POP())
	fmt.Println(stack.POP())
	fmt.Println(stack.POP())
}
