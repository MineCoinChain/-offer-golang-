/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:offer
 * @Description:在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，
 *每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
 * @Date:2019/7/18 13:30
 * @Version:v1.0
*/
package main

import "fmt"

var array = [5][5]int{
	{1, 2, 3, 4, 5},
	{2, 3, 4, 5, 6},
	{3, 4, 5, 6, 7},
	{4, 5, 6, 7, 8},
	{5, 6, 7, 8, 9},
}

func FindX(x int) bool {
	//使用二分查找的思想，找到斜对角线的元素,如果比斜对角线元素大往下找，如果比斜对角线元素小往左找;
	var h, l = 0,len(array[0])-1
	for {
		middle := array[h][l]
		if x > middle {
			if h == len(array)-1 {
				break
			}
			h = h + 1
		}
		if x < middle {
			if l == 0 {
				break
			}
			l = l - 1
		}
		if x == middle {
			return true
		}

	}
	return false
}
func main() {
	fmt.Println(FindX(10))
}
