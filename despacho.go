package main

import "fmt"

func jogarTurno(turno int, ambulancias []Ambulancia) {
	for _, a := range ambulancias {
		mov := "P"

		if a.Livre {
			melhor := -1

			for i, c := range chamados {
				if melhor == -1 ||
					distancia(a.X, a.Y, c.X, c.Y) <
						distancia(a.X, a.Y, chamados[melhor].X, chamados[melhor].Y) {
					melhor = i
				}
			}

			if melhor != -1 {
				mov = direcao(
					a.X,
					a.Y,
					chamados[melhor].X,
					chamados[melhor].Y,
				)
			}
		} else {
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
