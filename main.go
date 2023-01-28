package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
)

type Person struct{
	UserName string `faker:"username"`
	FirstName string `faker:"first_name"`
	LastName string `faker:"last_name"`
	Email string `faker:"email"`
	PhoneNumber string `faker:"phone_number"`
}

func main() {
	person1 := Person{}
	err := faker.FakeData(&person1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", person1)
}
