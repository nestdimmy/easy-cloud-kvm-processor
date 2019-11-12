#KVM-processor

####Project develops on Ubuntu 18.04 

[libvirt package documentation](https://godoc.org/github.com/libvirt/libvirt-go)

#####__TODO list__:
* [ ] Scruct README.md
* [x] Create project skeleton
* [x] Run VM with KVM in CLI mod
* [ ] Create .sh script for fast start (to update and install all needed packages and libs)
* [ ] Run VM with libvirt-go library   


#####Установка библиотек KVM
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
#####Список используемых пакетов:
 - `gorilla/mux` — URL-маршрутизатор и диспетчер.  
 Мы будем использовать этот пакет для сопоставления путей URL с их обработчиками;
 - `jinzhu/gorm` — ORM-библиотека для Golang.  
 Мы используем этот пакет, чтобы взаимодействовать с базой данных;
 - Для логирования используем `google/logger`;
 
   
     
 _________________________
 #####Тестовый запуск
 Запускаем скрипт createVM.sh, пока программное создание ВМ не готово.  
 Убеждаемся в отсутсвии ошибок. При успешном поднятии должно получиться что-то вроде:
 
 ````
[Tue, 12 Nov 2019 14:35:45 virt-install 10081] DEBUG (cli:451) Launching virt-viewer for graphics type 'vnc'
[Tue, 12 Nov 2019 14:35:45 virt-install 10081] DEBUG (cli:425) Running: virt-viewer --connect qemu:///system --wait demo
[Tue, 12 Nov 2019 14:38:14 virt-install 10081] DEBUG (virt-install:744) Domain state after install: 1
[Tue, 12 Nov 2019 14:38:16 virt-install 10081] DEBUG (virt-install:744) Domain state after install: 1
Domain creation completed.
````
 Далее запускаем наше go приложение  
 делаем Get запрос на "localhost:8080/api/v1/vm/test"
 
 в ответ мы должны получить имя созданного домена
 
 ```
{
    "domainName": "demo"
}
```
 В консоли же отображается ID, Name, UUID того же домена
 
 ````
 Version: 4.0.0
 ID      Name            UUID
 --------------------------------------------------------
 9       demo    b61295a01ffa4012a37b23c61805d009

 ````
 
 #####Полезные команды для управления инстансами KVM:

Синтаксис выключения ВМ:  
* `virsh shutdown domain`  
* `virsh shutdown vm`  
* `virsh shutdown freebsd`  
* `virsh shutdown ubuntu1`  

Синтаксис удаления инстансов:  
* `virsh destroy domain`  
* `virsh destroy vm`  
* `virsh destroy freebsd`  
* `virsh destroy ubuntu1`  