package main

import "testing"

func TestNaoSelecionaPacienteSemTempoDeEntrega(t *testing.T) {
	prepararTeste("")
	totalTurnos = 10
	defer func() { totalTurnos = 0 }()
	hospitais = []Hospital{{X: 0, Y: 0}}
	adicionarChamado(Chamado{ID: 10, X: 7, Y: 0, Valor: 100, Prazo: 100})
	adicionarChamado(Chamado{ID: 20, X: 2, Y: 0, Valor: 10, Prazo: 100})
	a := Ambulancia{ID: 0, X: 0, Y: 0, Livre: true}
	idx := selecionarChamado(a, 5, nil, -1)
	if idx == -1 || chamados[idx].ID != 20 {
		t.Fatalf("chamado esperado 20, indice obtido: %d", idx)
	}
}

func TestConsideraDistanciaDoChamadoAoHospital(t *testing.T) {
	prepararTeste("")
	totalTurnos = 0
	hospitais = []Hospital{{X: 0, Y: 0}}
	a := Ambulancia{ID: 0, X: 0, Y: 0, Livre: true}
	proximoDoHospital := Chamado{ID: 10, X: 2, Y: 0, Valor: 20, Prazo: 100}
	longeDoHospital := Chamado{ID: 20, X: 15, Y: 0, Valor: 20, Prazo: 100}
	if pontuacaoChamado(a, proximoDoHospital, 1, -1) <= pontuacaoChamado(a, longeDoHospital, 1, -1) {
		t.Fatal("a pontuacao deveria favorecer o ciclo de atendimento mais curto")
	}
}

func TestGuardaDuracaoDaPartida(t *testing.T) {
	entrada := `INICIO 200 200 500 100 2000 0 2
HOSPITAIS 1
0 0
AMBULANCIAS 1
0 0 0
BLOQUEIOS 0
FIM
`
	prepararTeste(entrada)
	lerInicio()
	if totalTurnos != 500 {
		t.Fatalf("esperado 500 turnos, obtido %d", totalTurnos)
	}
	totalTurnos = 0
}
