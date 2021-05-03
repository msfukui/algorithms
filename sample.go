package main

import "fmt"

func main() {

	A := []string{"a", "b", "c"}
	B := []string{"a", "c", "d"}
	C := []string{}

	fmt.Printf("A = %v, B = %v, C = %v\n", A, B, C)

	merge_flag := true

	for _, a := range A {
		for _, b := range B {
			if a == b {
				merge_flag = false
				break
			}
		}
		if merge_flag {
			C = append(C, a)
		}
		merge_flag = true
	}
	fmt.Printf("C = %v\n", C)
}
