package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	//1、sruct as a value type
	var pers1 Person
	pers1.firstName = "liao"
	pers1.lastName = "bei"
	upPerson(&pers1)
	fmt.Printf("the name person firstname is %s and lastname is %s\n", pers1.firstName, pers1.lastName)
	//2、sruct as a pointer
	pers2 := new(Person)
	(*pers2).firstName = "ye"
	(*pers2).lastName = "chuanrong"
	upPerson(pers2)
	fmt.Printf("the name person firstname is %s and lastname is %s\n", pers2.firstName, pers2.lastName)
	//3、sruct as a literal
	pers3 := &Person{"liao", "chuan"}
	upPerson(pers3)
	fmt.Printf("the name person firstname is %s and lastname is %s\n", pers3.firstName, pers3.lastName)
}
