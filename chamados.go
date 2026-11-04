package main

func removerPorID(id int) {
	for i := range chamados {
		if chamados[i].ID == id {
			chamados[i] = chamados[len(chamados)-1]
			chamados = chamados[:len(chamados)-1]
			return
		}
	}
}

func removerExpirados(turno int) {
	vivos := chamados[:0]

	for _, c := range chamados {
		if c.Prazo >= turno {
			vivos = append(vivos, c)
		}
	}

	chamados = vivos
}
