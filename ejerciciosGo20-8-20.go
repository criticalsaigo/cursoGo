package main

import "fmt"

func main() {
	dupla := [2]string{"laaaargo", "cortito"}
	fmt.Println("La dupla es:", dupla[:])
	fmt.Println("El primero es mas largo?", esMasLargo(dupla[:]))
	fmt.Println("Unida con Y:", unirConY(dupla[:]))
	fmt.Println("Es palíndromo?", esPalindromo("palindromo"))
	fmt.Println("La letra a aparece", contarLetra("a", "palabra"), "veces en palabra")
	fmt.Println("La palabra a aparece", contarPalabra("casa", "la casa se sintio como una casa nuevamente"), "veces en la frase")
	//	fmt.Println("La palabra a aparece", contarPalabra2("casa", "la casa se sintio como una casa nuevamente"), "veces en la frase")
	fmt.Println("Sr Simpson desde ahora", reemplazarPalabra("Usted es Homero Simpson", "Simpson", "Thompson"))
	palabras := [4]string{"azul", "amarillo", "rojo", "negro"}
	fmt.Println("Las palabras son:", palabras[:])
	fmt.Println(enumerar(palabras[:]))
	numeros := [10]int{3, 5, 7, 8, 9, 0, 1, 2, 4, 6}
	fmt.Println("Los numeros son:", numeros[:])
	fmt.Println("El menor de la lista es", minimo(numeros[:]))
	fmt.Println("El menor de la lista es", borrador(0, numeros[:]))

}

func esMasLargo(dupla []string) bool {
	return len(dupla[0]) > len(dupla[1])
}

func unirConY(dupla []string) string {
	return dupla[0] + " y " + dupla[1]
}

func esPalindromo(palabra string) bool {
	return palabra == "palíndromo"
}

func contarLetra(letra string, palabra string) int {
	contador := 0

	for i := 0; i < len(palabra); i++ {
		if palabra[i] == letra[0] {
			contador = contador + 1
		}
	}
	return contador
}

func contarPalabra(palabra string, frase string) int {
	contador := 0

	if len(frase) < len(palabra) {
		return contador
	}

	for i := 0; i < len(frase); i++ {
		if frase[i] == palabra[0] {
			vaBien := true
			for y := 1; y < len(palabra) && vaBien; y++ {

				if frase[i+y] != frase[i+y] {
					vaBien = false
				}
			}
			if vaBien {
				contador = contador + 1
			}
		}

	}
	contador = contador - 1
	return contador
}

/*Aca busco crear un string desde el pibote del for inicial para comprarlo de una con palabra

func contarPalabra2(palabra string, frase string) int {
	contador := 0
	var minifrase string = "a"
	if len(frase) < len(palabra) {
		return contador
	}

	for i := 0; i < len(frase); i++ {
		miniFrase := frase[i]

		for y := 0; y < len(palabra); y++ {
			pibote := i + y
			miniFrase = miniFrase + frase[pibote]
		}
		if miniFrase == palabra {
			contador = contador + 1
		}
	}
	return contador
}
*/

func reemplazarPalabra(frase string, palabra string, reemplazo string) string {
	for i := 0; i < len(frase); i++ {
		if frase[i] == palabra[0] {
			vaBien := true
			for y := 1; y < len(palabra) && vaBien; y++ {

				if frase[i+y] != frase[i+y] {
					vaBien = false
				}
			}
			if vaBien {
				salto := i + len(palabra)
				frase = frase[0:i] + reemplazo + frase[salto:]
			}
		}

	}
	return frase
}

func enumerar(palabras []string) string {
	resultado := palabras[0]

	if len(palabras) == 1 {
		return resultado
	}

	for i := 1; i < len(palabras); i++ {

		y := i + 1
		if y == len(palabras) {
			resultado = resultado + " y " + palabras[i]
		} else {
			resultado = resultado + ", " + palabras[i]
		}

	}

	return resultado
}

func minimo(numeros []int) int {
	minimo := numeros[0]

	for i := 1; i < len(numeros); i++ {
		if minimo > numeros[i] {
			minimo = numeros[i]
		}

	}
	return minimo
}

func orden(numeros []int) []int {
	largoNumeros := len(numeros)

	if largoNumeros == 1 {
		return numeros
	}

	ordenados := make([]int, largoNumeros)

	for i := 0; i < largoNumeros; i++ {
		minimo := minimo(numeros)
		ordenados[i] = minimo
		numeros = borrador(minimo, numeros)
	}
	return ordenados
}

func borrador(numero int, listaNumeros []int) []int {
	nuevaLista := make([]int, (len(listaNumeros) - 1))
	aunNoBorro := true

	for i := 0; i < len(listaNumeros); i++ {
		if listaNumeros[i] == numero && aunNoBorro {
			aunNoBorro = false
		} else {
			nuevaLista[i] = listaNumeros[i]
		}
	}
	return nuevaLista
}

func combinar(lista1 []int, lista2 []int) []int {
	listaResultado := make([]int, (len(lista1) + len(lista2)))

	for i := 0; i < (len(lista1) + len(lista2)); i++ {
		if i < len(lista1) {
			listaResultado[i] = lista1[i]
		} else {
			listaResultado[i] = lista2[i]
		}
	}
	listaResultado = orden(listaResultado)
	return listaResultado
}
