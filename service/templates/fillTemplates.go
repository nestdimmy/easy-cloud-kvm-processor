package templates

import (
	"../../models"
)

func FillVirtualMachineTemplate(vmBody models.VirtualMachine) string {
	return `
<domain type="kvm">
	<name>` + vmBody.Name + `</name>
	<memory>512000</memory>
	<currentMemory>512000</currentMemory>
	<vcpu>1</vcpu>
	<os>
		<type arch="x86_64">hvm</type>
		<boot dev="cdrom"/>
	</os>
	<features>
		<acpi/>
		<apic/>
	</features>
	<cpu mode="custom" match="exact">
		<model>Broadwell-noTSX-IBRS</model>
	</cpu>
	<clock offset="utc">
		<timer name="rtc" tickpolicy="catchup"/>
		<timer name="pit" tickpolicy="delay"/>
		<timer name="hpet" present="no"/>
	</clock>
	<pm>
		<suspend-to-mem enabled="no"/>
		<suspend-to-disk enabled="no"/>
	</pm>
	<devices>
		<emulator>/usr/bin/kvm-spice</emulator>
		<disk type="file" device="cdrom">
			<driver name="qemu" type="raw"/>
			<source file="/home/dmitry/GolandProjects/kvm-processor/kvm/imgs/ubuntu.iso"/>
			<target dev="hda" bus="ide"/>
			<readonly/>
		</disk>
		<controller type="usb" index="0" model="ich9-ehci1"/>
		<controller type="usb" index="0" model="ich9-uhci1">
			<master startport="0"/>
		</controller>
		<controller type="usb" index="0" model="ich9-uhci2">
			<master startport="2"/>
		</controller>
		<controller type="usb" index="0" model="ich9-uhci3">
			<master startport="4"/>
		</controller>
		<interface type="network">
			<source network="default"/>
			<mac address="52:54:00:be:21:bc"/>
			<model type="virtio"/>
		</interface>
		<input type="tablet" bus="usb"/>
		<graphics type="vnc" port="-1"/>
		<console type="pty"/>
	</devices>
</domain>`
}