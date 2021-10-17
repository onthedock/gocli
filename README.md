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

Go dispone del *package* [flag](https://pkg.go.dev/flag), que permite definir *flags* mediante `flag.String()`, `Bool()` o `Int()`.

Usamos el ejemplo del artículo [How To Use the Flag Package in Go
](https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go) de Digital Ocean.

En el artículo se crea una utilidad que muestra un mensaje por pantalla (no personalizable). Usando el paquete `flag`, se define un parámetro de tipo `bool`; si está presente, el mensaje se muestra en color; si no, se muestra en el color predeterminado.

La aplicación de ejemplo del artículo define un tipo de variable `Color` y un conjunto de constantes para los diferentes colores.

```go
// Define a new type of variable
type Color string

// Define colors as Constants of type Color
const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed    Color = "\u001b[31m"
	ColorGreen  Color = "\u001b[32m"
	ColorYellow Color = "\u001b[33m"
	ColorBlue   Color = "\u001b[34m"
	ColorReset  Color = "\u001b[0m"
)
```

Se define una función `colorize(color Color, message string)` que acepta un *color* (una variable de tipo `Color`) y el mensaje a mostrar.

```go
...
// Show output in the passed color
func colorize(color Color, message string) {
	// color and ColorReset son de tipo "Color", por lo que hay que convertirlas en string
	// antes de imprimirlas usando Println (que requiere un cadena)
	fmt.Println(string(color), message, string(ColorReset))
}
...
```

En la función principal `main` se define el *flag* `color`; mediante un `if` se valida si el *flag* está presente para llamar a la función y mostrar el texto en color.

```go
...
func main() {
	useColor := flag.Bool("color", false, "Display colorized output")
	flag.Parse()

	if *useColor {
		colorize(ColorGreen, "Hello world default message")
		return
	}
	fmt.Println("Hello with no color")
}
```

Vamos a modificar el ejemplo para que podamos especificar el color en el que se muestra el texto.

Empezamos modificando el tipo del *flag* para que sea de tipo *String*:

```go
	useColor := flag.String("color", "ColorGreen", "Display colorized output")
```

A continuación, usamos un `switch` para seleccionar el color adecuado en función del valor del *flag* `-color`. Para simplificar la evaluación de la condición en el comando `switch`, primero convertimos el valor del *flag* a minúsculas (usando `strings.ToLower(string)`):

```go
	switch strings.ToLower(string(*useColor)) {
	case "black":
		c = ColorBlack
	case "orange":
		c = ColorOrange
	case "green":
		c = ColorGreen
	case "blue":
		c = ColorBlue
	case "red":
		c = ColorRed
	default:
		c = ColorReset
	}
```

Mediante el caso por defecto, hacemos frente a otros colores que no tengamos definidos, usando en este caso el valor definido por defecto.

Finalmente, llamamos la función `colorize(c, msg)`.
