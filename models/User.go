package models

type User struct {
	Id              string
	Name            string
	VirtualMachines []VirtualMachine
}
