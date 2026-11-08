package main

const tamanhoCelula = 25

type GridEspacial struct {
	celulas            [][][]int
	posicoes           []int
	numCols            int
	numLinhas          int
	largura            int
	altura             int
	distanciasHospital []uint16
}

func novoGridEspacial(largura, altura int) *GridEspacial {
	cols := (largura + tamanhoCelula - 1) / tamanhoCelula
	linhas := (altura + tamanhoCelula - 1) / tamanhoCelula

	celulas := make([][][]int, cols)

	for x := range celulas {
		celulas[x] = make([][]int, linhas)
	}

	return &GridEspacial{
		celulas:   celulas,
		numCols:   cols,
		numLinhas: linhas,
		largura:   largura,
		altura:    altura,
	}
}

func (g *GridEspacial) Inserir(x, y, indice int) {
	bx, by := x/tamanhoCelula, y/tamanhoCelula

	if bx < 0 || bx >= g.numCols || by < 0 || by >= g.numLinhas {
		return
	}

	for len(g.posicoes) <= indice {
		g.posicoes = append(g.posicoes, 0)
	}

	g.posicoes[indice] = len(g.celulas[bx][by])
	g.celulas[bx][by] = append(g.celulas[bx][by], indice)
}

func (g *GridEspacial) Remover(indice int, todos []Chamado) {
	c := todos[indice]

	bx, by := c.X/tamanhoCelula, c.Y/tamanhoCelula

	indices := g.celulas[bx][by]
	pos := g.posicoes[indice]
	ultimo := indices[len(indices)-1]

	indices[pos] = ultimo
	g.posicoes[ultimo] = pos

	g.celulas[bx][by] = indices[:len(indices)-1]
}

func (g *GridEspacial) AtualizarIndice(antigo, novo int, c Chamado) {
	bx, by := c.X/tamanhoCelula, c.Y/tamanhoCelula
	pos := g.posicoes[antigo]

	g.celulas[bx][by][pos] = novo
	g.posicoes[novo] = pos
}

func (g *GridEspacial) BuscarMaisProximo(
	ax, ay int,
	todos []Chamado,
) int {
	if len(todos) == 0 {
		return -1
	}

	bx, by := ax/tamanhoCelula, ay/tamanhoCelula

	melhor := -1
	menorDist := int(^uint(0) >> 1)

	maxRaio := g.numCols + g.numLinhas

	for raio := 0; raio <= maxRaio; raio++ {
		xMin, xMax := bx-raio, bx+raio
		yMin, yMax := by-raio, by+raio

		for x := xMin; x <= xMax; x++ {
			if x < 0 || x >= g.numCols {
				continue
			}

			for j, y := range []int{yMin, yMax} {
				if y < 0 || y >= g.numLinhas ||
					(j == 1 && yMin == yMax) {
					continue
				}

				for _, idx := range g.celulas[x][y] {
					c := todos[idx]
					d := distancia(ax, ay, c.X, c.Y)

					if d < menorDist {
						menorDist = d
						melhor = idx
					}
				}
			}
		}

		for y := yMin + 1; y < yMax; y++ {
			if y < 0 || y >= g.numLinhas {
				continue
			}

			for _, x := range []int{xMin, xMax} {
				if x < 0 || x >= g.numCols ||
					(x == xMax && xMin == xMax) {
					continue
				}

				for _, idx := range g.celulas[x][y] {
					c := todos[idx]
					d := distancia(ax, ay, c.X, c.Y)

					if d < menorDist {
						menorDist = d
						melhor = idx
					}
				}
			}
		}

		if melhor != -1 {
			limite := int(^uint(0) >> 1)

			if xMin > 0 {
				d := ax - xMin*tamanhoCelula + 1

				if d < limite {
					limite = d
				}
			}

			if xMax < g.numCols-1 {
				d := (xMax+1)*tamanhoCelula - ax

				if d < limite {
					limite = d
				}
			}

			if yMin > 0 {
				d := ay - yMin*tamanhoCelula + 1

				if d < limite {
					limite = d
				}
			}

			if yMax < g.numLinhas-1 {
				d := (yMax+1)*tamanhoCelula - ay

				if d < limite {
					limite = d
				}
			}

			if menorDist <= limite {
				return melhor
			}
		}
	}

	return melhor
}

func (g *GridEspacial) PrepararHospitais(hospitais []Hospital) {
	if len(hospitais) == 0 {
		g.distanciasHospital = nil
		return
	}
	largura, altura := g.largura, g.altura
	valores := make([]uint16, largura*altura)
	for i := range valores {
		valores[i] = 65535
	}
	for _, h := range hospitais {
		if h.X >= 0 && h.X < largura && h.Y >= 0 && h.Y < altura {
			valores[h.Y*largura+h.X] = 0
		}
	}
	for y := 0; y < altura; y++ {
		for x := 0; x < largura; x++ {
			i := y*largura + x
			if x > 0 && valores[i-1] < valores[i] {
				valores[i] = valores[i-1] + 1
			}
			if y > 0 && valores[i-largura] < valores[i] {
				valores[i] = valores[i-largura] + 1
			}
		}
	}
	for y := altura - 1; y >= 0; y-- {
		for x := largura - 1; x >= 0; x-- {
			i := y*largura + x
			if x+1 < largura && valores[i+1] < valores[i] {
				valores[i] = valores[i+1] + 1
			}
			if y+1 < altura && valores[i+largura] < valores[i] {
				valores[i] = valores[i+largura] + 1
			}
		}
	}
	g.distanciasHospital = valores
}
