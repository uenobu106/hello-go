package main

import "fmt"

type Student struct {
	name string
}


type User struct {
	gender string
	age int
}

func (s Student) avg(math, english float64) float64 {
	return (math + english)/2
}

func main() {
	user1 := User{"hanako", 20}
	fmt.Println(user1)
	var user2 User
	user2.gender = "f"
	user2.age = 20
	fmt.Println(user2)

	hoge001 := Student{"sito"}
	hoge001.avg(79,90)
	fmt.Println(hoge001.avg(79,90))

}