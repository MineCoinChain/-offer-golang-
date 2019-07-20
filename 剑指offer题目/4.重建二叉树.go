/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:main
 * @Description:输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
 * 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
 * @Date:2019/7/18 17:05
 * @Version:v1.0
*/
package main

import "fmt"

//定义二叉树
type TreeNode struct {
	val       int
	leftNode  *TreeNode
	rightNode *TreeNode
}

//定义二叉树的先序以及中序序列
var preNodeOrder = []int{1, 2, 4, 7, 3, 5, 6, 8}
var midNodeOrder = []int{4, 7, 2, 1, 5, 3, 8, 6}

//重建二叉树的算法
func reConstructBinaryTree(pre []int, mid []int) *TreeNode {
	fmt.Println(pre)
	fmt.Println(mid)
	fmt.Println()
	if len(pre) == 0 || len(mid) == 0 {
		return nil
	}
	//查找当前二叉树的根节点
	var node = TreeNode{pre[0], nil, nil}

	for i := 0; i < len(mid); i++ {
		if node.val == mid[i] {
			node.leftNode = reConstructBinaryTree(pre[1:i+1], mid[:i])
			node.rightNode = reConstructBinaryTree(pre[i+1:], mid[i+1:])
		}
	}
	return &node
}

//二叉树的先序遍历
func PrintTreePre(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val," ")
	PrintTreePre(node.leftNode)
	PrintTreePre(node.rightNode)
}
func main() {
	node := reConstructBinaryTree(preNodeOrder, midNodeOrder)
	PrintTreePre(node)
}
