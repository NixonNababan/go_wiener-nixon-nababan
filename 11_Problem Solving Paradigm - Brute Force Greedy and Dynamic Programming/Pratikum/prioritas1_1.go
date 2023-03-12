package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2
	ans := make([]string, x+1)
	for i := 0; i <= x; i++ {
		ans[i] = strconv.FormatInt(int64(i), 2)
	}
	fmt.Println(ans)
}
