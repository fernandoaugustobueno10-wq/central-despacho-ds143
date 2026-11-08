package main

import "testing"

func TestSelecaoComCemAlvosReservados(t *testing.T) {
	prepararTeste("")
	alvosPorAmbulancia = nil
	totalTurnos = 0
	defer func() { totalTurnos = 0 }()

	reservados := make([]int, 0, 100)
	for id := 1; id <= 101; id++ {
		adicionarChamado(Chamado{ID: id, X: 2, Y: 2, Valor: 20, Prazo: 1000})
		if id <= 100 {
			reservados = append(reservados, id)
		}
	}

	idx := selecionarChamado(Ambulancia{ID: 0, X: 0, Y: 0, Livre: true}, 1, reservados, -1)
	if idx < 0 || chamados[idx].ID != 101 {
		t.Fatalf("alvo nao reservado esperado: 101, indice encontrado: %d", idx)
	}
}
