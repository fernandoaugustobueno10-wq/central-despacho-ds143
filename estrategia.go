package main

var (
	alvosPorAmbulancia []int
	totalTurnos        int
)

func alvoReservado(id int, reservados []int) bool {
	for _, reservado := range reservados {
		if reservado == id {
			return true
		}
	}
	return false
}

func distanciaHospital(x, y int) int {
	if gridEspacial != nil && len(gridEspacial.distanciasHospital) > 0 &&
		x >= 0 && x < gridEspacial.largura &&
		y >= 0 && y < gridEspacial.altura {
		return int(gridEspacial.distanciasHospital[y*gridEspacial.largura+x])
	}
	if len(hospitais) == 0 {
		return 0
	}

	menor := distancia(x, y, hospitais[0].X, hospitais[0].Y)
	for _, h := range hospitais[1:] {
		d := distancia(x, y, h.X, h.Y)
		if d < menor {
			menor = d
		}
	}
	return menor
}

func pontuacaoChamado(a Ambulancia, c Chamado, turno int, alvoAtual int) int64 {
	ida := distancia(a.X, a.Y, c.X, c.Y)
	if turno+ida > c.Prazo {
		return -1
	}

	volta := distanciaHospital(c.X, c.Y)
	if totalTurnos > 0 && turno+ida+volta > totalTurnos {
		return -1
	}

	pontos := int64(c.Valor) * 100000 / int64(ida+volta+10)
	folga := c.Prazo - turno - ida
	if folga < 15 {
		pontos += pontos * int64(15-folga) / 50
	}
	if c.ID == alvoAtual {
		pontos += pontos / 10
	}
	return pontos
}

func selecionarChamado(a Ambulancia, turno int, reservados []int, alvoAtual int) int {
	if len(chamados) == 0 || gridEspacial == nil {
		return -1
	}

	melhor := -1
	melhorPontos := int64(-1)
	menorDistancia := int(^uint(0) >> 1)

	considerar := func(indice int) {
		if indice < 0 || indice >= len(chamados) {
			return
		}
		c := chamados[indice]
		if alvoReservado(c.ID, reservados) {
			return
		}
		pontos := pontuacaoChamado(a, c, turno, alvoAtual)
		if pontos < 0 {
			return
		}
		d := distancia(a.X, a.Y, c.X, c.Y)
		if pontos > melhorPontos ||
			(pontos == melhorPontos && d < menorDistancia) ||
			(pontos == melhorPontos && d == menorDistancia &&
				(melhor == -1 || c.ID < chamados[melhor].ID)) {
			melhor = indice
			melhorPontos = pontos
			menorDistancia = d
		}
	}

	if alvoAtual >= 0 && tabelaChamados != nil {
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
				if x < 0 || x >= gridEspacial.numCols ||
					(x == xMax && xMin == xMax) {
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
