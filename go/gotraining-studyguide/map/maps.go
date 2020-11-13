package main

type user struct {
	name    string
	surname string
}

func main() {

	users1 := make(map[string]user)

	users1["abb"] = user{"hello", "world"}

	
}
