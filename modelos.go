package main

type Chamado struct {
	ID    int
	X, Y  int
	Valor int
	Prazo int
}

type Ambulancia struct {
	ID    int
	X, Y  int
	Livre bool
}

type Hospital struct {
	X, Y int
}

var (
	hospitais []Hospital
	chamados  []Chamado
)
