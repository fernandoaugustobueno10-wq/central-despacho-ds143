# Relatório de Projeto – Central de Despacho

**Disciplina:** DS143 – Estruturas de Dados II  
**Integrantes:** Fernando Augusto Bueno Canquerini e 
**Etapa 1 – Grade aberta**

## 1. Busca de chamados próximos

Para localizar os chamados, utilizamos um índice espacial (`espacial.go`) que divide o mapa em blocos de 25 × 25 células. A busca começa pelas regiões próximas das ambulâncias, evitando percorrer todos os chamados. Em blocos muito cheios, utilizamos uma seleção de candidatos reutilizada durante o turno (`candidatos.go`).

Localizar um bloco custa O(1), enquanto a busca depende da quantidade de blocos e candidatos examinados. Consideramos percorrer todos os chamados, mas descartamos essa alternativa pelo custo O(n) para cada ambulância.

## 2. Controle dos chamados expirados

Utilizamos um heap mínimo (`expiracao.go`) organizado pelo prazo dos chamados. Assim, conseguimos consultar rapidamente o próximo vencimento e remover os chamados expirados sem verificar toda a coleção.

A consulta ao menor prazo custa O(1), enquanto inserções e remoções do topo custam O(log n). Consideramos verificar o prazo de todos os chamados a cada turno, mas essa alternativa teria custo O(n), mesmo quando poucos estivessem vencidos.

## 3. Remoção de chamados pelo identificador

Implementamos uma tabela hash própria (`hash.go`) que relaciona o ID de cada chamado com sua posição no slice. Ao remover um chamado, substituímos sua posição pelo último elemento e atualizamos as referências na hash e no índice espacial.

As operações na hash possuem custo médio O(1), podendo chegar a O(n) em situações desfavoráveis. Consideramos buscar o ID diretamente no slice, mas isso exigiria uma busca linear O(n).

## 4. Distribuição dos alvos entre ambulâncias

Para evitar que duas ambulâncias escolham o mesmo chamado, mantemos os identificadores reservados durante cada turno (`despacho.go`). Na seleção, utilizamos uma pequena tabela hash própria para verificar se o candidato já foi escolhido (`estrategia.go`).

A consulta à reserva possui custo médio O(1). Uma alternativa seria percorrer a lista de IDs reservados para cada candidato, com custo O(a), sendo `a` a quantidade de alvos reservados.

## 5. Busca do hospital mais próximo

Como a primeira etapa não possui obstáculos, utilizamos a distância de Manhattan. Em `espacial.go`, pré-calculamos as distâncias de cada célula até o hospital mais próximo, permitindo consultas em O(1) após um processamento inicial proporcional à área do mapa.

Para ambulâncias ocupadas, também comparamos diretamente os hospitais disponíveis (`despacho.go`), com custo O(h), sendo `h` a quantidade de hospitais. O pré-processamento evita repetir essas comparações durante a avaliação de cada chamado.