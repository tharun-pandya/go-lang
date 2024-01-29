package main

import "fmt"

func main() {
	var r1, r2 float32
	func() {
		fmt.Println("Enter first Ratio number")
		fmt.Scan(&r1)
		fmt.Println("Enter second Ratio number")
		fmt.Scan(&r2)
		r1 = (r1 + r2) / r1
		r2 = (r1 + r2) / r2
		fmt.Printf("%+2v ratio1 \n", r1)
		fmt.Printf("%+2v ratio2 ", r2)
	}()
}
