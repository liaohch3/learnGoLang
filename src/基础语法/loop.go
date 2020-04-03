/*
* @Author: liaohch3
* @Date:   2020-04-03 14:26:13
* @Last Modified by:   liaohch3
* @Last Modified time: 2020-04-03 14:50:37
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {

	if n == 0 {
		return "0"
	}

	res := ""
	for ; n > 0; n /= 2 {
		res = strconv.Itoa(n%2) + res
	}
	return res
}

func printFile(filename string) {
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

func forever() {
	for {
		fmt.Println("kjl")
	}
}

func main() {
	fmt.Println(
		convertToBin(0),
		convertToBin(5),
		convertToBin(14),
		convertToBin(84),
	)

	printFile("file/lhc.txt")

	// forever()
}
