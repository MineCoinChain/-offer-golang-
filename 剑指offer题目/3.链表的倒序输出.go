/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:main
 * @Description:输入一个链表，按链表值从尾到头的顺序输出
 * @Date:2019/7/18 16:43
 * @Version:v1.0
*/
package main

import "fmt"

//这里简单起见直接定义一个值为int类型的链表
type List struct {
	Val  int
	Next *List
}

//添加一个链表的节点
func (list *List) AddNode(val int) {
	for list.Next != nil {
		list = list.Next
	}
	list.Next = &List{val, nil}
}

//倒序打印链表,这里采用递归的方法实现，缺点：链表过长会造成内存溢出
func reverPrintList(list *List) {
	if list == nil {
		return
	}
	reverPrintList(list.Next)
	fmt.Println(list.Val)
}
func main() {
	//定义一个头节点
	Head:=List{10,nil}
	Head.AddNode(9)
	Head.AddNode(8)
	Head.AddNode(7)
	Head.AddNode(6)
	Head.AddNode(5)
	Head.AddNode(4)
	Head.AddNode(3)
	Head.AddNode(2)
	Head.AddNode(1)
	reverPrintList(&Head)
}
