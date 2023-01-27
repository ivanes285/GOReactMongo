# GOReactMongo

## Iniciamos el proyecto go con el comando

- go mod init ES EQUIVALENTE en node a npm int  sin embargo nos pide ingresar la direccion de un repositorio donde se va alojar
```
go mod init github.com/ivanes285/GOReactMongo  
```

### Instalacion de modulos

- Instalamos un framework llamado fiber equivalente(express en js) para trabajar con go
```
go get github.com/gofiber/fiber/v2
```
- Importamos Cors para poder habilitarlos para el frontend
```
 "github.com/gofiber/fiber/v2/middleware/cors"
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

