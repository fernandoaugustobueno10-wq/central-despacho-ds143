package main

import (
	"testing"
)

// Teste para validar inserção e busca básica
func TestTabelaHashInserirEBuscar(t *testing.T) {
	h := novaTabelaHash()

	// TODO: Insere alguns elementos (ex: ID 100 com índice 1, ID 200 com índice 2)
	h.Inserir(100, 1)
	h.Inserir(200, 2)

	// TODO: Testa se a busca por um ID existente retorna o índice correto e true
	indice, encontrado := h.Buscar(100)
	if !encontrado || indice != 1 {
		t.Errorf("esperado índice 1 para o ID 100, obtido %d (encontrado: %v)", indice, encontrado)
	}

	// TODO: Testa se a busca por um ID inexistente retorna false
	_, encontrado = h.Buscar(999)
	if encontrado {
		t.Error("ID 999 não deveria existir na tabela")
	}
}

// Teste para validar a atualização de valores para um mesmo ID
func TestTabelaHashAtualizar(t *testing.T) {
	h := novaTabelaHash()

	// TODO: Insere um ID com um índice inicial
	h.Inserir(50, 5)

	// TODO: Insere o mesmo ID com um novo índice
	h.Inserir(50, 10)

	// TODO: Verifica se o valor foi atualizado corretamente na busca
	indice, encontrado := h.Buscar(50)
	if !encontrado || indice != 10 {
		t.Errorf("esperado índice 10 para o ID 50 após atualização, obtido %d (encontrado: %v)", indice, encontrado)
	}
}

// Teste para validar a remoção de elementos
func TestTabelaHashRemover(t *testing.T) {
	h := novaTabelaHash()

	// TODO: Insere elementos, remove um deles usando h.Remover(id)
	h.Inserir(100, 1)
	h.Inserir(200, 2)

	removido := h.Remover(200)
	// e valida se a remoção retorna true e se o elemento deixa de ser encontrado.
	indice, encontrado := h.Buscar(200)
	if !removido || encontrado {
		t.Errorf("esperado remoção do indíce 200, obtido na busca %d (encontrado: %v)", indice, encontrado)
	}
	// Valida também remover um elemento que não existe (deve retornar false).
	removido = h.Remover(300)
	if removido {
		t.Errorf("esperado falha na remoção do indíce inexistente 300, porém retornou true (removido: %v)", removido)
	}
}

// Teste para forçar e validar o redimensionamento (rehash)
func TestTabelaHashRedimensionamento(t *testing.T) {
	h := novaTabelaHash()

	// TODO: Insira uma quantidade grande de elementos (ex: 2000 itens)
	// para ultrapassar o limite inicial e acionar o redimensionamento automático.
	for i := 0; i < 2000; i++ {
		h.Inserir(i, i)
	}
	// No final, verifique se todos os itens inseridos continuam acessíveis via Buscar.
	for i := 0; i < 2000; i++ {
		indice, encontrado := h.Buscar(i)
		if !encontrado {
			t.Errorf("um dos elementos não foi encontrado após 2000 inserções, verificar redimensionamento. obtido na busca %d (encontrado: %v)", indice, encontrado)
		}
	}
}
