package dto

type DomainDto struct {
	Name          string
	Id            uint32
	UUID          string
	Memory        int64
	CurrentMemory int64
	VirtualCPUs   int64
}
