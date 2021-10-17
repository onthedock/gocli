package main

import (
	"flag"
	"fmt"
	"strings"
)

// Define a new type of variable
type Color string

// Define colors as Constants of type Color
const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed    Color = "\u001b[31m"
	ColorGreen  Color = "\u001b[32m"
	ColorOrange Color = "\u001b[33m"
	ColorBlue   Color = "\u001b[34m"
	ColorReset  Color = "\u001b[0m"
)

// Show output in the passed color
func colorize(color Color, message string) {
	// color and ColorReset son de tipo "Color", por lo que hay que convertirlas en string
	// antes de imprimirlas usando Println (que requiere un cadena)
	fmt.Println(string(color), message, string(ColorReset))
}

func main() {
	msg := "Texto por defecto a mostrar"
	c := ColorGreen
	useColor := flag.String("color", "ColorGreen", "Display colorized output")
	flag.Parse()

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
	colorize(c, msg)
}
