# Central de Despacho

Trabalho desenvolvido para a disciplina DS143 – Estruturas de Dados II, do curso de Tecnologia em Análise e Desenvolvimento de Sistemas (TADS), da Universidade Federal do Paraná (UFPR).

**Semestre:** 2026/2

## Integrantes

- Fernando Augusto Bueno Canquerini – GRR20254201
- Fabricio Renan Plautz - GRR20253136

## Sobre o projeto

Neste trabalho, desenvolvemos um bot em Go para controlar uma equipe de ambulâncias em uma cidade representada por uma grade.

A cada turno, o servidor informa quais chamados surgiram, quais foram recolhidos e onde estão as ambulâncias. Com essas informações, nosso programa decide para onde cada ambulância deve se movimentar.

O objetivo é atender os chamados dentro do prazo e levar os pacientes até os hospitais para somar pontos. Para isso, precisamos escolher bons destinos, evitar que duas ambulâncias sejam enviadas ao mesmo chamado e fazer com que o programa tome essas decisões rapidamente, mesmo quando há muitos chamados no mapa.

## Etapa 1 – Grade aberta

Na primeira etapa, trabalhamos com um mapa sem obstáculos. Por isso, utilizamos a distância de Manhattan para calcular os deslocamentos das ambulâncias.

Um dos principais desafios foi lidar com a quantidade de chamados. No cenário de estresse, podem existir aproximadamente 500 mil chamados abertos ao mesmo tempo. Percorrer todos eles para tomar cada decisão deixaria o programa muito lento.

Para resolver esse problema, organizamos o código em módulos e implementamos estruturas de dados específicas para as principais operações.

### Estruturas de dados utilizadas

**Tabela hash (`hash.go`)**

Utilizamos uma tabela hash própria para localizar os chamados pelo identificador. Isso é importante porque o servidor informa apenas o ID quando um chamado é recolhido.

Com a hash, conseguimos encontrar e remover esses chamados sem precisar procurar por toda a lista.

**Heap mínimo (`expiracao.go`)**

Criamos um heap mínimo para controlar os prazos dos chamados. A estrutura mantém os menores prazos no topo, permitindo identificar quais chamados expiraram.

Dessa forma, não precisamos verificar individualmente todos os chamados a cada turno.

**Índice espacial (`espacial.go`)**

Dividimos o mapa em blocos de 25 por 25 células para organizar os chamados de acordo com sua localização.

Quando uma ambulância precisa escolher um destino, o programa começa procurando nas regiões próximas, em vez de percorrer o mapa inteiro.

Também utilizamos esse arquivo para calcular antecipadamente as distâncias até os hospitais mais próximos.

**Seleção de candidatos (`candidatos.go`)**

Durante os testes, percebemos que alguns blocos concentravam muitos chamados, aumentando o tempo necessário para escolher os destinos.

Para melhorar isso, implementamos uma seleção de candidatos que combina chamados com bom potencial de pontuação e uma amostragem dos demais. Essa seleção é reaproveitada pelas ambulâncias durante o mesmo turno, evitando repetir parte do processamento.

**Gerenciamento de alvos (`despacho.go` e `estrategia.go`)**

Implementamos um controle dos chamados escolhidos pelas ambulâncias.

Quando uma ambulância livre recebe um alvo, seu identificador é reservado. Assim, as próximas ambulâncias não selecionam o mesmo chamado naquele turno.

O programa também considera o alvo que a ambulância já estava seguindo, para evitar mudanças desnecessárias de destino.

### Como escolhemos os chamados

Nossa estratégia não considera apenas qual chamado está mais perto.

Para decidir o destino de cada ambulância, levamos em conta o valor do chamado, a distância necessária para chegar até ele, o caminho de volta ao hospital, o prazo de atendimento e a quantidade de turnos restantes na partida.

Também verificamos se ainda existe tempo suficiente para recolher o paciente e entregá-lo a um hospital antes do fim do jogo.

Quando uma ambulância já está transportando um paciente, ela segue em direção ao hospital mais próximo.

## Organização dos arquivos

Separamos o projeto em arquivos para facilitar o desenvolvimento, os testes e a manutenção do código.

| Arquivo | Responsabilidade |
|---|---|
| `main.go` | Inicia o programa e controla a execução dos turnos. |
| `modelos.go` | Define as estruturas que representam chamados, ambulâncias e hospitais. |
| `protocolo.go` | Realiza a comunicação com o servidor. |
| `chamados.go` | Controla a inserção, remoção e expiração dos chamados. |
| `hash.go` | Implementa a tabela hash utilizada na busca por identificador. |
| `expiracao.go` | Implementa o heap mínimo utilizado no controle dos prazos. |
| `espacial.go` | Organiza os chamados por região e calcula distâncias até hospitais. |
| `candidatos.go` | Seleciona e reutiliza candidatos nos blocos com muitos chamados. |
| `estrategia.go` | Avalia os chamados e escolhe os destinos das ambulâncias. |
| `despacho.go` | Distribui os alvos e envia os comandos de movimento. |
| `movimento.go` | Calcula as distâncias e as direções de movimento. |
| `*_test.go` | Contém os testes automatizados do projeto. |

## Requisitos

Para compilar e executar o projeto, é necessário ter o Go 1.22 ou uma versão mais recente instalada.

Também utilizamos o servidor e o corretor disponibilizados pelo professor para testar as partidas.

Seguindo as restrições do trabalho, não utilizamos o tipo `map` nativo do Go, os pacotes `container/heap`, `container/list` e `container/ring`, nem bibliotecas de terceiros na implementação.

## Como compilar

Abra o terminal na pasta principal do projeto e execute:

```bash
go build .
```

## Como executar os testes

Para executar todos os testes automatizados:

```bash
go test ./... -count=1
```

Para verificar possíveis problemas no código:

```bash
go vet ./...
```

## Resultados da Etapa 1

Testamos o programa nos três cenários disponibilizados pelo professor: pequeno, médio e estresse.

Durante o desenvolvimento, fizemos ajustes na estratégia e nas estruturas utilizadas para melhorar a pontuação e reduzir o tempo de processamento.

A principal dificuldade foi conseguir avaliar bons chamados sem deixar o programa lento no cenário de estresse. Depois das otimizações, nossa versão alcançou **45 de 45 pontos no corretor da Etapa 1**, utilizando as dez sementes de treinamento de cada cenário.

Esse foi o resultado obtido nos nossos testes. A correção definitiva poderá utilizar outras sementes.

## Etapa 2

Na segunda etapa, o mapa passa a ter quarteirões e ruas bloqueadas.

Com isso, a distância de Manhattan deixa de ser suficiente para determinar os caminhos das ambulâncias, sendo necessário considerar os obstáculos e calcular rotas válidas.

Essas funcionalidades serão desenvolvidas na continuação do projeto.