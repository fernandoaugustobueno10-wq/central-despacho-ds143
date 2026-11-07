package main

const tamanhoCelula = 25

type GridEspacial struct {
	celulas   [][][]int // Matriz 2D onde cada posição guarda uma lista de índices de chamados
	numCols   int
	numLinhas int
}

func novoGridEspacial(largura, altura int) *GridEspacial {
	numCols := (largura / tamanhoCelula) + 1
	numLinhas := (altura / tamanhoCelula) + 1

	celulas := make([][][]int, numCols) // Inicializa a matriz 2D de slices
	for i := range celulas {
		celulas[i] = make([][]int, numLinhas)
	}

	return &GridEspacial{
		celulas:   celulas,
		numCols:   numCols,
		numLinhas: numLinhas,
	}
}

func (g *GridEspacial) Inserir(x, y, indice int) {
	bx := x / tamanhoCelula //coordenada convertida x do bloco em que a posição x se enquadra
	by := y / tamanhoCelula //coordenada convertida y do bloco em que a posição y se enquadra

	// Proteção contra coordenadas fora dos limites
	if bx >= 0 && bx < g.numCols && by >= 0 && by < g.numLinhas {
		g.celulas[bx][by] = append(g.celulas[bx][by], indice) // Adiciona o id do chamado no slice do bloco em que se enquadrou
	}
}

func (g *GridEspacial) BuscarMaisProximo(ax, ay int, chamados []Chamado) int {
	melhorIndice := -1
	menorDist := -1

	bx := ax / tamanhoCelula //coordenada convertida x do bloco em que a posição x da ambulância se enquadra
	by := ay / tamanhoCelula //coordenada convertida y do bloco em que a posição y da ambulância se enquadra

	// Percorre o bloco atual e os 8 vizinhos
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx, ny := bx+dx, by+dy

			if nx >= 0 && nx < g.numCols && ny >= 0 && ny < g.numLinhas {
				indices := g.celulas[nx][ny] // Pega o slice de indices de chamados naquele bloco

				for _, idx := range indices { // Itera cada indice de chamado e compara a distância da ambulância para ele com a menor distância achada até agora
					chamado := chamados[idx]
					distancia := distancia(ax, ay, chamado.X, chamado.Y)

					if melhorIndice == -1 || distancia < menorDist {
						menorDist = distancia
						melhorIndice = idx
					}
				}
			}
		}
	}

	// se não achou nada, busca no slice global com varredura linear
	if melhorIndice == -1 {
		for i, chamado := range chamados {
			distancia := distancia(ax, ay, chamado.X, chamado.Y)
			if melhorIndice == -1 || distancia < menorDist {
				menorDist = distancia
				melhorIndice = i
			}
		}
	}

	return melhorIndice
}
