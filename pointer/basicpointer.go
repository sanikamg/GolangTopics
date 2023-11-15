package main

import "fmt"

//when we need update a state
//pointer = 8 byte
// when we want to optimize the memory for large objects that are getting called alot.

type User struct {
	username string
	email    string
	age      int
}

func (u User) Email() string {
	return u.email
}

func (u *User) UpdateEmail(email string) {
	u.email = email
}
func main() {
	user := User{
		email: "sanik@gamil.com",
	}
	user.UpdateEmail("hiii@.com")
	fmt.Println(user.Email())
}
