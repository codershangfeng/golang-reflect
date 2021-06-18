package main

import (
	"fmt"
)

func main() {
	var u1 User
	u2 := u1
	fmt.Printf("user1: %#v\n user2:%#v\n", u1, u2)
	var salary int32 = 10000
	u1.salary = &salary
	u1.age++
	fmt.Printf("user1: %#v\n user2:%#v\n", u1, u2)
}

type User struct {
	age    int32
	salary *int32
}
