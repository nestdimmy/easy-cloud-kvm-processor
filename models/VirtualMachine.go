package models

type VirtualMachine struct {
	Id       *int
	UUID     string
	Name     string
	HostName string
	Owner    *User
}

func NewVirtualMachine(id *int, UUID string, name string, hostname string, owner *User) *VirtualMachine {
	vm := new(VirtualMachine)
	vm.Name = name
	vm.Id = id
	vm.UUID = UUID
	vm.HostName = hostname
	vm.Owner = owner

	return vm
}
