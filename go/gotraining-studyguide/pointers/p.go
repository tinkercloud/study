package main

import "fmt"

func main() {

	slice1 := make([]string, 4)

	slice1[0] = "apple"
	slice1[1] = "green"
	slice1[2] = "red"
	slice1[3] = "xxxxx"

	slice1 = append(slice1, "kk")
	slice1 = append(slice1, "bb")
	slice1 = append(slice1, "dd")
	slice1 = append(slice1, "ee")
	slice1 = append(slice1, "ss")
	fmt.Println(cap(slice1), len(slice1), slice1)
}

func printV1(v int) {
	v++

	fmt.Printf("value:%d , Point: %p \n", v, &v)
}

func printV2(v *int) {

	// *v++

	fmt.Printf("value:%d , Point: %p, old value Point:%p", *v, &v, v)
}
