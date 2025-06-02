package main

import "fmt"

// tipo personalizado
type MyType interface {
	int | string
}

func main() {
	// criando um mapa
	var mapa map[string]int
	// mapa := map[string]int{}
	// mapa := make(map[string]int)

	// criando uma chave
	mapa["Benjamin"] = 20

	// deletando uma chave
	delete(mapa, "Benjamin")

	// percorrendo todo o mapa
	for key, value := range mapa {
		fmt.Printf("key => %s, value => %d", key, value)
	}

	// genericMap := make(map[string]interface{})
	genericMap := make(map[string]any)
	genericMap["batata"] = 20
	genericMap["cenoura"] = "ralada"
}

// func Sum(...any) {

// }
