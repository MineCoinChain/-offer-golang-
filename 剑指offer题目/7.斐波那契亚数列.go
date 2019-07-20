/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:main
 * @Description:大家都知道斐波那契数列，现在要求输入一个整数n，请你输出斐波那契数列的第n项
 * 斐波那契亚数列通常使用递归解法，但考虑到求解的重复性过高，因此面试时不要使用递归直接使用循环实现
 * @Date:2019/7/20 23:54
 * @Version:v1.0
*/
package main

import "fmt"

func Fabcci(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var temp1 = 0
	var temp2 = 1
	var sum int
	for i := n; i > 2; i-- {
		sum = temp1 + temp2
		temp1 = temp2
		temp2 = sum
	}
	return sum
}
func main() {
	fmt.Println(Fabcci(10))
	fmt.Println(Fabcci(3))
}


