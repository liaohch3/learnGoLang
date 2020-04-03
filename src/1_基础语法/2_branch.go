/*
* @Author: liaohch3
* @Date:   2020-04-03 00:57:21
* @Last Modified by:   liaohch3
* @Last Modified time: 2020-04-03 14:21:47
 */

package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Bad score: %d", score))
	case score < 60:
		g = "C"
	case score < 80:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "file/lhc.txt"
	// contents, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		// grade(-15),
		grade(0),
		grade(46),
		grade(75),
		grade(89),
		grade(100),
	)
	// fmt.Println(
	// 	grade(-15),
	// )
}
