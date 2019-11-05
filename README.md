KVM-processor

Project develops on Ubuntu 18.04 

[libvirt package documentation](https://libvirt.org/html/index.html)

1. Install kvm libraries
 - use - `sudo apt-get install qemu-kvm libvirt-bin ubuntu-vm-builder bridge-utils`
 - type `kvm-ok` to understand that kvm is working  
    * if you see smth like `Your cpu does not support KVM extentions`  
          go to BIOS and turn on Inter/AMD Virtualization  
 
 __Remember VirtualBox or Parallels for MacOs can't do parallel virtualization, 
 so you need install ubuntu or other Linux based instance to work properly__
 
 _if you see this, you are ready to start using KVM_
 ```
 $ kvm-ok  
INFO: /dev/kvm exists 
KVM acceleration can be used
 