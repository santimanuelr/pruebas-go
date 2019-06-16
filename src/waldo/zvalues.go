package waldo

import "fmt"

var i int
var f float64
var b bool
var s string

func init() { fmt.Println("Se ejecuta init") }

func Print0Val() { fmt.Printf("%v %v %v %q\n", i, f, b, s) }

/*Si le pongo mayus a los metodos señala que son públicos. Minuscula -> privados*/
/*No hay exceptiones en go*/
