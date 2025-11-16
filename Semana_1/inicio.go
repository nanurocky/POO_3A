/*
@titulo: programa principal del proyecto
@autor: Juan Carlos Pérez
@fecha: 27/10/2025
@Este programa presenta los aspectos iniciales de go
*/

// Llamar al paquete principal
package main

// Llamar a las diferentes librerías
import "fmt"

// Librería permite imprimir en consola

// función principal
func main() {
	// coloca el código inicial
	fmt.Println("¡Bienvenidos a la asignatura de POO en GO!")

	// variable:
	// tipos de datos: boleano, flotante, texto, entero
	// cuando vamos a declarar una varible debemos utilizar
	// palabra reservada var, := operador de asignación

	var edad int = 40
	var nombre string = "Juan Carlos"
	fmt.Println("Mi nombre es:", nombre, "y mi edad es", edad)

	// declaramos la variable con :=
	apellidos := "Pérez"
	numeroProductos := 100
	fmt.Println("Mis apellidos son:", apellidos, "y el número de producto es", numeroProductos)

	// variables constantes const
	const PI = 3.1416
	fmt.Println("El valor de la constante PI es:", PI)
	// tipos de datos
	// bool, int, float32, float64, string

	fmt.Println("Bienvenido a nuestro primer programa:")
	fmt.Print("Ingresa tu nombre:")
	fmt.Scanln(&nombre)
	fmt.Println("Ingresa tu edad:")
	fmt.Scanln(&edad)
	fmt.Println("Hola ", nombre, "tu edad es: ", edad)

	// operadores aritméticos + - * / %
	// operadores de comparación == != > < >= <=
	// operadores lógicos && || !
	// operadores de asignación += -= *= /= %=
	// operadores de incremento y decremento ++ --

	suma := 10 + 5
	resta := 10 - 5
	multiplicacion := 10 * 5
	division := 10 / 5
	modulo := 10 % 5
	fmt.Println("La suma es:", suma)
	fmt.Println("La resta es:", resta)
	fmt.Println("La multiplicacion es:", multiplicacion)
	fmt.Println("La division es:", division)
	fmt.Println("El modulo es:", modulo)

}
