package main

import "fmt"

func jogarTurno(turno int, ambulancias []Ambulancia) {
	reservados := make([]int, 0, len(ambulancias))

	for _, a := range ambulancias {
		if a.ID >= len(alvosPorAmbulancia) {
			anteriores := alvosPorAmbulancia
			alvosPorAmbulancia = make([]int, a.ID+1)
			copy(alvosPorAmbulancia, anteriores)
		}
	}

	for _, a := range ambulancias {
		if !a.Livre {
			alvosPorAmbulancia[a.ID] = 0
			continue
		}

		alvoAtual := alvosPorAmbulancia[a.ID] - 1
		melhor := selecionarChamado(a, turno, reservados, alvoAtual)

		if melhor == -1 {
			alvosPorAmbulancia[a.ID] = 0
			continue
		}

		id := chamados[melhor].ID
		alvosPorAmbulancia[a.ID] = id + 1
		reservados = append(reservados, id)
	}

	for _, a := range ambulancias {
		mov := "P"

		if a.Livre {
			id := alvosPorAmbulancia[a.ID] - 1

			if id >= 0 {
				if indice, existe := tabelaChamados.Buscar(id); existe {
					c := chamados[indice]
					mov = direcao(a.X, a.Y, c.X, c.Y)
				}
			}
		} else if len(hospitais) > 0 {
			melhor := 0

			for i, h := range hospitais {
				if distancia(a.X, a.Y, h.X, h.Y) <
					distancia(a.X, a.Y, hospitais[melhor].X, hospitais[melhor].Y) {
					melhor = i
				}
			}

			mov = direcao(
				a.X,
				a.Y,
				hospitais[melhor].X,
				hospitais[melhor].Y,
			)
		}

		fmt.Fprintln(escritor, a.ID, mov)
	}

	fmt.Fprintln(escritor, "FIM")
	escritor.Flush()
}
