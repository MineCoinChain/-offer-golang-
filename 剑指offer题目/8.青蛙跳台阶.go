/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:main
 * @Description:一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法
 * 递推公式：f(n)=f(n-1)+f(n-2),本质是一个斐波那契亚数列
 * @Date:2019/7/21 0:15
 * @Version:v1.0
*/
package main

import "fmt"

func Jump(n int)int{
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
func main(){
	fmt.Println(Jump(12))
}