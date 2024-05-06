package main

import "fmt"

type User struct {
	Name string
}

type Account struct {
	Name   string
	Age    int
	Adress string
	User   // legacy methods, legacy fields,have hierarchy
}

func main() {
	// var Adlet User = User{
	// 	Name: "AK",
	// }
	var Didar Account = Account{
		Name: "Didar",
		Age:  19,
		User: User{
			Name: "Amirlan",
		},
	}

	// var ptr *User = &Adlet
	// (ptr.Name) = "Amirlan"
	Didar.setname("habib") // more priority
	Didar.User.setname("AK")
	fmt.Println(Didar)
	Didar.updatename2("ggg") // legacy
	fmt.Println(Didar)
}

func (p *Account) setname(s string) {
	p.Name = s
}

func (p *User) setname(s string) {
	p.Name = s
}

func (p *User) updatename2(s string) {
	p.Name = s
}
