package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")
	hanio(3, "START POLE", "MIDDLE POLE", "END POLE")
}

func hanio(n int, a string, b string, c string) {
	if n == 1 {
		fmt.Printf("move %d from %s to %s\n", n, a, c)
		return
	}

	hanio(n-1, a, c, b)
	fmt.Printf("move %d from %s to %s\n", n, a, c)
	hanio(n-1, b, a, c)
}
