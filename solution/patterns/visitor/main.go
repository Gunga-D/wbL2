package main

import "fmt"

type Action interface {
}

type Privilege interface {
	GetName() string
	IsActionValid(act Action) bool
	Process(pp PrivilegeProcess)
}

type User struct {
}

func (u User) GetName() string {
	return "User"
}

func (u User) IsActionValid(act Action) bool {
	return false
}

type Admin struct {
}

func (a Admin) GetName() string {
	return "Admin"
}

func (u Admin) IsActionValid(act Action) bool {
	return false
}

type PrivilegeProcess interface {
	ProcessUser()
	ProcessAdmin()
}

type PrivilegeCalculator struct {
}

func (pc PrivilegeCalculator) ProcessUser() {
	fmt.Println("Calculate user...")
}

func (pc PrivilegeCalculator) ProcessAdmin() {
	fmt.Println("Calculate admin...")
}
