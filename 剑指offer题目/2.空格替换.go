/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:main
 * @Description:请实现一个函数，将一个字符串中的每个空格替换成“%20”。例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
 * @Date:2019/7/18 14:36
 * @Version:v1.0
*/
package main

import (
	"fmt"
	"strings"
)

//golang 中的字符串库实现，不推荐
func replaceBlank1(s string) string {
	//sarr := strings.Split(s, " ")
	//return strings.Join(sarr,"%20")
	return strings.Replace(s, " ", "%20", -1)

}

/**正常方法：
***正序遍历，如果碰到空格就替换，因为替换字符长度为3，所以，所有数据都需要后移，时间复杂度为O（n^2）
**/
/**
	考虑牺牲时间换空间的方法：倒序添加
	1.将字符串转换为字符切片
	2.遍历字符切片统计空格的个数
    3.准备两个指针，一个指向切片末尾，一个指向追加空间的末尾（追加空格字符空间至切片末尾）
	4.一个指针指向切片末尾，一个指针指向追加末尾，将切片末尾的数据拷贝至追加末尾
	5.如果遇到空格，切片指针直接减1，追加指针减3，并追加替换字符
**/

func replaceBlank(s string) string {
	//将字符串转换为字符切片
	sarr := []byte(s)
	var blankcount = 0
	for _, sbyte := range sarr {
		if sbyte == 32 {
			blankcount++
		}
	}
	sarr = append(sarr, make([]byte, blankcount*2)...)
	pointsarr := len(s) - 1
	pointappend := len(sarr) - 1
	for i := 0; i < len(s); i++ {
		if sarr[pointsarr] == 32 {
			pointsarr--
			pointappend -= 3
			sarr[pointappend+1] = '%'
			sarr[pointappend+2] = '2'
			sarr[pointappend+3] = '0'
		} else {
			sarr[pointappend] = sarr[pointsarr]
			pointappend--
			pointsarr--
		}
	}
	return string(sarr)
}
func main() {
	s := "I am a student"
	fmt.Println(replaceBlank(s))
}
