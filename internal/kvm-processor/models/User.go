package models

type User struct {
	Id              string
	Name            string
	Surname         string
	VirtualMachines []VirtualMachine
}

func ExampleUser() *User {

	var user = User{
		Id:      "1234567890",
		Name:    "Dmirty",
		Surname: "Nesterov",
	}

	return &user
}
