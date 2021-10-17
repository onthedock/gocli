# README

Para alojar el proyecto, he creado una carpeta llamada `gocli`.

En la subcarpeta `code/` es donde residirá el código en Go de la aplicación, mientras que la documentación (este README.md) se encuentra en la raíz de `gocli/`.

Desde la línea de comando, inicializo el módulo mediante

```bash
go mod init github.com/onthedock/gocli
```

Dentro de la carpeta `code/` genero el fichero `main.go`.

```go
package main

import "fmt"

func main() {
	fmt.Println("gocli project will be here")
}
```

Para verificar que todo funciona por ahora, desde `gocli/` ejecuto:

```bash
$ go run github.com/onthedock/gocli
gocli project will be here
```
