package main

import (
	"fmt"
	"strconv"
)

// 参考资料：https://www.runoob.com/w3cnote/bit-operation.html
//			https://www.runoob.com/go/go-operators.html

//左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运
//算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。
//
//右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运
//算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。

// 将十进制数字转化为二进制字符串
func convertToBin(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return s
}

func main() {
	/*	A =    0011 1100

		B =    0000 1101

		A&B =  0000 1100
		A|B =  0011 1101
		A^B =  0011 0001*/

	var a uint = 60 /* 60 = 0011 1100 */
	//var b uint = 13      /* 13 = 0000 1101 */

	fmt.Println(a << 3)
	fmt.Println(a >> 3)

}
