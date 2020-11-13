package main

import "fmt"

type user struct {
	name  string
	email string
}

// fmt.Printf("Sending User Emain:%s, User:%s",u.name,u.email)
// print current email or users
func (u user) notify() {
	fmt.Printf("Sending Emain:%s, User:%s\n", u.email, u.name)
}

// change email func
func (u *user) changeEmail(email string) {
	u.email = email
	fmt.Printf("New Email:%s\n", email)
}

// 更改用户实现
func (u *user) changeUser(user string) {
	u.name = user
	fmt.Printf("New UserName:%s\n", user)

}

// func main() {

// 	bill := user{"bill", "worss"}

// 	bill.notify()

// 	bill.changeEmail("ak47")

// 	bill.notify()

// }
