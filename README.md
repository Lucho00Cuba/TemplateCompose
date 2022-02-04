
### Template Compose

```shell
go run core.go

        Rancher Template Compose - Continues Delivery

        [*] AYUDA

        Uso: core -xml [XML] -yml [YML] -out [OUT]


lucho@cloud:~/00/dev/go$ go run core.go -xml ./compose.xml -yml ./compose.yml -out ./deploy.yml

Parser

 nginx:
   image: nginx:alpine
   environment:
     ENV:prod
     USER:admin
     PASS:admin
   ports:
    - 2200:3300

 nginx-1:
   environment:
     ENV: test
     USER: root
     PASS: root
   image: nginx-1:alpine

```