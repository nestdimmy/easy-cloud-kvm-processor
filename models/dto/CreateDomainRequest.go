package dto

type CreateDomainDto struct {
	Name          string
	Memory        int64
	CurrentMemory int64
	VirtualCPUs   int64
}
