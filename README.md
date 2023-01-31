# GOReactMongo

# SERVER

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
 ### Uso del mÓdulo air
 - En la carpeta raiz del server ejecutamos en la terminal el comando
 ```
 air .
 ```
## Ejecución
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
 ## Preview Deploy in Railway
 - Ya que el proyecto esta desplegado en railway podemos acceder a el a travez de la siguiente url
 ```
 https://goreactmongo-production.up.railway.app/
 ``` 
 
 ## MOGODB

 - Instalamos paquete para leer una cadena de conexión de MongoDB desde una variable de entorno
 ```
 go get github.com/joho/godotenv
 ```
 
 # CLIENT
 - Usamos vite y typeScript para crear el proyecto frontend 
 - Por alguna razon la configuracion de proxy en el vite.config para establecer la url del server, dio error en el server al momento de comprobar las peticiones http, por lo que se opto por usar una variable de entornor en vite en un archivo .env  para la url, sin embargo hay que tomar en cuenta que en vite se maneja distinto, incluso desde la declaracion que debe empezar por VITE como se muestra en la primera linea , y se lo llama como se muestra en la segunda linea
 ```
 VITE_URLSERVER=http://localhost:3000/api/v1
 ```
 ```
  await axios.post(import.meta.env.VITE_URLSERVER+"/users"); 
 ```
 
 
 
 
 
 
 
 
