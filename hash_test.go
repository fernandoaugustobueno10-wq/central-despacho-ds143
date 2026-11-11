package main

import (
	"testing"
)

func TestTabelaHashInserirEBuscar(t *testing.T) {
	h := novaTabelaHash()

	h.Inserir(100, 1)
	h.Inserir(200, 2)

	indice, encontrado := h.Buscar(100)
	if !encontrado || indice != 1 {
		t.Errorf("esperado índice 1 para o ID 100, obtido %d (encontrado: %v)", indice, encontrado)
	}

	_, encontrado = h.Buscar(999)
	if encontrado {
		t.Error("ID 999 não deveria existir na tabela")
	}
}

func TestTabelaHashAtualizar(t *testing.T) {
	h := novaTabelaHash()

	h.Inserir(50, 5)

	h.Inserir(50, 10)

	indice, encontrado := h.Buscar(50)
	if !encontrado || indice != 10 {
		t.Errorf("esperado índice 10 para o ID 50 após atualização, obtido %d (encontrado: %v)", indice, encontrado)
	}
}

func TestTabelaHashRemover(t *testing.T) {
	h := novaTabelaHash()

	h.Inserir(100, 1)
	h.Inserir(200, 2)

	removido := h.Remover(200)

	indice, encontrado := h.Buscar(200)
	if !removido || encontrado {
		t.Errorf("esperado remoção do indíce 200, obtido na busca %d (encontrado: %v)", indice, encontrado)
	}

	removido = h.Remover(300)
	if removido {
		t.Errorf("esperado falha na remoção do indíce inexistente 300, porém retornou true (removido: %v)", removido)
	}
}

func TestTabelaHashRedimensionamento(t *testing.T) {
	h := novaTabelaHash()

	for i := 0; i < 2000; i++ {
		h.Inserir(i, i)
	}

	for i := 0; i < 2000; i++ {
		indice, encontrado := h.Buscar(i)
		if !encontrado {
			t.Errorf("um dos elementos não foi encontrado após 2000 inserções, verificar redimensionamento. obtido na busca %d (encontrado: %v)", indice, encontrado)
		}
	}
}
