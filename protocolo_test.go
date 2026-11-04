package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func prepararTeste(entrada string) *bytes.Buffer {
	leitor = bufio.NewScanner(strings.NewReader(entrada))
	leitor.Buffer(make([]byte, 1<<16), 1<<20)

	var saida bytes.Buffer
	escritor = bufio.NewWriter(&saida)

	hospitais = nil
	chamados = nil

	return &saida
}

func TestLerInicio(t *testing.T) {
	entrada := `INICIO 200 200 500 100 2000 0 2
HOSPITAIS 2
20 30
150 170
AMBULANCIAS 2
0 20 30
1 150 170
BLOQUEIOS 0
FIM
`

	prepararTeste(entrada)
	lerInicio()

	if len(hospitais) != 2 {
		t.Fatalf("esperado 2 hospitais, obtido %d", len(hospitais))
	}

	if hospitais[0].X != 20 || hospitais[0].Y != 30 {
		t.Errorf("coordenadas incorretas do primeiro hospital")
	}

	if hospitais[1].X != 150 || hospitais[1].Y != 170 {
		t.Errorf("coordenadas incorretas do segundo hospital")
	}
}

func TestLerTurno(t *testing.T) {
	entrada := `TURNO 1
CHAMADOS 2
1683149752 25 30 15 120
2096790314 140 160 75 60
RECOLHIDOS 0
AMBULANCIAS 2
0 20 30 LIVRE
1 150 170 OCUPADA
PLACAR 2
0 0
1 0
FIM
`

	prepararTeste(entrada)

	turno, ambulancias, ok := lerTurno()

	if !ok {
		t.Fatal("falha ao ler turno")
	}

	if turno != 1 {
		t.Errorf("esperado turno 1, obtido %d", turno)
	}

	if len(chamados) != 2 {
		t.Errorf("esperados 2 chamados, obtidos %d", len(chamados))
	}

	if len(ambulancias) != 2 {
		t.Fatalf("esperadas 2 ambulancias, obtidas %d", len(ambulancias))
	}

	if !ambulancias[0].Livre {
		t.Error("ambulancia 0 deveria estar livre")
	}

	if ambulancias[1].Livre {
		t.Error("ambulancia 1 deveria estar ocupada")
	}
}

func TestRemoverChamadoRecolhido(t *testing.T) {
	entrada := `TURNO 2
CHAMADOS 0
RECOLHIDOS 1
100
AMBULANCIAS 0
PLACAR 2
0 0
1 0
FIM
`

	prepararTeste(entrada)

	chamados = []Chamado{
		{ID: 100, X: 10, Y: 10, Valor: 15, Prazo: 50},
		{ID: 200, X: 20, Y: 20, Valor: 30, Prazo: 60},
	}

	_, _, ok := lerTurno()

	if !ok {
		t.Fatal("falha ao ler turno")
	}

	if len(chamados) != 1 {
		t.Fatalf("esperado 1 chamado, obtidos %d", len(chamados))
	}

	if chamados[0].ID != 200 {
		t.Errorf("chamado incorreto permaneceu na lista")
	}
}

func TestRemoverChamadoExpirado(t *testing.T) {
	entrada := `TURNO 10
CHAMADOS 0
RECOLHIDOS 0
AMBULANCIAS 0
PLACAR 2
0 0
1 0
FIM
`

	prepararTeste(entrada)

	chamados = []Chamado{
		{ID: 100, X: 10, Y: 10, Valor: 15, Prazo: 9},
		{ID: 200, X: 20, Y: 20, Valor: 30, Prazo: 10},
	}

	_, _, ok := lerTurno()

	if !ok {
		t.Fatal("falha ao ler turno")
	}

	if len(chamados) != 1 {
		t.Fatalf("esperado 1 chamado, obtidos %d", len(chamados))
	}

	if chamados[0].ID != 200 {
		t.Error("chamado com prazo valido foi removido")
	}
}

func TestEnviarPronto(t *testing.T) {
	saida := prepararTeste("")

	enviarPronto()

	if saida.String() != "PRONTO\n" {
		t.Errorf("resposta incorreta: %q", saida.String())
	}
}

func TestRespostaTurno(t *testing.T) {
	saida := prepararTeste("")

	hospitais = []Hospital{
		{X: 0, Y: 0},
	}

	chamados = []Chamado{
		{ID: 100, X: 5, Y: 0, Valor: 15, Prazo: 50},
	}

	ambulancias := []Ambulancia{
		{ID: 0, X: 0, Y: 0, Livre: true},
		{ID: 1, X: 0, Y: 2, Livre: false},
	}

	jogarTurno(1, ambulancias)

	esperado := "0 L\n1 N\nFIM\n"

	if saida.String() != esperado {
		t.Errorf(
			"resposta incorreta\nEsperado:\n%sObtido:\n%s",
			esperado,
			saida.String(),
		)
	}
}

func TestFimDaEntrada(t *testing.T) {
	prepararTeste("")

	_, _, ok := lerTurno()

	if ok {
		t.Error("esperado encerramento ao fim da entrada")
	}
}
