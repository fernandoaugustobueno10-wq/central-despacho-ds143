package main

import (
	"math/rand"
	"testing"
)

func TestIndiceEspacialPermaneceConsistente(t *testing.T) {
	gridEspacial = novoGridEspacial(200, 200)
	inicializarChamados()

	for i := 1; i <= 2000; i++ {
		adicionarChamado(Chamado{
			ID:    i,
			X:     i % 197,
			Y:     (i * 37) % 199,
			Prazo: 50 + i%20,
		})
	}

	for i := 1; i <= 2000; i += 3 {
		removerPorID(i)
	}

	for i, c := range chamados {
		idx, ok := tabelaChamados.Buscar(c.ID)

		if !ok || idx != i {
			t.Fatalf("hash incorreta para %d", c.ID)
		}

		bx, by := c.X/tamanhoCelula, c.Y/tamanhoCelula
		pos := gridEspacial.posicoes[i]

		if pos >= len(gridEspacial.celulas[bx][by]) ||
			gridEspacial.celulas[bx][by][pos] != i {
			t.Fatalf("indice espacial incorreto para %d", c.ID)
		}
	}

	removerExpirados(70)

	if len(chamados) != 0 {
		t.Fatalf("prazo expirado: sobraram %d", len(chamados))
	}
}

func TestBuscaEspacialEquivaleVarredura(t *testing.T) {
	rng := rand.New(rand.NewSource(123))

	for repeticao := 0; repeticao < 10; repeticao++ {
		gridEspacial = novoGridEspacial(200, 200)
		inicializarChamados()

		for i := 0; i < 500; i++ {
			adicionarChamado(Chamado{
				ID:    i + 1,
				X:     rng.Intn(200),
				Y:     rng.Intn(200),
				Prazo: 100,
			})
		}

		for k := 0; k < 100; k++ {
			x, y := rng.Intn(200), rng.Intn(200)

			idx := gridEspacial.BuscarMaisProximo(
				x, y, chamados,
			)

			if idx < 0 {
				t.Fatal("nenhum candidato")
			}

			d := distancia(
				x, y,
				chamados[idx].X,
				chamados[idx].Y,
			)

			for _, c := range chamados {
				if distancia(x, y, c.X, c.Y) < d {
					t.Fatalf("nao retornou mais proximo")
				}
			}
		}
	}
}
