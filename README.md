# GOReactMongo

## Iniciamos el proyecto go con el comando

- go mod init ES EQUIVALENTE en node a npm int  sin embargo nos pide ingresar la direccion de un repositorio donde se va alojar
```
go mod init github.com/ivanes285/GOReactMongo  
```

### Instalacion de modulos

- Instalamos un framework llamado fiber equivalente(express en nodejs) para trabajar con go
```
go get github.com/gofiber/fiber/v2
```
- Importamos Cors para poder habilitarlos para el frontend
```
 "github.com/gofiber/fiber/v2/middleware/cors"
```
- Para evitar parar el servidor y de nuevo ejecutarlo para aplicar los cambios se instalo el siguiente modulo equivalente (nodemon en nodejs)
 ```
 go install github.com/cosmtrek/air@latest
 ```
### Ejecucion
- Podemos compilar de forma normal con el comando 
```
go run .
```
o tambien
```
 go run .\main.go
```

## Despliegue 
- Desplegamos en railway a travez de un repositorio de github tomando en cuenta que el static file en este caso dist se 
 encuentra en el mismo nivel de la carperta server. 
- Para usar statics files en fiber realizamos la siguiente configuracion en el main.go
 ```
  app.Static("/", "../dist") 
 ``` 
- Nota: Por alguna razón que aun desconozco no se pudo desplegar en railway tomando en cuenta que el frontend(folder client) esta en el mismo nivel de la carpeta server, sin embargo se acostumbra a trabajar por separado el frontend y el backend. Por lo tanto se asume que la carpeta dist que contiene los archivos estaticos siempre se la traera del frontend. Por lo tanto la configuracion que no me permitió desplegar en railway es la siguiente.
 ```
  app.Static("/", "../client/dist") 
 ```
 - Ya que el proyecto esta desplegado en railway podemos acceder a el a travez de la siguiente url
 ```
 https://goreactmongo-production.up.railway.app/
 ``` 
 
 ## MOGODB

 - Instalamos paquete para leer una cadena de conexión de MongoDB desde una variable de entorno
 ```
 go get github.com/joho/godotenv
 ```
 
 
 
 
 
 
 
 
 
 
