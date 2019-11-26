package defaults

/*
	Operational System
*/
const OsTypeArch string = "x86_64"
const OsBootDev string = "cdrom"
const OsTypeDefault string = "hvm"

/*
	CPU
*/
const CpuMode string = "custom"
const CpuMatch string = "exact"
const CpuModel string = "Broadwell-noTSX-IBRS"

/*
	Domain
*/
const DomainType string = "kvm"
const DomainDefaultName string = "newDomain"
const DomainMinMemory int64 = 512000
const DomainVirtualCpusDefaultCount int64 = 1

/*
	Disk
*/
const DiskType string = "file"
const DiskDevice string = "cdrom"
const DiskDriverName string = "qemu"
const DiskDriverType string = "raw"
const DiskSourceFile string = "/home/dmitry/GolandProjects/kvm-processor/kvm/imgs/ubuntu.iso"
const DiskTargetDev string = "hda"
const DiskTargetBus string = "ide"

/*
	Network interface
*/
const InterfaceType string = "network"
const InterfaceSourceNetwork string = "default"
const InterfaceModelType = "virtio"

/*
	Devices
*/
const DeviceInputType string = "tablet"
const DeviceInputBus string = "usb"
const DeviceGraphicsType string = "vnc"
const DeviceGraphicsPort string = "-1"
const DeviceConsoleType string = "pty"
const DeviceEmulator string = "/usr/bin/kvm-spice"
