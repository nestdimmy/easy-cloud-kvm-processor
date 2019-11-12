package models

type VirtualMachine struct {
	Id       string
	Name     string
	HostName string
	Owner    User
}
