/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:tree
 * @Description:构建二叉树，并对其进行先序/后序/中序遍历
 * @Date:2019/7/18 17:20
 * @Version:v1.0
*/
package main

import (
	"fmt"
)

type TreeNode struct {
	val       int
	leftNode  *TreeNode
	rightNode *TreeNode
}

//二叉树的序列（可以是先序/后序/中序）
var treearr = []int{1, 2, 3, 4, 0, 5, 6, 7, 8, 9}
//定义一个可供递归时调用的索引变量
var i = -1

func CreateTreeNode(arr []int) *TreeNode {
	i++
	if i >= len(arr) {
		return nil
	}

	var t *TreeNode
	if arr[i]!=0 {
		t = &TreeNode{arr[i], nil, nil}
		t.leftNode = CreateTreeNode(arr)
		t.rightNode = CreateTreeNode(arr)
	}else{
		return nil
	}
	return t
}
//二叉树的先序遍历
func PrevPrintTree(t *TreeNode){
	if t == nil{
		return
	}
	fmt.Print(t.val,"  ")
	PrevPrintTree(t.leftNode)
	PrevPrintTree(t.rightNode)
}

//二叉树的中序遍历
func MidPrintTree(t *TreeNode){
	if t == nil{
		return
	}
	PrevPrintTree(t.leftNode)
	fmt.Println(t.val,"  ")
	PrevPrintTree(t.rightNode)
}
func main() {
	t := CreateTreeNode(treearr)
	PrevPrintTree(t)
	fmt.Println()
	MidPrintTree(t)
}
