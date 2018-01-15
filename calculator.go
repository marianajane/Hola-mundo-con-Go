package main

import "fmt"

type Perrito struct {
	raza string
	color string
	precio float32
	genero bool
}

func main() {
	var perrito_negro = Perrito{"Husky", "negro", 25.50, true}

	fmt.Println(perrito_negro.raza)


	var numero1 float32 = 132
	var numero2 float32 = 23

	fmt.Println("Calculadora 1")
	calculadora(numero1, numero2)
	fmt.Println("-----------------------")

	holaMundo()
	u2()

	var numero3 float32 = 44
	var numero4 float32 = 7

	fmt.Println("Calculadora 2")
	calculadora(numero3, numero4)

}

func holaMundo(){
	fmt.Println("Hola mundo")
}

func u2(){
	fmt.Println("I still haven't foundd what im looking for")
}

func operacion(n1 float32, n2 float32, op string) float32{
	var resultado float32
	if(op == "+") {
		resultado = n1 + n2
	}
	if(op == "-") {
		resultado = n1 - n2
	}
	if(op == "*") {
		resultado = n1 * n2
	}
	if(op == "/") {
		resultado = n1 / n2
	}
	return resultado
}

func calculadora(numero1 float32, numero2 float32){
	fmt.Print("La suma es : ")
	fmt.Println(operacion (numero1, numero2, "+"))

	fmt.Print("La resta es : ")
	fmt.Println(operacion (numero1, numero2, "-"))

	fmt.Print("La multiplicación es : ")
	fmt.Println(operacion (numero1, numero2, "*"))

	fmt.Print("La división es : ")
	fmt.Println(operacion (numero1, numero2, "/"))
}
