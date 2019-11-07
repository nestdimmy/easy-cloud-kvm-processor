package models

type VirtualMachine struct {
	Id    string
	Name  string
	Mask  string
	Owner User
}
