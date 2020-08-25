package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	dupla := [2]string{"laaaargo", "cortito"}
	fmt.Println("La dupla es:", dupla[:])
	fmt.Println("El primero es mas largo?", esMasLargo(dupla[:]))
	fmt.Println("Unida con Y:", unirConY(dupla[:]))
	fmt.Println("Es palíndromo?", esPalindromo("palindromo"))
	fmt.Println("La letra a aparece", contarLetra("a", "palabra"), "veces en palabra")
	fmt.Println("La palabra a aparece", contarPalabra("casa", "la casa se sintio como una casa nuevamente"), "veces en la frase")
	fmt.Println("Sr Simpson desde ahora", reemplazarPalabra("Usted es Homero Simpson", "Simpson", "Thompson"))
	palabras := [4]string{"azul", "amarillo", "rojo", "negro"}
	fmt.Println("Las palabras son:", palabras[:])
	fmt.Println(enumerar(palabras[:]))
	numeros := [10]int{3, 5, 7, 8, 9, 0, 1, 2, 4, 6}
	fmt.Println("Los numeros son:", numeros[:])
	fmt.Println("El menor de la lista es", minimo(numeros[:]))
	fmt.Println("Lista ordenada", orden(numeros[:]))
	fmt.Println("Lista sin 0", borrador(0, numeros[:]))
	fmt.Println("Lista ordenada borrando", ordenBorrando(numeros[:]))
	numeros2 := [10]int{17, 13, -1, 8 - 20, 19, 10, 11, 12, 14, 16}
	fmt.Println("Los numeros para combinar son:", numeros2[:])
	fmt.Println("La lista resultante es", combinar(numeros[:], numeros2[:]))
	fmt.Println("La cant de parentecis es equilibrda en:()(())()es:", tieneParentesisBalanceados("()(())()"))
	fmt.Println("La cant de parentecis es equilibrda en:))((es:", tieneParentesisBalanceados("))(("))
	fmt.Println("La cant de parentecis es equilibrda en:(()es:", tieneParentesisBalanceados("(()"))
	fmt.Println("Distancia entre", distanciaDeLevenshtein("casa", "calle"))
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

func ordenBorrando(numeros []int) []int {
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

func orden(numeros []int) []int {
	valorTemporal := 0
	for i := 0; i < len(numeros); i++ {
		for x := 0; x < len(numeros); x++ {
			if numeros[x] > numeros[i] {
				valorTemporal = numeros[x]
				numeros[x] = numeros[i]
				numeros[i] = valorTemporal
			}
		}
	}
	return numeros
}

func borrador(numero int, listaNumeros []int) []int {
	largoLista := len(listaNumeros)
	nuevaLista := make([]int, (len(listaNumeros) - 1))
	aunNoBorro := true
	pibote := 0

	for i := 0; i < largoLista; i++ {
		if listaNumeros[i] == numero && aunNoBorro {
			aunNoBorro = false
		} else {
			nuevaLista[pibote] = listaNumeros[i]
			pibote = pibote + 1
		}
	}
	return nuevaLista
}

func combinar(lista1 []int, lista2 []int) []int {

	longitudCombinada := len(lista1) + len(lista2)
	listaResultado := make([]int, longitudCombinada)

	for i := 0; i < len(lista1); i++ {
		listaResultado[i] = lista1[i]
	}
	x := 0
	for i := len(lista1); i < longitudCombinada; i++ {
		listaResultado[i] = lista2[x]
		x = x + 1
	}

	listaResultado = orden(listaResultado)
	return listaResultado
}

func tieneParentesisNeutrlizados(frase string) bool {
	contadorAbiertos := contarLetra("(", frase)
	contadorCerrados := contarLetra(")", frase)

	return contadorAbiertos == contadorCerrados
}

func tieneParentesisBalanceados(frase string) bool {
	if !tieneParentesisNeutrlizados(frase) {
		return false
	}

	abierto := "("
	cerrado := ")"
	flagBorrro := true
	largoFrace := len(frase)
	resultado := true

	for i := 0; i < largoFrace; i++ {
		if frase[i] == abierto[0] {
			resultado = false
			for x := i + 1; x < largoFrace && flagBorrro; x++ {
				if frase[x] == cerrado[0] {
					frase = strings.Replace(frase, "(", "8", 1)
					frase = strings.Replace(frase, ")", "8", 1)
					flagBorrro = false
					resultado = true
				}
			}
			flagBorrro = true
		}
	}
	return resultado
}

func distanciaDeLevenshtein(s string, t string) []float64 {

	resultado := make([]float64, 2)

	// Costo penaliza si hay un cambio en el caracter.
	var costo float64 = 0

	// Obtenemos la longitud de las cadenas para crear los slices.
	m, n := len(s), len(t)
	d := make([][]float64, m+1, m*n)
	// Generamos el slice interno.
	for idx := range d {
		d[idx] = make([]float64, n+1)
	}

	// Recorre la matriz  de acuerdo a los cambios que encuentre en la cadena.
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				costo = 0
			} else {
				costo = 1
			}
			oper := math.Min(math.Min(d[i-1][j]+1, // Eliminacion
				d[i][j-1]+1), // Inserccion
				d[i-1][j-1]+costo) // Sustitucion
			d[i][j] = oper
		}
	}

	// Calcula el ratio de cambios en la palabra.
	var porcentaje float64 = 0
	if len(s) > len(t) {
		porcentaje = d[m][n] / float64(len(s))
	} else {
		porcentaje = d[m][n] / float64(len(t))
	}

	// Regresa distancia y ratio.
	resultado[0] = d[m][n]
	resultado[1] = porcentaje

	return resultado
}
