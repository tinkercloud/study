package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("User:%s, Email:%s\n", u.name, u.email)
}

type notify struct {
	user
	url string
}

func main() {

	n := notify{
		user: user{
			"wang",
			"tk@qq.com",
		},
		url: "www.baidu.com",
	}

	n.user.notify()
	fmt.Println(n.url)
}
