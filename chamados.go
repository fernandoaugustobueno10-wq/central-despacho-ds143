package main

var (
	tabelaChamados *TabelaHash
	filaExpiracao  *FilaExpiracao
)

func inicializarChamados() {
	chamados = nil
	tabelaChamados = novaTabelaHash()
	filaExpiracao = novaFilaExpiracao()
}

func adicionarChamado(c Chamado) {
	if tabelaChamados == nil {
		inicializarChamados()
	}
	if _, existe := tabelaChamados.Buscar(c.ID); existe {
		return
	}

	indice := len(chamados)
	chamados = append(chamados, c)
	tabelaChamados.Inserir(c.ID, indice)
	filaExpiracao.Inserir(c.ID, c.Prazo)

	if gridEspacial != nil {
		gridEspacial.Inserir(c.X, c.Y, indice)
	}
}

func removerPorID(id int) {
	if tabelaChamados == nil {
		return
	}

	indice, existe := tabelaChamados.Buscar(id)
	if !existe {
		return
	}

	ultimo := len(chamados) - 1

	if gridEspacial != nil {
		gridEspacial.Remover(indice, chamados)
	}

	if indice != ultimo {
		movido := chamados[ultimo]
		chamados[indice] = movido
		tabelaChamados.Inserir(movido.ID, indice)

		if gridEspacial != nil {
			gridEspacial.AtualizarIndice(ultimo, indice, movido)
		}
	}

	chamados = chamados[:ultimo]
	tabelaChamados.Remover(id)
}

func removerExpirados(turno int) {
	if filaExpiracao == nil || tabelaChamados == nil {
		return
	}

	for {
		proximo, existe := filaExpiracao.Proximo()

		if !existe || proximo.prazo >= turno {
			return
		}

		filaExpiracao.Retirar()

		indice, encontrado := tabelaChamados.Buscar(proximo.id)

		if encontrado && chamados[indice].Prazo == proximo.prazo {
			removerPorID(proximo.id)
		}
	}
}
