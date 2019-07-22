/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:main
 * @Description:我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？
 * 破题：我们可以先定义一个2*8的矩形，这样使用2*1的矩形有两种方法放置：
 * 1.横着放：此时下方必须放置一个2*1的小矩形，此时变换为2*6的矩形，此时有f（6）种放置方法
 * 2.竖着放：此时变换为2*7的矩形，此时有f（7）种解法
 * 所以f（8）=f(7)+f(6)
 * @Date:2019/7/21 0:41
 * @Version:v1.0
*/
package main

import "fmt"

func MatrixSlov(n int) int {
	if n <= 0 {
		return -1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	var sum, temp1, temp2 int = 0, 1, 1
	for i := n; i > 2; i-- {
		sum = temp1 + temp2
		temp1 = temp2
		temp2 = sum
	}
	return sum
}
func main() {
	fmt.Print(MatrixSlov(4))
}
