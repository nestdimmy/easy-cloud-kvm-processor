package models

type VirtualMachine struct {
	Id       *int
	UUID     string
	Name     string
	State    string
	HostName string
	OwnerId  string
}

func VM(id *int, UUID string, name string, state string, hostname string, owner string) *VirtualMachine {
	vm := new(VirtualMachine)
	vm.Name = name
	vm.Id = id
	vm.UUID = UUID
	vm.State = state
	vm.HostName = hostname
	vm.OwnerId = owner

	return vm
}
