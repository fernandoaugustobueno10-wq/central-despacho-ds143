package main

type resumoBloco struct {
	turno   int
	tamanho int
	indices []int
}

var resumos []resumoBloco
var ultimaGrade *GridEspacial

func candidatosDoBloco(bx, by, turno int) []int {
	g := gridEspacial

	if ultimaGrade != g || len(resumos) != g.numCols*g.numLinhas {
		resumos = make([]resumoBloco, g.numCols*g.numLinhas)
		ultimaGrade = g
	}

	pos := by*g.numCols + bx
	r := &resumos[pos]
	todos := g.celulas[bx][by]

	if r.turno == turno && r.tamanho == len(todos) && r.indices != nil {
		return r.indices
	}

	r.turno, r.tamanho = turno, len(todos)
	r.indices = r.indices[:0]

	const melhores = 96

	type opcao struct {
		indice int
		score  int64
	}

	var heap [melhores]opcao
	n := 0

	for _, indice := range todos {
		c := chamados[indice]
		d := distanciaHospital(c.X, c.Y)
		rank := int64(c.Valor) * 100000 / int64(d+25)

		if n < melhores {
			i := n
			n++

			for i > 0 && heap[(i-1)/2].score > rank {
				heap[i] = heap[(i-1)/2]
				i = (i - 1) / 2
			}
			heap[i] = opcao{indice, rank}
		} else if rank > heap[0].score {
			i := 0

			for {
				child := 2*i + 1
				if child >= n {
					break
				}
				if child+1 < n && heap[child+1].score < heap[child].score {
					child++
				}
				if heap[child].score >= rank {
					break
				}
				heap[i] = heap[child]
				i = child
			}

			heap[i] = opcao{indice, rank}
		}
	}

	for i := 0; i < n; i++ {
		r.indices = append(r.indices, heap[i].indice)
	}

	const amostras = 96

	passo := len(todos) / amostras
	if passo == 0 {
		passo = 1
	}

	off := (bx*31 + by*17 + turno) % passo

	for i, k := off, 0; i < len(todos) && k < amostras; i, k = i+passo, k+1 {
		r.indices = append(r.indices, todos[i])
	}

	return r.indices
}
