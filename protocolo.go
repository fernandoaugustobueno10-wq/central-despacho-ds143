package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	leitor   *bufio.Scanner
	escritor *bufio.Writer
)

func lerCampos() []string {
	if !leitor.Scan() {
		return nil
	}

	return strings.Fields(leitor.Text())
}

func inteiro(s string) int {
	v, err := strconv.Atoi(s)

	if err != nil {
		panic("número inválido na entrada: " + s)
	}

	return v
}

func lerInicio() {
	camposInicio := lerCampos()

	largura := inteiro(camposInicio[1])
	altura := inteiro(camposInicio[2])

	gridEspacial = novoGridEspacial(largura, altura)
	inicializarChamados()

	h := inteiro(lerCampos()[1])

	for i := 0; i < h; i++ {
		c := lerCampos()

		hospitais = append(hospitais, Hospital{
			inteiro(c[0]),
			inteiro(c[1]),
		})
	}

	k := inteiro(lerCampos()[1])

	for i := 0; i < k; i++ {
		lerCampos()
	}

	b := inteiro(lerCampos()[1])

	for i := 0; i < b; i++ {
		lerCampos()
	}

	lerCampos()
}

func configurarEntradaSaida() {
	leitor = bufio.NewScanner(os.Stdin)
	leitor.Buffer(make([]byte, 1<<16), 1<<20)

	escritor = bufio.NewWriter(os.Stdout)
}

func lerTurno() (int, []Ambulancia, bool) {
	c := lerCampos()

	if c == nil {
		return 0, nil, false
	}

	turno := inteiro(c[1])

	n := inteiro(lerCampos()[1])

	for i := 0; i < n; i++ {
		c := lerCampos()

		adicionarChamado(Chamado{
			inteiro(c[0]),
			inteiro(c[1]),
			inteiro(c[2]),
			inteiro(c[3]),
			inteiro(c[4]),
		})
	}

	r := inteiro(lerCampos()[1])

	for i := 0; i < r; i++ {
		removerPorID(inteiro(lerCampos()[0]))
	}

	removerExpirados(turno)

	k := inteiro(lerCampos()[1])
	ambulancias := make([]Ambulancia, k)

	for i := 0; i < k; i++ {
		c := lerCampos()

		ambulancias[i] = Ambulancia{
			inteiro(c[0]),
			inteiro(c[1]),
			inteiro(c[2]),
			c[3] == "LIVRE",
		}
	}

	e := inteiro(lerCampos()[1])

	for i := 0; i < e; i++ {
		lerCampos()
	}

	lerCampos()

	return turno, ambulancias, true
}

func enviarPronto() {
	fmt.Fprintln(escritor, "PRONTO")
	escritor.Flush()
}
