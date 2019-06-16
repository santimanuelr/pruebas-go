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



/*Hacer un programa que reciba un parámetro p por consola y que
calcule la suma de los valores v tales que 0 <= v <= p,
siendo v un número divisible por 3 o por 5*/

func SumaDivisibles (entero int) int {
	if entero <= 0 {
		return 0
	} else {
		if entero % 3 == 0 || entero % 5 == 0 {
			return entero + SumaDivisibles(entero - 1)
		} else {
			return SumaDivisibles(entero - 1)
		}
	}
}




/*

Segunda clase de golang
Punteros.
en los metodos se pasa por parametro una copia del puntero.

type Vertex struct
No tenemos clases pero tenemos struct!


Slices
Arreglos dinamicos.
Se le puede agregar elementos. Ojo! con la capacidad. Sobreescribe el arreglo original o agrega una posición.


package main

import "fmt"

func main() {
    months := [...]string{1: "January", 2:"February",3:"March", 4: "April" ,5:"May", 6:"June",
        7:"July",8:"August",9:"September",10:"October",11:"November", 12: "December"}

    summer := months[6:9]

    Q3yQ4 := months[7:]
    Q3yQ4 = append(Q3yQ4, "New_month")
    Q3yQ4[1] = "new_August"
    fmt.Println(Q3yQ4)
    fmt.Println(summer)
}




*/