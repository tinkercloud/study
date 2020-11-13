package main

import "fmt"

type data struct {
	name string
	age  int
}

// return data struct data info
func (d data) DisplayInfo() {
	fmt.Printf("name:%s\tage:%d\n", d.name, d.age)
}

func (d *data) SetAge(newAge int) {
	d.age = newAge
	fmt.Printf("New Age:%d\n", d.age)
}

func (d *data) SetName(newName string) {
	d.name = newName
	fmt.Printf("New Name:%s\n", d.name)
}

func main() {
	var d data

	d.age = 19
	d.name = "hello"

	d.DisplayInfo()
	d.SetAge(22)

	f1 := d.DisplayInfo

	f1()

}
