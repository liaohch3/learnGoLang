/*
* @Author: liaohch3
* @Date:   2020-04-03 14:50:43
* @Last Modified by:   liaohch3
* @Last Modified time: 2020-04-03 20:51:22
 */

package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("Bad operator: %s", op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	q, r = a/b, a%b
	return
}

// 函数式编程
func apply(op func(int, int) int, a, b int) int {
	opName := runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name()
	fmt.Println("Calling %s with args %d, %d\n", opName, a, b)
	return op(a, b)
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0

	// 此处的i是下标，不是item值
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	if res, err := eval(14, 5, "/"); err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}

	if res, err := eval(14, 5, "%"); err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}

	fmt.Println(div(13, 3))
	fmt.Println(div2(13, 3))

	fmt.Println(apply(
		// 传入匿名函数
		func(a, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(5, 4, 3, 2, 1))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

}
