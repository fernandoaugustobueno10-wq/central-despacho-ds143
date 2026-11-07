package main

var tabelaChamados *TabelaHash

func removerPorID(id int) {
	indice, achou := tabelaChamados.Buscar(id)
	if !achou {
		return
	}

	// Remove de chamados (troca de posição com o último)
	ultimo := len(chamados) - 1
	if indice != ultimo {
		chamados[indice] = chamados[ultimo]                 // Move o último elemento para a posição do removido
		tabelaChamados.Inserir(chamados[indice].ID, indice) // atualize na hash
	}

	// Encolhe o slice
	chamados = chamados[:ultimo]

	tabelaChamados.Remover(id)
}

func removerExpirados(turno int) {
	// filtra os ativos e reconstrui/atualiza a hash.
	ativos := make([]Chamado, 0, len(chamados))
	tabelaChamados = novaTabelaHash()

	for _, c := range chamados {
		if c.Prazo >= turno {
			novoIndice := len(ativos)
			ativos = append(ativos, c)
			tabelaChamados.Inserir(c.ID, novoIndice) // Insere o ID e o novo índice na tabela hash
		}
	}

	chamados = ativos
}
