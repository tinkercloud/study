package main

import "fmt"

type notifier interface {
	notify()
}

type printer interface {
	print()
}

type user struct {
	name  string
	email string
}

func (u user) print() {
	fmt.Printf("Print User:%s, Email:%s\n", u.name, u.email)
}

func (u *user) notify() {
	fmt.Printf("User:%s, Email:%s\n", u.name, u.email)
}

func main() {

	u := user{"wang", "tinker@gmail.com"}

	// SendNotification(&u)

	enters := []printer{
		u,
		&u,
	}

	u.name = "ak47"
	u.email = "tinker101@gmail.com"

	//
	for _, e := range enters {
		fmt.Println(e)
	}
}

// 将接口传递给函数，并实现接口中实现的notify函数调用
func SendNotification(f notifier) {
	f.notify()
}
