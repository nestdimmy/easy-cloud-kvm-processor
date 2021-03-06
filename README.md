KVM-processor

Project develops on Ubuntu 18.04 

[Организация кода](https://golang.org/doc/code.html)

[libvirt package documentation](https://godoc.org/github.com/libvirt/libvirt-go)

[Создание автоустоновочного образа ubuntu](https://habr.com/ru/post/104029/) 

__TODO list__:
* [ ] Scruct README.md
* [x] Create project skeleton
* [x] Run VM with KVM in CLI mod
* [ ] Create .sh script for fast start (to update and install all needed packages and libs)
* [x] Run VM with libvirt-go library   
* [x] Delete VM with libvirt-go library   
* [ ] Run to go-dep project structure

Установка библиотек KVM
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
 - https://libvirt.org/formatdomain.html - подробное описание xml конфигураций 
 для создания виртуальных машин с помощью libvirt
 
  
Тестовый запуск  

 Создание домена
 POST localhost:8080/api/v1/vm with body
  
  ```json
{
	"Name": "testDomain" 
}
```

  return  
  ```json
{
    "Id": 4,
    "UUID": "1fb62aad-8e82-44ee-ab1c-e9c4dcf48938",
    "Name": "testDomain",
    "HostName": "",
    "Owner": {
        "Id": "1234567890",
        "Name": "Dmirty",
        "Surname": "Nesterov",
        "VirtualMachines": null
    }
}
```
Удаление  
 REST - DELETE  
1. localhost:8080/api/v1/vm/4
2. localhost:8080/api/v1/vm/1fb62aad-8e82-44ee-ab1c-e9c4dcf48938 
3. localhost:8080/api/v1/vm/testDomain
 
 
 Получение домена
 REST - GET  
1. localhost:8080/api/v1/vm/4
2. localhost:8080/api/v1/vm/1fb62aad-8e82-44ee-ab1c-e9c4dcf48938 
3. localhost:8080/api/v1/vm/testDomain
 
return  
   ```json
 {
     "Id": 4,
     "UUID": "1fb62aad-8e82-44ee-ab1c-e9c4dcf48938",
     "Name": "testDomain",
     "HostName": "",
     "Owner": {
         "Id": "1234567890",
         "Name": "Dmirty",
         "Surname": "Nesterov",
         "VirtualMachines": null
     }
 }