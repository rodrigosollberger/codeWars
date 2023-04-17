//Given a string of digits, you should replace any digit below 5 with '0' and
//any digit 5 and above with '1'. Return the resulting string.

package main

import "fmt"

func main() {
	FakeBin("7684")
	fmt.Println(FakeBin("7684"))
}

func FakeBin(x string) string {
	// your code here
	y := ""
	for _, v := range x {
		if v < 5 {
			y = "0" + y
		} else {
			y = "1" + y
		}
	}
	return y
}
