package main

import (
	"strings"
	"testing"
)

func TestAlvosDistintos(t *testing.T) {
	saida := prepararTeste("")
	alvosPorAmbulancia = nil
	hospitais = []Hospital{{X: 0, Y: 0}}

	adicionarChamado(Chamado{ID: 101, X: 5, Y: 0, Valor: 10, Prazo: 40})
	adicionarChamado(Chamado{ID: 102, X: 0, Y: 5, Valor: 10, Prazo: 40})

	jogarTurno(1, []Ambulancia{
		{ID: 0, X: 0, Y: 0, Livre: true},
		{ID: 1, X: 0, Y: 0, Livre: true},
	})

	if alvosPorAmbulancia[0] == alvosPorAmbulancia[1] ||
		alvosPorAmbulancia[0] == 0 ||
		alvosPorAmbulancia[1] == 0 {
		t.Fatalf("alvos nao exclusivos: %v", alvosPorAmbulancia)
	}

	if !strings.Contains(saida.String(), "FIM\n") {
		t.Fatal("resposta nao terminou")
	}
}

func TestPrazoInviavel(t *testing.T) {
	prepararTeste("")
	alvosPorAmbulancia = nil

	adicionarChamado(Chamado{ID: 101, X: 25, Y: 0, Valor: 100, Prazo: 2})
	adicionarChamado(Chamado{ID: 102, X: 3, Y: 0, Valor: 5, Prazo: 15})

	a := Ambulancia{ID: 0, X: 0, Y: 0, Livre: true}
	idx := selecionarChamado(a, 1, nil, -1)

	if idx == -1 || chamados[idx].ID != 102 {
		t.Fatalf("selecionou chamado inviavel: %d", idx)
	}
}

func TestValorCompensaDistancia(t *testing.T) {
	prepararTeste("")
	alvosPorAmbulancia = nil

	adicionarChamado(Chamado{ID: 101, X: 3, Y: 0, Valor: 2, Prazo: 60})
	adicionarChamado(Chamado{ID: 102, X: 7, Y: 0, Valor: 90, Prazo: 60})

	idx := selecionarChamado(
		Ambulancia{ID: 0, X: 0, Y: 0, Livre: true},
		1, nil, -1,
	)

	if idx == -1 || chamados[idx].ID != 102 {
		t.Fatalf("nao priorizou valor: %d", idx)
	}
}

func TestReatribuiAlvoRemovido(t *testing.T) {
	prepararTeste("")
	alvosPorAmbulancia = nil

	adicionarChamado(Chamado{ID: 101, X: 5, Y: 0, Valor: 50, Prazo: 60})
	adicionarChamado(Chamado{ID: 102, X: 0, Y: 5, Valor: 15, Prazo: 60})

	a := []Ambulancia{{ID: 0, X: 0, Y: 0, Livre: true}}

	jogarTurno(1, a)

	if alvosPorAmbulancia[0] != 102 {
		t.Fatal("alvo inicial incorreto")
	}

	removerPorID(101)
	jogarTurno(2, a)

	if alvosPorAmbulancia[0] != 103 {
		t.Fatal("alvo removido nao foi reatribuido")
	}
}

func TestPrioridadePrazoCurto(t *testing.T) {
	prepararTeste("")
	alvosPorAmbulancia = nil

	adicionarChamado(Chamado{ID: 101, X: 2, Y: 0, Valor: 20, Prazo: 60})
	adicionarChamado(Chamado{ID: 102, X: 0, Y: 2, Valor: 20, Prazo: 3})

	idx := selecionarChamado(
		Ambulancia{ID: 0, X: 0, Y: 0, Livre: true},
		1, nil, -1,
	)

	if idx == -1 || chamados[idx].ID != 102 {
		t.Fatalf("nao priorizou chamado com prazo curto: %d", idx)
	}
}
