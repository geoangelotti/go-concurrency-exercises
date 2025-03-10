package main

import "fmt"

type user struct {
	name  string
	email string
}

// TODO: Implement custom formating for user struct values.
func (u user) String() string {
	return fmt.Sprintf("%s <%ss>", u.name, u.email)
}

func main() {
	u := user{
		name:  "John Doe",
		email: "johndoe@example.com",
	}
	fmt.Println(u)
}
