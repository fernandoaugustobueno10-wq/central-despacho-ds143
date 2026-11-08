package main

var alvosPorAmbulancia []int

func alvoReservado(id int, reservados []int) bool {
	for _, reservado := range reservados {
		if reservado == id {
			return true
		}
	}
	return false
}

func pontuacaoChamado(a Ambulancia, c Chamado, turno int, alvoAtual int) int64 {
	d := distancia(a.X, a.Y, c.X, c.Y)

	if turno+d > c.Prazo {
		return -1
	}

	pontuacao := int64(c.Valor) * 100000 / int64(d+10)
	folga := c.Prazo - turno - d

	if folga < 15 {
		pontuacao += pontuacao * int64(15-folga) / 50
	}

	if c.ID == alvoAtual {
		pontuacao += pontuacao / 10
	}

	return pontuacao
}

func selecionarChamado(a Ambulancia, turno int, reservados []int, alvoAtual int) int {
	if gridEspacial == nil || len(chamados) == 0 {
		return -1
	}

	melhor := -1
	melhorPontuacao := int64(-1)
	melhorDistancia := int(^uint(0) >> 1)

	considerar := func(indice int) {
		if indice < 0 || indice >= len(chamados) {
			return
		}

		c := chamados[indice]

		if alvoReservado(c.ID, reservados) {
			return
		}

		p := pontuacaoChamado(a, c, turno, alvoAtual)

		if p < 0 {
			return
		}

		d := distancia(a.X, a.Y, c.X, c.Y)

		if p > melhorPontuacao ||
			(p == melhorPontuacao && d < melhorDistancia) ||
			(p == melhorPontuacao && d == melhorDistancia &&
				(melhor == -1 || c.ID < chamados[melhor].ID)) {

			melhor = indice
			melhorPontuacao = p
			melhorDistancia = d
		}
	}

	if alvoAtual != -1 && tabelaChamados != nil {
		if indice, existe := tabelaChamados.Buscar(alvoAtual); existe {
			considerar(indice)
		}
	}

	bx, by := a.X/tamanhoCelula, a.Y/tamanhoCelula
	maxRaio := gridEspacial.numCols + gridEspacial.numLinhas

	for raio := 0; raio <= maxRaio; raio++ {
		xMin, xMax := bx-raio, bx+raio
		yMin, yMax := by-raio, by+raio

		for x := xMin; x <= xMax; x++ {
			if x < 0 || x >= gridEspacial.numCols {
				continue
			}

			for j, y := range []int{yMin, yMax} {
				if y < 0 || y >= gridEspacial.numLinhas ||
					(j == 1 && yMin == yMax) {
					continue
				}

				for _, indice := range gridEspacial.celulas[x][y] {
					considerar(indice)
				}
			}
		}

		for y := yMin + 1; y < yMax; y++ {
			if y < 0 || y >= gridEspacial.numLinhas {
				continue
			}

			for _, x := range []int{xMin, xMax} {
				if x < 0 || x >= gridEspacial.numCols {
					continue
				}

				if x == xMax && xMin == xMax {
					continue
				}

				for _, indice := range gridEspacial.celulas[x][y] {
					considerar(indice)
				}
			}
		}

		if melhor != -1 && raio >= 3 {
			break
		}
	}

	return melhor
}
