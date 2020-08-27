package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	x1, x2 := 0, 0

	return func() int {

		result := x1 + x2
		x1 = x2
		if x2 == 0 {
			x2 = 1
		} else {
			x2 = result
		}

		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
