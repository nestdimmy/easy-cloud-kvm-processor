KVM-processor

Project develops on Ubuntu 18.04 

[libvirt package documentation](https://libvirt.org/html/index.html)

1. Установка библиотек KVM
 - use - `sudo apt-get install qemu-kvm libvirt-bin ubuntu-vm-builder bridge-utils`
 - type `kvm-ok` to understand that kvm is working  
    * Если вы видите это - `Your cpu does not support KVM extentions`, то у вас отключена виртуализация в BIOS, исправте это
     
 __Внимание! VirtualBox или Parallels для MacOs не поддерживают параллельную виртуализацию, т.е. не
 могут запусать вируталку в виртуалке. Для работы вам необходимо будет установиь любой Linux-based дистрибутив 
 с поддержкой KVM__
 
 _Если `kvm-ok` выводит это, значит вы все настроили правильно и готовы к работе с гипервизором_
 ```
 $ kvm-ok  
INFO: /dev/kvm exists 
KVM acceleration can be used
```
Для загрузки необходимых пакетов прописываем в териминале 
```
go get github.com/{package-name}
```
Список используемых пакетов:
 - `gorilla/mux` — URL-маршрутизатор и диспетчер.  
 Мы будем использовать этот пакет для сопоставления путей URL с их обработчиками;
 - `jinzhu/gorm` — ORM-библиотека для Golang.  
 Мы используем этот пакет, чтобы взаимодействовать с базой данных;
 - Для логирования используем `google/logger`;
 
 