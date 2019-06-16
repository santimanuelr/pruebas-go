package waldo

import (
	"fmt"
	"sync"
)

/*

Manejo de errores
Concurrencia
Como construir servicios rest

err := doSth()
if err != nil {

}

Siempre manejar los errores. De esto esta lleno los archivos GO


Panic es parecido al runtimeException

defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()

eso es como el finally en java. Es importante para cerrar bien todos los flujos


Ejercicio
Armar un programa que permita simular un punto de un partido de Tenis
entre Del Potro y Djokovic, usando goroutines y unbuffered channels.
Extra: Tratemos de que gane Del Potro.

*/


func Partido() {

	// Create a wait group.
	var wg sync.WaitGroup
	var sChannelDjokovic chan string = make(chan string, 2)
	var sChannelDelPotro chan string = make(chan string, 2)

	wg.Add(2)
	go func() {
		fmt.Println("Hello del potro!")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello Djokovic!")
		wg.Done()
	}()




	// Wait until everyone finishes.
	wg.Wait()
}