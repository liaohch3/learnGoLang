/*
* @Author: liaohch3
* @Date:   2020-04-03 00:26:53
* @Last Modified by:   liaohch3
* @Last Modified time: 2020-04-03 00:55:54
 */

package main

import (
	"fmt"
	"math"
)

func variableZeroValue() {
	var a int
	var s string
	// 默认初始化为0
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 2, 3
	var s string = "liaohch3"
	fmt.Printf("%d %d %q\n", a, b, s)
}

func variableTypeDeduction() {
	// 编译器自动判断类型
	var a, b, s, h = 2, 3, "liaohch3", false
	fmt.Println(a, b, s, h)
}

func variableShorter() {
	// 更短的初始化方法
	a, b, s, h := 2, 3, "liaohch3", false
	b = 8
	fmt.Println(a, b, s, h)
}

// 包内全局变量
var (
	gloA = 3
	gloS = "jkl"
	gloB = true
)

func triangle() {
	// GoLang需要强制类型转换
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(a, b, c)
}

func consts() {
	// 常量
	const filename = "xxx.txt"
	const a, b = 3, 4
	// 常量类似于文本替换，传参可以不用强制类型转换
	var c int = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, a, b, c)
}

func enums() {
	// 枚举
	const (
		cpp = iota
		java
		python
		golang
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {

	// 变量声明
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(gloA, gloS, gloB)

	triangle()
	consts()
	enums()

}
