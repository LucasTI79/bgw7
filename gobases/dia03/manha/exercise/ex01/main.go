// Crie um programa que atenda aos seguintes pontos:
// Ter uma estrutura chamada Product com os campos ID, Name, Price, Description e Category.
// Ter uma fatia global de Produto chamada Produtos instanciada com valores. 2 métodos
// associados à estrutura Produto: Save(), GetAll(). O método Save() deve pegar a fatia de
// Products e adicionar o produto a partir do qual o método é chamado. O método GetAll()
// imprime todos os produtos salvos na fatia Products.
// Uma função getById() para a qual um INT deve ser passado como parâmetro e retorna o
// produto correspondente ao parâmetro passado.
// Execute pelo menos uma vez cada método e função definidos em main().

package main

type Marketplace struct {
	Products []Product
}

func (m *Marketplace) Save()  {}
func (m Marketplace) GetAll() {}

type Product struct{}
