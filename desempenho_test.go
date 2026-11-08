package main

import "testing"

func TestNaoBuscaChamadoInviavelNoFimDaPartida(t *testing.T) {
	prepararTeste("")

	totalTurnos = 1000
	defer func() { totalTurnos = 0 }()

	hospitais = []Hospital{{X: 0, Y: 0}}
	gridEspacial.PrepararHospitais(hospitais)

	adicionarChamado(Chamado{
		ID: 1, X: 175, Y: 175, Valor: 100, Prazo: 1200,
	})

	alvo := selecionarChamado(
		Ambulancia{ID: 0, X: 0, Y: 0, Livre: true},
		999, nil, -1,
	)

	if alvo != -1 {
		t.Fatalf("chamado sem tempo de entrega selecionado: %d", alvo)
	}
}

func TestBuscaAindaSelecionaChamadoViavel(t *testing.T) {
	prepararTeste("")

	totalTurnos = 1000
	defer func() { totalTurnos = 0 }()

	hospitais = []Hospital{{X: 0, Y: 0}}
	gridEspacial.PrepararHospitais(hospitais)

	adicionarChamado(Chamado{
		ID: 1, X: 2, Y: 0, Valor: 30, Prazo: 1000,
	})

	alvo := selecionarChamado(
		Ambulancia{ID: 0, X: 0, Y: 0, Livre: true},
		995, nil, -1,
	)

	if alvo < 0 || chamados[alvo].ID != 1 {
		t.Fatalf("chamado viavel nao selecionado: %d", alvo)
	}
}
