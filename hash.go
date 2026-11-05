package main

type entradaHash struct {
	id      int
	indice  int
	proximo *entradaHash
}

type TabelaHash struct {
	baldes     []*entradaHash
	quantidade int
}

func novaTabelaHash() *TabelaHash {
	return &TabelaHash{
		baldes: make([]*entradaHash, 1024),
	}
}

func (h *TabelaHash) posicao(id int) int {
	x := uint64(id)
	x ^= x >> 33
	x *= 0xff51afd7ed558ccd
	x ^= x >> 33
	x *= 0xc4ceb9fe1a85ec53
	x ^= x >> 33

	return int(x % uint64(len(h.baldes)))
}

func (h *TabelaHash) Inserir(id, indice int) {
	pos := h.posicao(id)

	for atual := h.baldes[pos]; atual != nil; atual = atual.proximo {
		if atual.id == id {
			atual.indice = indice
			return
		}
	}

	if h.quantidade >= len(h.baldes)*3/4 {
		h.redimensionar()
		pos = h.posicao(id)
	}

	h.baldes[pos] = &entradaHash{
		id:      id,
		indice:  indice,
		proximo: h.baldes[pos],
	}
	h.quantidade++
}

func (h *TabelaHash) Buscar(id int) (int, bool) {
	pos := h.posicao(id)

	for atual := h.baldes[pos]; atual != nil; atual = atual.proximo {
		if atual.id == id {
			return atual.indice, true
		}
	}

	return 0, false
}

func (h *TabelaHash) Remover(id int) bool {
	pos := h.posicao(id)

	var anterior *entradaHash
	atual := h.baldes[pos]

	for atual != nil {
		if atual.id == id {
			if anterior == nil {
				h.baldes[pos] = atual.proximo
			} else {
				anterior.proximo = atual.proximo
			}

			h.quantidade--
			return true
		}

		anterior = atual
		atual = atual.proximo
	}

	return false
}

func (h *TabelaHash) redimensionar() {
	antigos := h.baldes
	h.baldes = make([]*entradaHash, len(antigos)*2)

	for _, entrada := range antigos {
		for entrada != nil {
			proximo := entrada.proximo
			pos := h.posicao(entrada.id)

			entrada.proximo = h.baldes[pos]
			h.baldes[pos] = entrada

			entrada = proximo
		}
	}
}
