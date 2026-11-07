package main

import "fmt"

func jogarTurno(turno int, ambulancias []Ambulancia) {
	for _, a := range ambulancias {
		mov := "P"

		if a.Livre {
			melhor := gridEspacial.BuscarMaisProximo(a.X, a.Y, chamados)

			if melhor != -1 {
				mov = direcao(
					a.X,
					a.Y,
					chamados[melhor].X,
					chamados[melhor].Y,
				)
			}
		} else { // Para os hospitais, que são poucos e fixos, a busca não precisa usar gridEspacial
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
