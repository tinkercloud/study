package main

import "fmt"

type user struct {
	name    string
	surname string
}

func main() {

	users1 := make(map[string]user)

	users1["Roy"] = user{"Rob", "Roy"}

	users1["Ford"] = user{"Henry", "Ford"}

	users1["Mouse"] = user{"Mickey", "Mount"}

	for k, v := range users1 {
		fmt.Println(k, v)
	}
}
