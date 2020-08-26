package main

import "fmt"

type Golondrina struct {
	nombre  string
	energia float64
	peso    float64
	vida    float64
	ataque  float64
}

type Paloma struct {
	nombre  string
	energia float64
	peso    float64
	vida    float64
	ataque  float64
}

type Volador interface {
	volar(km float64)
}

func main() {
	pepita := naceGolondrina("Pepita")

	fmt.Println("Nace una golondrina:")
	pepita.imprimirEstado()

	fmt.Println("Vive, come y vuela")

	comer(&pepita, 5)
	pepita.volar(2)

	pepita.imprimirEstado()

	palomon := nacePaloma("Palomon")
	fmt.Println("Nace una paloma:")
	palomon.imprimirEstado()

	robar(&palomon, 100)
	palomon.volar(2)

	fmt.Println("Vive, roba y vuela")
	palomon.imprimirEstado()

	//	volarMucho(&palomon, &pepita)
	fmt.Println("Ahora combaten a muerte")
	fmt.Println(combateGolondrinaVsPaloma(&pepita, &palomon))

	pepita.imprimirEstado()
	palomon.imprimirEstado()

}

func naceGolondrina(nombre string) Golondrina {
	var nuevaAve Golondrina
	nuevaAve.nombre = nombre
	nuevaAve.energia = 10
	nuevaAve.peso = 1
	nuevaAve.vida = 100
	nuevaAve.ataque = 25
	return nuevaAve
}

func (golondrina Golondrina) imprimirEstado() {
	fmt.Println("Nombre", golondrina.nombre)
	fmt.Println("Energia", golondrina.energia)
	fmt.Println("Peso", golondrina.peso)
	fmt.Println("Vida", golondrina.vida)
	fmt.Println("Ataque", golondrina.ataque)
}
func (paloma Paloma) imprimirEstado() {
	fmt.Println("Nombre", paloma.nombre)
	fmt.Println("Energia", paloma.energia)
	fmt.Println("Peso", paloma.peso)
	fmt.Println("Vida", paloma.vida)
	fmt.Println("Ataque", paloma.ataque)
}

func nacePaloma(nombre string) Paloma {
	var nuevaAve Paloma
	nuevaAve.nombre = nombre
	nuevaAve.energia = 8
	nuevaAve.peso = 2
	nuevaAve.vida = 120
	nuevaAve.ataque = 15
	return nuevaAve
}

func (golondrina *Golondrina) volar(km float64) {
	golondrina.energia -= km * 1.3
}

func comer(golondrina *Golondrina, kgComida float64) {
	golondrina.energia += kgComida * 0.8
}

func (paloma *Paloma) volar(km float64) {
	paloma.energia -= km * 1.3 * (paloma.peso * 0.1)
}

func robar(paloma *Paloma, kgComida float64) {
	paloma.energia += kgComida * 0.1
}

func volarMucho(voladores ...Volador) {
	for i := 0; i < len(voladores); i++ {
		voladores[i].volar(100)
	}
}

func combateGolondrinaVsPaloma(golondrina *Golondrina, paloma *Paloma) string {
	for golondrina.vida > 0 && paloma.vida > 0 {
		if golondrina.energia > paloma.energia {
			paloma.vida -= golondrina.ataque
			if paloma.vida > 0 {
				golondrina.vida -= paloma.ataque
			}
		} else {
			golondrina.vida -= paloma.ataque
			if golondrina.vida > 0 {
				paloma.vida -= golondrina.ataque
			}
		}
	}
	if golondrina.vida > 0 {
		return "palomon murio"
	} else {
		return "golondrina murio"
	}
}
