package main

type vencimento struct {
	id    int
	prazo int
}

type FilaExpiracao struct {
	itens []vencimento
}

func novaFilaExpiracao() *FilaExpiracao {
	return &FilaExpiracao{
		itens: make([]vencimento, 0, 2048),
	}
}

func (f *FilaExpiracao) Inserir(id, prazo int) {
	f.itens = append(f.itens, vencimento{
		id:    id,
		prazo: prazo,
	})

	i := len(f.itens) - 1

	for i > 0 {
		pai := (i - 1) / 2

		if f.itens[pai].prazo <= f.itens[i].prazo {
			break
		}

		f.itens[pai], f.itens[i] = f.itens[i], f.itens[pai]
		i = pai
	}
}

func (f *FilaExpiracao) Proximo() (vencimento, bool) {
	if len(f.itens) == 0 {
		return vencimento{}, false
	}

	return f.itens[0], true
}

func (f *FilaExpiracao) Retirar() (vencimento, bool) {
	if len(f.itens) == 0 {
		return vencimento{}, false
	}

	primeiro := f.itens[0]
	ultimo := len(f.itens) - 1

	f.itens[0] = f.itens[ultimo]
	f.itens = f.itens[:ultimo]

	n := len(f.itens)

	for i := 0; i < n; {
		esq := 2*i + 1

		if esq >= n {
			break
		}

		menor := esq
		dir := esq + 1

		if dir < n && f.itens[dir].prazo < f.itens[esq].prazo {
			menor = dir
		}

		if f.itens[i].prazo <= f.itens[menor].prazo {
			break
		}

		f.itens[i], f.itens[menor] = f.itens[menor], f.itens[i]
		i = menor
	}

	return primeiro, true
}
