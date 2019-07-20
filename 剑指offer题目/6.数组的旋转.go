/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:main
 * @Description:把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
 *              输入一个非减排序的数组的一个旋转，输出旋转数组的最小元素。
 *				例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。
 *              NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。
 *              解法：通过二分查找法破题
 * @Date:2019/7/20 22:39
 * @Version:v1.0
*/
package main

import (
	"fmt"
)

//定义一个旋转数组
var xarray = []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4}

func minNumberInArray(array []int) int {
	//这里加上异常判定，如果数组不存在直接返回
	if len(array) == 0 {
		return -1
	}
	//简单起见不考虑重复情况，如果旋转0个元素直接返回
	if xarray[0] < xarray[len(xarray)-1] {
		return xarray[0]
	}
	//使用二分查找对数组进行查找
	var high = len(xarray) - 1
	var low = 0
	//取出中点值

	for high-low != 1 {
		var mid = (high + low) / 2
		if xarray[low] >= xarray[mid] {
			high = mid
		} else {
			low = mid
		}
	}
	return xarray[high]
}
func main() {
	mid := minNumberInArray(xarray)
	fmt.Println(mid)
}
